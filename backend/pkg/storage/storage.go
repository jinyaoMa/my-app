package storage

import (
	"crypto/md5"
	"crypto/sha512"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"my-app/backend/pkg/helper"
	"os"
	"path/filepath"
	"sort"
	"strconv"
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
	File       *os.File
	Path       string
	RangeStart uint64
	RangeEnd   uint64
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

				var rangeStart, rangeEnd uint64
				pathSplit := strings.Split(path, ".")
				if rangeStart, err = strconv.ParseUint(pathSplit[len(pathSplit)-2], 10, 64); err != nil {
					return err
				}
				if rangeEnd, err = strconv.ParseUint(pathSplit[len(pathSplit)-1], 10, 64); err != nil {
					return err
				}
				files = append(files, &StorageCacheFile{
					File:       file,
					Path:       path,
					RangeStart: rangeStart,
					RangeEnd:   rangeEnd,
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
		return files[i].RangeStart < files[j].RangeStart
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
		if helper.CheckIfFileExists(path) {
			file, err = os.OpenFile(path, os.O_RDWR, os.ModePerm)
			return
		}
	}
	return nil, "", nil
}

// Cache implements Interface.
func (s *Storage) Cache(filename string, data []byte, rangeStart int64, rangeEnd int64, size int64, forceCache bool) (ok bool, path string, err error) {
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
			return int64(wSize) == size, path, nil
		}

		return false, "", errors.New(cacheFilename + "had already existed")
	}

	var u MountpointUsage
	u, err = s.GetMountpointUsage()
	if err != nil {
		return
	}

	sPath := u.PickAPath(uint64(size))
	cacheFilePath := filepath.Join(sPath.Cache, cacheFilename)
	file, err = os.Create(cacheFilePath)
	if err != nil {
		println("err (SearchFile) =", err.Error())
		return
	}
	defer file.Close()

	var wSize int
	wSize, err = file.Write(data)
	if err != nil {
		return
	}
	return int64(wSize) == size, cacheFilePath, nil
}

// Checksum implements Interface.
func (s *Storage) Checksum(filename string, isCache bool) (checksum string, paths []string, err error) {
	var data []byte
	buffer := make([]byte, 4096)
	size := 0
	md5New := md5.New()
	sha512New := sha512.New()

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

				temp := buffer[0:n]
				_, err = md5New.Write(temp)
				if err != nil {
					break
				}
				_, err = sha512New.Write(temp)
				if err != nil {
					break
				}
				data = append(data, temp...)
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

			temp := buffer[0:n]
			_, err = md5New.Write(temp)
			if err != nil {
				break
			}
			_, err = sha512New.Write(temp)
			if err != nil {
				break
			}
			data = append(data, temp...)
			size += n
		}
	}

	checksum = fmt.Sprintf("%x-%x-%d", md5New.Sum(nil), sha512New.Sum(nil), size)
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
func (s *Storage) ClearCache(filename string) (err error) {
	for _, sPath := range s.paths {
		err = filepath.WalkDir(sPath.Cache, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if !d.IsDir() && strings.HasPrefix(filepath.Base(path), filename) {
				err = os.Remove(path)
				if err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// Load implements Interface.
func (s *Storage) Load(filename string, rangeStart int64, rangeEnd int64) (data []byte, err error) {
	if rangeStart < 0 || rangeStart > rangeEnd {
		return nil, errors.New("rangeStart should be greater than zero or less/equal to rangeEnd")
	}

	var file *os.File
	file, _, err = s.SearchFile(filename, false)
	if err != nil {
		return
	}
	if file == nil {
		return
	}
	defer file.Close()

	var n int
	dataLength := rangeEnd - rangeStart
	data = make([]byte, dataLength)
	n, err = file.ReadAt(data, rangeStart)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return data[0:n], nil
}

// Persist implements Interface.
func (s *Storage) Persist(filename string, cacheFilepaths []string, totalSize int64) (ok bool, path string, err error) {
	var u MountpointUsage
	u, err = s.GetMountpointUsage()
	if err != nil {
		return
	}

	var targetFile *os.File
	targetFile, path, err = s.SearchFile(filename, false)
	if err != nil {
		return
	}
	if targetFile != nil {
		defer targetFile.Close()
		ok = true
		return
	}

	sPath := u.PickAPath(uint64(totalSize))
	path = filepath.Join(sPath.Dir, filename)
	targetFile, err = os.Create(path)
	if err != nil {
		return
	}
	defer targetFile.Close()

	var cacheFile *os.File
	for _, cacheFilepath := range cacheFilepaths {
		cacheFile, err = os.Open(cacheFilepath)
		if err != nil {
			return
		}

		_, err = io.Copy(targetFile, cacheFile)
		if err != nil {
			cacheFile.Close()
			return
		}
		cacheFile.Close()
	}

	ok = true
	return
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
				helper.CheckIfDirectoryExists(path) {
				if !helper.CheckIfDirectoryExists(cache) &&
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
