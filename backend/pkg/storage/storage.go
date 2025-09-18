package storage

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/disk"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/base/hash2"
	"majinyao.cn/my-app/backend/pkg/utils"
)

type IStorage interface {
	GetHashAlgs() (algs []string)
	GetLibraryAPathMaxLength() int
	PersistTemporary(ctx context.Context, namespace string, checksum string, size uint64, ext string) (checksums []string, err error)
	CacheTemporary(ctx context.Context, namespace string, checksum string, size uint64, offset uint64, data io.Reader) (err error)
	SearchTemporary(namespace string, checksum string) (file *os.File, path string, dir string, err error)
	SearchDirectory(ext string, size uint64, checksums ...string) (file *os.File, path string, dir string, err error)
	BuildPersistedFilename(ext string, size uint64, checksums ...string) string
	ChooseLibraryTemporaryBySize(ctx context.Context, size uint64, namespace string) (nsDir string, err error)
	ChooseBySize(ctx context.Context, size uint64) (partition Partition, dir string, err error)
	EnableLibraryByMountpoint(mountpoint string) (err error)
	DisableLibraryByMountpoint(mountpoint string)
	GetLibraryTemporaryByMountpoint(ctx context.Context, mountpoint string) (tmp string, ok bool)
	GetLibraryDirectoryByMountpoint(ctx context.Context, mountpoint string) (dir string, ok bool)
	GetPartitions(ctx context.Context) (partitions []Partition, err error)
	AddLibraries(ctx context.Context, ls ...Library) (err error)
	LoadLibraries(ctx context.Context, temp string, ls ...Library) (err error)
	GetLibraries() []Library
	CreateLibraryTemporary(l Library) (err error)
}

func MustNew(options Options) IStorage {
	s, err := New(options)
	if err != nil {
		panic(err)
	}
	return s
}

func New(options Options) (IStorage, error) {
	return new(storage).init(options)
}

type storage struct {
	libraries     []Library
	temporary     string
	bufferSize    uint64
	combo         hasher.ICombo
	maxPathLength int
}

func (s *storage) GetHashAlgs() (algs []string) {
	s.combo.Using(func(hs hash2.Hash2s) {
		algs = hs.Algs()
	})
	return
}

func (s *storage) GetLibraryAPathMaxLength() int {
	// assume max path length 254 bytes,
	// filename:
	// 64 bytes for sha3-256 hex,
	// 16 bytes for crc64 hex,
	// 16 bytes for extra xxh3 hex,
	// 16 bytes for file size hex,
	// 1/2 of 64+16+16+16 bytes for extension,
	// then left 86 bytes for library absolute path
	return s.maxPathLength - (utils.SliceSum(s.combo.OutputLengths())*2 + 16) - (utils.SliceSum(s.combo.OutputLengths())*2+16)/2
}

func (s *storage) PersistTemporary(ctx context.Context, namespace string, checksum string, size uint64, ext string) (checksums []string, err error) {
	tmpFile, path, dir, err := s.SearchTemporary(namespace, checksum)
	if err != nil {
		return nil, err
	}
	stat, err := tmpFile.Stat()
	if err != nil {
		return nil, err
	}
	if stat.Size() != int64(size) {
		return nil, errors.New("storage persist temporary: size not match")
	}

	s.combo.Using(func(hs hash2.Hash2s) {
		defer func() {
			if err == nil {
				checksums = hs.SumBase64()
			}
		}()

		buffer := make([]byte, s.bufferSize)
		for {
			select {
			case <-ctx.Done():
				err = ctx.Err()
				return
			default:
				n, errRead := tmpFile.Read(buffer)
				if errRead == io.EOF {
					return
				}
				if errRead != nil {
					err = errRead
					return
				}

				m, errWrite := hs.Write(buffer[:n])
				if errWrite != nil {
					err = errWrite
					return
				}
				if n != m {
					err = errors.New("storage persist temporary: combo hasher write not all bytes")
					return
				}
			}
		}
	})

	tmpFile.Close()
	if err != nil {
		return nil, err
	}
	if !slices.Contains(checksums, checksum) {
		return nil, errors.New("storage persist temporary: checksum not match")
	}

	newPath := filepath.Join(dir, s.BuildPersistedFilename(ext, size, checksums...))
	err = os.Rename(path, newPath)
	if err != nil {
		return nil, err
	}
	return
}

