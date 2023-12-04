package fstore

import (
	"errors"
	"fmt"
	"io/fs"
	"my-app/pkg/base"
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
		e := fmt.Sprintf("storage %s invalid", storage.APath)
		return "", errors.New(e)
	}

	if len(cache) > 0 && cache[0] {
		apath = filepath.Join(storage.CPath, filename)
	} else {
		apath = filepath.Join(storage.APath, filename)
	}

	if !base.IsFileExists(apath) {
		e := fmt.Sprintf("file %s not exists", apath)
		return "", errors.New(e)
	}
	return
}

func (storage *Storage) SearchFileByChecksum(checksum string, cache ...bool) (apath string, filename string, err error) {
	if !storage.Valid {
		e := fmt.Sprintf("storage %s invalid", storage.APath)
		return "", "", errors.New(e)
	}

	target := storage.APath
	if len(cache) > 0 && cache[0] {
		target = storage.CPath
	}

	err = filepath.WalkDir(target, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if target == path {
			return nil
		}
		if d.IsDir() {
			return filepath.SkipDir
		}
		if strings.HasPrefix(d.Name(), checksum) {
			apath = path
			filename = d.Name()
			return filepath.SkipAll
		}
		return nil
	})
	return
}

func (storage *Storage) SearchCache(cacheId string) (apaths []string, err error) {
	if !storage.Valid {
		e := fmt.Sprintf("storage %s invalid", storage.APath)
		return nil, errors.New(e)
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
