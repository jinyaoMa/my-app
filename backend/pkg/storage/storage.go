package storage

import (
	"crypto/md5"
	"crypto/sha512"
	"errors"
	"fmt"
	"io/fs"
	"my-app/backend/pkg/utils"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const (
	B  uint64 = 1
	KB        = 1024 * B
	MB        = 1024 * KB
	GB        = 1024 * MB
	TB        = 1024 * GB
)

type StorageCacheFile struct {
	File *os.File
	Path string
}

type StoragePath struct {
	Dir   string
	Cache string
}

type Storage struct {
	paths []*StoragePath
}

// GetCacheFiles implements Interface.
func (s *Storage) GetCacheFiles(filename string) (files []*StorageCacheFile, err error) {
	for _, sPath := range s.paths {
		err = filepath.WalkDir(sPath.Cache, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if !d.IsDir() && strings.HasPrefix(filepath.Base(path), filename) {
				var file *os.File
				file, err = os.Open(path)
				if err != nil {
					return err
				}
				files = append(files, &StorageCacheFile{
					File: file,
					Path: path,
				})
			}
			return nil
		})
		if err != nil {
			for _, file := range files {
				defer file.File.Close()
			}
			return nil, err
		}
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Path < files[j].Path
	})
	return
}

// SearchFile implements Interface.
func (s *Storage) SearchFile(filename string, isCache bool) (file *os.File, path string, err error) {
	for _, sPath := range s.paths {
		if isCache {
			path = filepath.Join(sPath.Cache, filename)
		} else {
			path = filepath.Join(sPath.Dir, filename)
		}
		if utils.CheckIfFileExists(path) {
			file, err = os.OpenFile(path, os.O_RDWR|os.O_TRUNC, 0666)
			return
		}
	}
	return nil, "", nil
}

// Cache implements Interface.
func (s *Storage) Cache(filename string, data []byte, rangeStart uint64, rangeEnd uint64, size uint64, forceCache bool) (ok bool, path string, err error) {
	cacheFilename := fmt.Sprintf("%s.%d.%d", filename, rangeStart, rangeEnd)

	var file *os.File
	file, path, err = s.SearchFile(cacheFilename, true)
	if err != nil {
		return
	}

	if file != nil {
		defer file.Close()

		if forceCache {
			var wSize int
			wSize, err = file.Write(data)
			if err != nil {
				return
			}
			return uint64(wSize) == size, path, nil
		}

		return false, "", errors.New(cacheFilename + "had already existed")
	}

	var u MountpointUsage
	u, err = s.GetMountpointUsage()
	if err != nil {
		return
	}

	sPath := u.PickAPath(size)
	cacheFilePath := filepath.Join(sPath.Cache, cacheFilename)
	file, err = os.Create(cacheFilePath)
	if err != nil {
		return
	}
	defer file.Close()

	var wSize int
	wSize, err = file.Write(data)
	if err != nil {
		return
	}
	return uint64(wSize) == size, path, nil
}

// Checksum implements Interface.
func (s *Storage) Checksum(filename string, isCache bool) (checksum string, paths []string, err error) {
	var data []byte
	buffer := make([]byte, 4096)
	size := 0

	if isCache {
		var cacheFiles []*StorageCacheFile
		cacheFiles, err = s.GetCacheFiles(filename)
		if err != nil {
			return
		}
		for _, cacheFile := range cacheFiles {
			defer cacheFile.File.Close()
		}

		for _, cacheFile := range cacheFiles {
			paths = append(paths, cacheFile.Path)
			for {
				n, err := cacheFile.File.Read(buffer)
				if err != nil {
					break
				}
				data = append(data, buffer[0:n]...)
				size += n
			}
		}
	} else {
		var file *os.File
		var path string
		file, path, err = s.SearchFile(filename, false)
		if err != nil {
			return
		}
		if file == nil {
			checksum = ""
			return
		}
		defer file.Close()
		paths = append(paths, path)

		for {
			n, err := file.Read(buffer)
			if err != nil {
				break
			}
			data = append(data, buffer[0:n]...)
			size += n
		}
	}

	md5Sum := md5.Sum(data)
	sha512Sum := sha512.Sum512(data)
	checksum = fmt.Sprintf("%x:%x:%d", md5Sum, sha512Sum, size)
	return
}

// VerifyChecksum implements Interface.
func (s *Storage) VerifyChecksum(filename string, isCache bool, checksum string) (ok bool, paths []string, err error) {
	var sum string
	sum, paths, err = s.Checksum(filename, isCache)
	if err != nil {
		return
	}
	ok = sum == checksum
	return
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
func (*Storage) Persist(filepaths []string) (err error) {
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