func (s *storage) CacheTemporary(ctx context.Context, namespace string, checksum string, size uint64, offset uint64, data io.Reader) (err error) {
	tmpFile, _, _, err := s.SearchTemporary(namespace, checksum)
	if err != nil {
		var nsDir string
		nsDir, err = s.ChooseLibraryTemporaryBySize(ctx, size, namespace)
		if err != nil {
			return err
		}

		err = os.MkdirAll(nsDir, os.ModeDir)
		if err != nil {
			return err
		}

		path := filepath.Join(nsDir, checksum)
		tmpFile, err = os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
	}
	defer tmpFile.Close()

	ret, err := tmpFile.Seek(int64(offset), io.SeekStart)
	if err != nil {
		return err
	}
	if ret < 0 {
		return errors.New("storage cache temporary: seek not to offset")
	}

	buffer := make([]byte, s.bufferSize)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			n, err := data.Read(buffer)
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}

			m, err := tmpFile.Write(buffer[:n])
			if err != nil {
				return err
			}
			if n != m {
				return errors.New("storage cache temporary: write not all bytes")
			}
		}
	}
}

func (s *storage) SearchTemporary(namespace string, checksum string) (file *os.File, path string, dir string, err error) {
	ls := s.GetLibraries()
	for _, l := range ls {
		path = filepath.Join(l.Directory, s.temporary, namespace, checksum)
		file, err = os.OpenFile(path, os.O_RDWR, 0644)
		if err == nil {
			return file, path, l.Directory, nil
		}
	}
	return nil, "", "", errors.New("storage search temporary: not found")
}

func (s *storage) SearchDirectory(ext string, size uint64, checksums ...string) (file *os.File, path string, dir string, err error) {
	ls := s.GetLibraries()
	for _, l := range ls {
		path = filepath.Join(l.Directory, s.BuildPersistedFilename(ext, size, checksums...))
		file, err = os.OpenFile(path, os.O_RDWR, 0644)
		if err == nil {
			return file, path, l.Directory, nil
		}
	}
	return nil, "", "", errors.New("storage search temporary: not found")
}

func (s *storage) BuildPersistedFilename(ext string, size uint64, checksums ...string) string {
	// assume `{sha256:b64},{crc64:b64},{xxh3:b64},{size:b64}.{ext}`
	if ext != "" && !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}
	commaParts := append(checksums, utils.ConvertUint64ToBase64(size))
	return strings.Join(commaParts, ",") + ext
}

func (s *storage) ChooseLibraryTemporaryBySize(ctx context.Context, size uint64, namespace string) (nsDir string, err error) {
	_, dir, err := s.ChooseBySize(ctx, size)
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, s.temporary, namespace), nil
}

func (s *storage) ChooseBySize(ctx context.Context, size uint64) (partition Partition, dir string, err error) {
	var partitions []Partition
	partitions, err = s.GetPartitions(ctx)
	if err != nil {
		return
	}
	if len(partitions) == 0 {
		err = errors.New("storage choose partition by size: no partition")
		return
	}

	ls := s.GetLibraries()
	partitions = slices.DeleteFunc(partitions, func(p Partition) bool {
		return slices.ContainsFunc(ls, func(l Library) bool {
			return !strings.EqualFold(l.Mountpoint, p.Mountpoint) || l.Disabled
		})
	})

	slices.SortStableFunc(partitions, func(a, b Partition) int {
		if a.Free > b.Free {
			return -1
		}
		return 1
	})
	if partitions[0].Free < size {
		err = errors.New("storage choose partition by size: partition not enough space")
		return
	}

	return partitions[0], ls[slices.IndexFunc(ls, func(l Library) bool {
		return strings.EqualFold(l.Mountpoint, partitions[0].Mountpoint)
	})].Directory, nil
}

func (s *storage) EnableLibraryByMountpoint(mountpoint string) (err error) {
	for i := range s.libraries {
		if strings.EqualFold(s.libraries[i].Mountpoint, mountpoint) {
			err = s.CreateLibraryTemporary(s.libraries[i])
			if err != nil {
				return errors.Join(errors.New("storage enable library by mountpoint: create library temporary failed"), err)
			}
			s.libraries[i].Disabled = false
		}
	}
	return errors.New("storage enable library by mountpoint: library not found")
}

func (s *storage) DisableLibraryByMountpoint(mountpoint string) {
	for i := range s.libraries {
		if strings.EqualFold(s.libraries[i].Mountpoint, mountpoint) {
			s.libraries[i].Disabled = true
		}
	}
}

