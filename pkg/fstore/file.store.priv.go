package fstore

import (
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/zeebo/xxh3"
)

func (fileStore *FileStore) prepareCache(cpath string, cacheId string, size uint64) (err error) {
	fragmentBuffer := make([]byte, fileStore.options.FragmentSize)
	for i := uint64(0); size > 0; {
		fragmentFilename := ""
		if size > fileStore.options.FragmentSize {
			fragmentFilename = fileStore.getCacheFilename(cacheId, i, i+fileStore.options.FragmentSize)
		} else {
			fragmentFilename = fileStore.getCacheFilename(cacheId, i, i+size)
			fragmentBuffer = make([]byte, size)
		}

		err = os.WriteFile(filepath.Join(cpath, fragmentFilename), fragmentBuffer, os.ModePerm)
		if err != nil {
			return err
		}

		size -= fileStore.options.FragmentSize
		i += fileStore.options.FragmentSize
	}
	return
}

// cache fragment filename format `{cacheId}.{rangeStartIndex:inclusive}-{rangeEndIndex:exclusive}`
func (fileStore *FileStore) getCacheFilename(cacheId string, rangeStart uint64, rangeEnd uint64) string {
	return fmt.Sprintf("%s.%d-%d", cacheId, rangeStart, rangeEnd)
}

func (fileStore *FileStore) loadCacheIds(cpath string) (err error) {
	err = filepath.WalkDir(cpath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		filenameParts := strings.Split(filepath.Base(path), ".")
		if len(filenameParts) == 2 {
			if _, err := uuid.Parse(filenameParts[0]); err == nil {
				fileStore.allowedCacheIdMap[filenameParts[0]] = true
			}
		}
		return nil
	})
	return err
}

// persist checksum format `{sha1:160bit:40hex}{xxh3:128bit:32hex}{size:64bit:16hex}`
func (fileStore *FileStore) checksum(loading func(buffer []byte) error, apaths ...string) (sum string, err error) {
	if len(apaths) == 0 {
		return "", nil
	}

	buffer := make([]byte, fileStore.options.BufferSize)
	size := uint64(0)
	sha1New := sha1.New()
	xxh3New := xxh3.New()
	for _, apath := range apaths {
		f, err := os.Open(apath)
		if err != nil {
			return "", err
		}

		for {
			n, err := f.Read(buffer)
			if err == io.EOF {
				break
			}
			if err != nil {
				f.Close()
				return "", err
			}

			temp := buffer[0:n]
			_, err = sha1New.Write(temp)
			if err != nil {
				f.Close()
				return "", err
			}
			_, err = xxh3New.Write(temp)
			if err != nil {
				f.Close()
				return "", err
			}
			size += uint64(n)
			if loading != nil {
				if err = loading(temp); err != nil {
					return "", err
				}
			}
		}
		f.Close()
	}

	bsize := make([]byte, 8)
	binary.BigEndian.PutUint64(bsize, size)
	return fmt.Sprintf("%x%x%x", sha1New.Sum(nil), xxh3New.Sum128().Bytes(), bsize), nil
}

// persist checksum format `{sha1:160bit:40hex}{xxh3:128bit:32hex}{size:64bit:16hex}`
// ext => ".txt", ".go"...
func (fileStore *FileStore) persist(desDir string, ext string, tmpFilename string, apaths ...string) (filename string, err error) {
	tmp, err := os.Create(tmpFilename)
	if err != nil {
		return "", err
	}
	defer tmp.Close()

	sum, err := fileStore.checksum(func(buffer []byte) error {
		_, err := tmp.Write(buffer)
		return err
	}, apaths...)
	if err != nil {
		return "", err
	}
	filename = sum + ext

	tmp.Close()
	err = os.Rename(tmpFilename, filepath.Join(desDir, filename))
	if err != nil {
		return "", err
	}
	return
}
