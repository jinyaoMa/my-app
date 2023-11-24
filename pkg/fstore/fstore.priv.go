package fstore

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

// cache fragment filename format `{cacheId}.{rangeStartIndex:inclusive}-{rangeEndIndex:exclusive}`
func (fstore *FStore) prepareCache(storage *Storage, cacheId string, size uint64) (err error) {
	fragmentBuffer := make([]byte, fstore.options.FragmentSize)
	for i := uint64(0); size > 0; {
		fragmentFilename := ""
		if size > fstore.options.FragmentSize {
			fragmentFilename = fstore.getCacheFilename(cacheId, i, i+fstore.options.FragmentSize)
		} else {
			fragmentFilename = fstore.getCacheFilename(cacheId, i, i+size)
			fragmentBuffer = make([]byte, size)
		}

		err = os.WriteFile(filepath.Join(storage.CPath, fragmentFilename), fragmentBuffer, os.ModePerm)
		if err != nil {
			return
		}

		size -= fstore.options.FragmentSize
		i += fstore.options.FragmentSize
	}
	return
}

func (fstore *FStore) getCacheFilename(cacheId string, rangeStart uint64, rangeEnd uint64) string {
	return fmt.Sprintf("%s.%d-%d", cacheId, rangeStart, rangeEnd)
}

func (fstore *FStore) loadCacheIds(cpath string) (err error) {
	err = filepath.WalkDir(cpath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		filenameParts := strings.Split(filepath.Base(path), ".")
		if len(filenameParts) == 2 {
			if _, err := uuid.Parse(filenameParts[0]); err == nil {
				fstore.allowedCacheIdMap[filenameParts[0]] = true
			}
		}
		return nil
	})
	return err
}
