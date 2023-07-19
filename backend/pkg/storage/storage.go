package storage

import (
	"io/fs"
	"my-app/backend/pkg/utils"
	"os"
	"path/filepath"
	"strings"
)

const (
	B  uint64 = 1
	KB        = 1024 * B
	MB        = 1024 * KB
	GB        = 1024 * MB
	TB        = 1024 * GB
)

type StoragePath struct {
	Dir   string
	Cache string
}

type Storage struct {
	paths []*StoragePath
}

// Search implements Interface.
func (s *Storage) Search(filename string, isCache bool) (file *os.File) {
	panic("unimplemented")
}

// Cache implements Interface.
func (s *Storage) Cache(filename string, data []byte, rangeStart uint64, rangeEnd uint64, totalSize uint64) (err error) {
	// var u MountpointUsage
	// u, err = s.GetMountpointUsage()
	// if err != nil {
	// 	return
	// }

	// targetPath := u.PickAPath(totalSize)
	// return
	panic("unimplemented")
}

// Checksum implements Interface.
func (*Storage) Checksum(filename string, checksum string, isCache bool) (ok bool) {
	panic("unimplemented")
}

// ClearCache implements Interface.
func (*Storage) ClearCache(filename string) (err error) {
	panic("unimplemented")
}

// Load implements Interface.
func (*Storage) Load(filename string, rangeStart uint64, rangeEnd uint64) (file fs.File, err error) {
	panic("unimplemented")
}

// Persist implements Interface.
func (*Storage) Persist(filename string) (err error) {
	panic("unimplemented")
}

// GetMountpointUsage implements Interface.
func (s *Storage) GetMountpointUsage() (u MountpointUsage, err error) {
	return NewMountpointUsage(s.paths)
}

// AddPaths implements Interface
func (s *Storage) AddPaths(paths ...string) (added int, err error) {
	var u MountpointUsage
	u, err = s.GetMountpointUsage()
	if err != nil {
		return
	}

	for mountpoint, mStat := range u {
		for _, path := range paths {
			cache := filepath.Join(path, ".cache")

			if mStat.StoragePath == nil &&
				strings.HasPrefix(path, mountpoint) &&
				utils.CheckIfDirectoryExists(path) {
				if !utils.CheckIfDirectoryExists(cache) &&
					os.MkdirAll(cache, os.ModeDir) != nil {
					continue
				}

				sPath := &StoragePath{
					Dir:   path,
					Cache: cache,
				}
				s.paths = append(s.paths, sPath)
				mStat.StoragePath = sPath
				added += 1
			}
		}
	}
	return
}

func New() Interface {
	return &Storage{
		paths: make([]*StoragePath, 0, 5),
	}
}
