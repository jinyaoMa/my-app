package fstore

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type StorageMap map[string]*Storage // pid: storage

func (storageMap StorageMap) ToList() (list []*Storage) {
	for pid := range storageMap {
		list = append(list, storageMap[pid])
	}
	return
}

type Storage struct {
	*Partition
	PID   string `json:"pid"`   // define storage by partition mountpoint
	APath string `json:"apath"` // the absolute path in partition to use for storage
	CPath string `json:"cpath"` // the absolute cache path of storage
	Valid bool   `json:"valid"` // if storage is valid, the path may disappear somehow
}

func (storage *Storage) SearchFile(filename string, cache ...bool) (apath string, err error) {
	if !storage.Valid {
		return "", errors.New("storage invalid")
	}

	if len(cache) > 0 && cache[0] {
		apath = filepath.Join(storage.CPath, filename)
	} else {
		apath = filepath.Join(storage.APath, filename)
	}

	var fi fs.FileInfo
	if fi, err = os.Stat(apath); err == nil && !fi.IsDir() {
		return apath, err
	}
	if err != nil {
		return "", err
	}
	return "", errors.New("filename points to a directory")
}

func (storage *Storage) SearchCache(cacheId string) (apaths []string, err error) {
	if !storage.Valid {
		return nil, errors.New("storage invalid")
	}
	err = filepath.WalkDir(storage.CPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		filenameParts := strings.Split(filepath.Base(path), ".")
		if len(filenameParts) == 2 && filenameParts[0] == cacheId && regDataRange.MatchString(filenameParts[1]) {
			apaths = append(apaths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return
}