func (s *storage) GetLibraryTemporaryByMountpoint(ctx context.Context, mountpoint string) (tmp string, ok bool) {
	dir, ok := s.GetLibraryDirectoryByMountpoint(ctx, mountpoint)
	if !ok {
		return "", false
	}
	return filepath.Join(dir, s.temporary), true
}

func (s *storage) GetLibraryDirectoryByMountpoint(ctx context.Context, mountpoint string) (dir string, ok bool) {
	if index := slices.IndexFunc(s.GetLibraries(), func(l Library) bool {
		return strings.EqualFold(l.Mountpoint, mountpoint) && !l.Disabled
	}); index != -1 {
		return s.libraries[index].Directory, true
	}
	return "", false
}

func (s *storage) GetPartitions(ctx context.Context) (partitions []Partition, err error) {
	pStats, err := disk.PartitionsWithContext(ctx, false)
	if err != nil {
		return nil, err
	}

	for _, pStat := range pStats {
		var uStat *disk.UsageStat
		uStat, err = disk.UsageWithContext(ctx, pStat.Mountpoint)
		if err != nil {
			return nil, err
		}

		partitions = append(partitions, Partition{
			Device:      pStat.Device,
			Mountpoint:  pStat.Mountpoint,
			Fstype:      pStat.Fstype,
			Opts:        pStat.Opts,
			Total:       uStat.Total,
			Free:        uStat.Free,
			Used:        uStat.Used,
			UsedPercent: uStat.UsedPercent,
		})
	}
	return
}

func (s *storage) AddLibraries(ctx context.Context, ls ...Library) (err error) {
	return s.LoadLibraries(ctx, s.temporary, append(s.libraries, ls...)...)
}

func (s *storage) LoadLibraries(ctx context.Context, temp string, ls ...Library) (err error) {
	partitions, err := s.GetPartitions(ctx)
	if err != nil {
		return errors.Join(errors.New("storage load libraries: get partitions failed"), err)
	}

	s.temporary = temp
	slices.SortStableFunc(ls, func(a, b Library) int {
		if a.Disabled {
			return -1
		}
		return 1
	})

	libraries := utils.SliceMap(partitions, func(p Partition) Library {
		return Library{
			Mountpoint: p.Mountpoint,
		}
	})
	for _, opt := range ls {
		if index := slices.IndexFunc(libraries, func(l Library) bool {
			return strings.EqualFold(l.Mountpoint, opt.Mountpoint)
		}); index != -1 {
			libraries[index] = opt
		}
	}

	maxLength := s.GetLibraryAPathMaxLength()
	s.libraries = slices.DeleteFunc(libraries, func(l Library) bool {
		return l.Directory == "" || len(l.Directory) > maxLength || !strings.HasPrefix(l.Directory, l.Mountpoint)
	})
	for _, l := range s.libraries {
		if !l.Disabled {
			temporaryPath := filepath.Join(l.Directory, s.temporary)
			err = os.MkdirAll(temporaryPath, os.ModeDir)
			if err != nil {
				return errors.Join(errors.New("storage load libraries: create temporary folder failed"), err)
			}
		}
	}
	return nil
}

func (s *storage) GetLibraries() []Library {
	for i := range s.libraries {
		if s.libraries[i].Disabled {
			continue
		}

		if err := s.CreateLibraryTemporary(s.libraries[i]); err != nil {
			s.libraries[i].Disabled = true
		}
	}
	return s.libraries
}

func (s *storage) CreateLibraryTemporary(l Library) (err error) {
	temporaryPath := filepath.Join(l.Directory, s.temporary)
	err = os.MkdirAll(temporaryPath, os.ModeDir)
	if err != nil {
		return err
	}
	return nil
}

func (s *storage) init(options Options) (*storage, error) {
	combo, err := hasher.NewCombo(options.Hashers...)
	if err != nil {
		return nil, errors.Join(errors.New("storage init failed"), err)
	}
	s.combo = combo

	// default 2 * (hashers output length sum + 16 of size b64)
	maxPathLength := (utils.SliceSum(s.combo.OutputLengths()) + 16) * 2
	if options.MaxPathLength > maxPathLength {
		maxPathLength = options.MaxPathLength
	}
	s.maxPathLength = maxPathLength
	s.bufferSize = options.BufferSize

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	defer cancel()

	err = s.LoadLibraries(ctx, options.Temporary, options.Libraries...)
	if err != nil {
		return nil, errors.Join(errors.New("storage init failed"), err)
	}
	return s, nil
}
