package fstore

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
)

type Storage struct {
	*Partition
	UID   string `json:"uid"`   // id for user uploading files to valid location
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
		return
	}
	if err != nil {
		return "", err
	}
	return "", errors.New("filename points to a directory")
}
