package fstore

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"hash/crc64"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"github.com/zeebo/xxh3"
)

var (
	regDataRange = regexp.MustCompile(`^\d+\-\d+$`)
)

func (fstore *FStore) prepareCache(cpath string, cacheId string, size uint64) (err error) {
	fragmentBuffer := make([]byte, fstore.options.FragmentSize)
	for i := uint64(0); size > 0; {
		fragmentFilename := ""
		if size > fstore.options.FragmentSize {
			fragmentFilename = fstore.getCacheFilename(cacheId, i, i+fstore.options.FragmentSize)
		} else {
			fragmentFilename = fstore.getCacheFilename(cacheId, i, i+size)
			fragmentBuffer = make([]byte, size)
		}

		err = os.WriteFile(filepath.Join(cpath, fragmentFilename), fragmentBuffer, os.ModePerm)
		if err != nil {
			return err
		}

		size -= fstore.options.FragmentSize
		i += fstore.options.FragmentSize
	}
	return
}

// cache fragment filename format `{cacheId}.{rangeStartIndex:inclusive}-{rangeEndIndex:exclusive}`
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

// persist checksum format `{sha1:160bit:40hex}-{size:64bit:16hex}-{xxh3:128bit:32hex}-{crc64:64bit:16hex}`
func (fstore *FStore) checksum(apaths ...string) (sum string, err error) {
	if len(apaths) == 0 {
		return "", errors.New("no apaths")
	}

	buffer := make([]byte, fstore.options.BufferSize)
	size := uint64(0)
	sha1New := sha1.New()
	xxhNew := xxh3.New()
	crc64New := crc64.New(fstore.crc64Table)
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
			size += uint64(n)
			_, err = sha1New.Write(temp)
			if err != nil {
				f.Close()
				return "", err
			}
			_, err = xxhNew.Write(temp)
			if err != nil {
				f.Close()
				return "", err
			}
			_, err = crc64New.Write(temp)
			if err != nil {
				f.Close()
				return "", err
			}
		}
		f.Close()
	}
	return fmt.Sprintf("%x-%x-%x-%x", sha1New.Sum(nil), size, xxhNew.Sum128().Bytes(), crc64New.Sum(nil)), nil
}
