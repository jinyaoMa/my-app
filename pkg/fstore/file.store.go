package fstore

import (
	"errors"
	"fmt"
	"hash/crc32"
	"io"
	"my-app/pkg/base"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

type FileStore struct {
	mount             IMount
	options           *FileStoreOptions
	storageMap        StorageMap      // pid: storage
	allowedCacheIdMap map[string]bool // cacheId: active bool
	crc32Table        *crc32.Table
}

// SearchFileByChecksum implements IFileStore.
func (fileStore *FileStore) SearchFileByChecksum(checksum string, cache ...bool) (apath string, filename string, err error) {
	for _, s := range fileStore.GetCurrentStorageMap() {
		if s.Valid {
			if apath, filename, err = s.SearchFileByChecksum(checksum, cache...); err == nil {
				return apath, filename, nil
			}
		}
	}
	e := fmt.Sprintf("checksum %s not found", checksum)
	return "", "", errors.New(e)
}

// SearchAndCopyFile implements IFileStore.
func (fileStore *FileStore) SearchAndCopyFile(filename string, newExt string, cache ...bool) (err error) {
	apath, err := fileStore.SearchFile(filename, cache...)
	if err != nil {
		return err
	}

	newExt = strings.TrimSpace(newExt)
	oldExt := filepath.Ext(apath)
	if strings.EqualFold(newExt, oldExt) {
		e := fmt.Sprintf("newExt %s is the same as the old one %s", newExt, oldExt)
		return errors.New(e)
	}

	newApath := base.GetFilepathWithoutExtension(apath) + newExt
	if base.IsFileExists(newApath) {
		e := fmt.Sprintf("newExt %s has been used", newExt)
		return errors.New(e)
	}

	file, err := os.Open(apath)
	if err != nil {
		return err
	}
	defer file.Close()
	newFile, err := os.Create(newApath)
	if err != nil {
		return err
	}
	defer newFile.Close()

	buffer := make([]byte, fileStore.options.BufferSize)
	_, err = io.CopyBuffer(newFile, file, buffer)
	return
}

// SearchFileAndGetData implements IFileStore.
func (fileStore *FileStore) SearchFileAndGetData(filename string, rangeStart uint64, rangeEnd uint64, cache ...bool) (data []byte, err error) {
	if rangeEnd-rangeStart == 0 {
		e := fmt.Sprintf("rangeStart %d and rangeEnd %d are the same", rangeStart, rangeEnd)
		return nil, errors.New(e)
	}

	at := int64(rangeStart)
	if at < 0 {
		e := fmt.Sprintf("rangeStart %d error (ReadAt: %d)", rangeStart, at)
		return nil, errors.New(e)
	}

	f, err := fileStore.SearchAndOpenFile(filename, os.O_RDONLY, cache...)
	if err != nil {
		return nil, err
	}

	data = make([]byte, rangeEnd-rangeStart)
	_, err = f.ReadAt(data, at)
	if err != nil && err != io.EOF {
		f.Close()
		return nil, err
	}
	f.Close()
	return
}

// GetFragmentSize implements IFileStore.
func (fileStore *FileStore) GetFragmentSize() uint64 {
	return fileStore.options.FragmentSize
}

// RemoveStorage implements IFileStore.
func (fileStore *FileStore) RemoveStorage(pid string) (err error) {
	_, ok := fileStore.storageMap[pid]
	if !ok {
		e := fmt.Sprintf("pid %s invalid", pid)
		return errors.New(e)
	}

	delete(fileStore.storageMap, pid)
	return
}

// SearchAndOpenFile implements IFileStore.
func (fileStore *FileStore) SearchAndOpenFile(filename string, flag int, cache ...bool) (file *os.File, err error) {
	apath, err := fileStore.SearchFile(filename, cache...)
	if err != nil {
		return nil, err
	}

	return os.OpenFile(apath, flag, os.ModePerm)
}

// ClearCache implements IFileStore.
func (fileStore *FileStore) ClearCache(pid string, cacheId string, progress func(c int, t int)) (err error) {
	active, ok := fileStore.allowedCacheIdMap[cacheId]
	if !ok || !active {
		e := fmt.Sprintf("cacheId %s invalid", cacheId)
		return errors.New(e)
	}

	storageMap := fileStore.GetCurrentStorageMap()
	s, ok := storageMap[pid]
	if !ok || !s.Valid {
		e := fmt.Sprintf("pid %s invalid", pid)
		return errors.New(e)
	}

	apaths, err := s.SearchCache(cacheId)
	if err != nil {
		return err
	}

	total := len(apaths)
	for i, apath := range apaths {
		if err := os.Remove(apath); err != nil {
			return err
		}
		if progress != nil {
			progress(i+1, total)
		}
	}
	delete(fileStore.allowedCacheIdMap, cacheId)
	return
}

// Persist implements IFileStore.
func (fileStore *FileStore) Persist(pid string, cacheId string, ext string) (filename string, err error) {
	active, ok := fileStore.allowedCacheIdMap[cacheId]
	if !ok || !active {
		e := fmt.Sprintf("cacheId %s invalid", cacheId)
		return "", errors.New(e)
	}

	storageMap := fileStore.GetCurrentStorageMap()
	s, ok := storageMap[pid]
	if !ok || !s.Valid {
		e := fmt.Sprintf("pid %s invalid", pid)
		return "", errors.New(e)
	}

	apaths, err := s.SearchCache(cacheId)
	if err != nil {
		return "", err
	}

	cacheFilename := filepath.Join(s.CPath, cacheId)
	filename, err = fileStore.persist(s.APath, ext, cacheFilename, apaths...)
	if err != nil {
		return "", err
	}
	return
}

// FillCache implements IFileStore.
func (fileStore *FileStore) FillCache(pid string, cacheId string, rangeStart uint64, rangeEnd uint64, data []byte) (checksum string, err error) {
	active, ok := fileStore.allowedCacheIdMap[cacheId]
	if !ok || !active {
		e := fmt.Sprintf("cacheId %s invalid", cacheId)
		return "", errors.New(e)
	}
	if rangeEnd-rangeStart != fileStore.options.FragmentSize {
		e := fmt.Sprintf("data range %d-%d invalid", rangeStart, rangeEnd)
		return "", errors.New(e)
	}
	if uint64(len(data)) != fileStore.options.FragmentSize {
		e := fmt.Sprintf("data length must be equal to fragment size %d", fileStore.options.FragmentSize)
		return "", errors.New(e)
	}

	cacheFilename := fileStore.getCacheFilename(cacheId, rangeStart, rangeEnd)
	storageMap := fileStore.GetCurrentStorageMap()
	s, ok := storageMap[pid]
	if !ok || !s.Valid {
		e := fmt.Sprintf("pid %s invalid", pid)
		return "", errors.New(e)
	}

	f, err := os.OpenFile(filepath.Join(s.CPath, cacheFilename), os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return "", err
	}

	crc32New := crc32.New(fileStore.crc32Table)
	return fmt.Sprintf("%x", crc32New.Sum(nil)), nil
}

// GetUsage implements IFileStore.
func (fileStore *FileStore) GetUsage() (usage *Usage, err error) {
	usage = &Usage{}
	count := 0
	for _, s := range fileStore.GetCurrentStorageMap() {
		if s.Valid {
			if u, err := s.GetUsage(); err != nil {
				return nil, err
			} else {
				usage.Total += u.Total
				usage.Free += u.Free
				usage.Used += u.Used
				usage.UsedPercent += u.UsedPercent
				count++
			}
		}
	}
	if count > 0 {
		usage.UsedPercent /= float64(count)
	}
	return
}

// PickAStorage implements IFileStore.
func (fileStore *FileStore) PickAStorage(size uint64) (storage *Storage, cacheId string, err error) {
	storageMap := fileStore.GetCurrentStorageMap()
	maxPID := ""
	maxSize := uint64(0)
	requiredSize := size*2 + fileStore.options.ThresholdSize
	for pid, s := range storageMap {
		if s.Valid {
			if s := fileStore.mount.GetFreeSize(s.APath); s > requiredSize && s > maxSize {
				maxSize = s
				maxPID = pid
			}
		}
	}
	if maxPID != "" {
		cacheId = uuid.NewString()
		if err := fileStore.prepareCache(storageMap[maxPID].CPath, cacheId, size); err != nil {
			return nil, "", err
		}
		fileStore.allowedCacheIdMap[cacheId] = true
		return storageMap[maxPID], cacheId, nil
	}
	e := fmt.Sprintf("no valid storages with space more than %d + %d * 2", fileStore.options.ThresholdSize, size)
	return nil, "", errors.New(e)
}

// SearchFile implements IFileStore.
func (fileStore *FileStore) SearchFile(filename string, cache ...bool) (apath string, err error) {
	for _, s := range fileStore.GetCurrentStorageMap() {
		if s.Valid {
			if apath, err = s.SearchFile(filename, cache...); err == nil {
				return apath, nil
			}
		}
	}
	e := fmt.Sprintf("file %s not found", filename)
	return "", errors.New(e)
}

// CreateStorage implements IFileStore.
func (fileStore *FileStore) CreateStorage(apath string, replace ...bool) (storage *Storage, err error) {
	cachePath := filepath.Join(apath, fileStore.options.CacheFolderName)
	if err := os.MkdirAll(cachePath, os.ModeDir); err != nil {
		return nil, err
	}

	p, err := fileStore.mount.FindPartition(apath)
	if err != nil {
		e := fmt.Sprintf("partition for %s not exist", apath)
		return nil, errors.New(e)
	}

	pid := fmt.Sprintf("%x", crc32.Checksum([]byte(p.Mountpoint), fileStore.crc32Table))
	s, ok := fileStore.storageMap[pid]
	if !ok { // add new storage
		if err := fileStore.loadCacheIds(cachePath); err != nil {
			return nil, err
		}
		fileStore.storageMap[pid] = &Storage{
			Partition: p,
			PID:       pid,
			APath:     apath,
			CPath:     cachePath,
			Valid:     true,
		}
		return fileStore.storageMap[pid], nil
	}

	if len(replace) == 0 || !replace[0] {
		e := fmt.Sprintf("do you want to replace the exist storage (%s)?", s.APath)
		return nil, errors.New(e)
	}

	if err := fileStore.loadCacheIds(cachePath); err != nil {
		return nil, err
	}
	s.Partition = p
	s.PID = pid
	s.APath = apath
	s.CPath = cachePath
	s.Valid = true
	return s, nil
}

// GetCurrentStorageMap implements IFileStore.
func (fileStore *FileStore) GetCurrentStorageMap() StorageMap {
	for pid := range fileStore.storageMap {
		fileStore.storageMap[pid].Valid = base.IsDirectoryExists(fileStore.storageMap[pid].APath)
	}
	return fileStore.storageMap
}

func NewFileStore(mount IMount, options *FileStoreOptions) (fileStore *FileStore, iFilestore IFileStore, err error) {
	options, err = NewFileStoreOptions(options)
	if err != nil {
		return nil, nil, err
	}

	fileStore = &FileStore{
		mount:             mount,
		options:           options,
		storageMap:        make(StorageMap),
		allowedCacheIdMap: make(map[string]bool),
		crc32Table:        crc32.MakeTable(crc32.Castagnoli),
	}
	return fileStore, fileStore, nil
}
