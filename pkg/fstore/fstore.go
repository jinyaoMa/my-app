package fstore

import (
	"errors"
	"fmt"
	"hash/crc32"
	"io"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type FStore struct {
	mount             IMount
	options           *Options
	storageMap        StorageMap      // pid: storage
	allowedCacheIdMap map[string]bool // cacheId: active bool
	crc32Table        *crc32.Table
}

// SearchFileAndGetData implements IFStore.
func (fstore *FStore) SearchFileAndGetData(filename string, rangeStart uint64, rangeEnd uint64, cache ...bool) (data []byte, err error) {
	if rangeEnd-rangeStart == 0 {
		e := fmt.Sprintf("rangeStart %d and rangeEnd %d are the same", rangeStart, rangeEnd)
		return nil, errors.New(e)
	}

	at := int64(rangeStart)
	if at < 0 {
		e := fmt.Sprintf("rangeStart %d error (ReadAt: %d)", rangeStart, at)
		return nil, errors.New(e)
	}

	f, err := fstore.SearchAndOpenFile(filename, os.O_RDONLY, cache...)
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

// GetFragmentSize implements IFStore.
func (fstore *FStore) GetFragmentSize() uint64 {
	return fstore.options.FragmentSize
}

// RemoveStorage implements IFStore.
func (fstore *FStore) RemoveStorage(pid string) (err error) {
	_, ok := fstore.storageMap[pid]
	if !ok {
		invalid := fmt.Sprintf("pid %s invalid", pid)
		return errors.New(invalid)
	}

	delete(fstore.storageMap, pid)
	return
}

// SearchAndOpenFile implements IFStore.
func (fstore *FStore) SearchAndOpenFile(filename string, flag int, cache ...bool) (file *os.File, err error) {
	apath, err := fstore.SearchFile(filename, cache...)
	if err != nil {
		return nil, err
	}

	return os.OpenFile(apath, flag, os.ModePerm)
}

// ClearCache implements IFStore.
func (fstore *FStore) ClearCache(pid string, cacheId string, progress func(c int, t int)) (err error) {
	active, ok := fstore.allowedCacheIdMap[cacheId]
	if !ok || !active {
		invalid := fmt.Sprintf("cacheId %s invalid", cacheId)
		return errors.New(invalid)
	}

	storageMap := fstore.GetCurrentStorageMap()
	s, ok := storageMap[pid]
	if !ok || !s.Valid {
		invalid := fmt.Sprintf("pid %s invalid", pid)
		return errors.New(invalid)
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
	delete(fstore.allowedCacheIdMap, cacheId)
	return
}

// Persist implements IFStore.
func (fstore *FStore) Persist(pid string, cacheId string, ext string) (filename string, err error) {
	active, ok := fstore.allowedCacheIdMap[cacheId]
	if !ok || !active {
		invalid := fmt.Sprintf("cacheId %s invalid", cacheId)
		return "", errors.New(invalid)
	}

	storageMap := fstore.GetCurrentStorageMap()
	s, ok := storageMap[pid]
	if !ok || !s.Valid {
		invalid := fmt.Sprintf("pid %s invalid", pid)
		return "", errors.New(invalid)
	}

	apaths, err := s.SearchCache(cacheId)
	if err != nil {
		return "", err
	}

	cacheFilename := filepath.Join(s.CPath, cacheId)
	filename, err = fstore.persist(s.APath, ext, cacheFilename, apaths...)
	if err != nil {
		return "", err
	}
	return
}

// FillCache implements IFStore.
func (fstore *FStore) FillCache(pid string, cacheId string, rangeStart uint64, rangeEnd uint64, data []byte) (checksum string, err error) {
	active, ok := fstore.allowedCacheIdMap[cacheId]
	if !ok || !active {
		invalid := fmt.Sprintf("cacheId %s invalid", cacheId)
		return "", errors.New(invalid)
	}
	if rangeEnd-rangeStart != fstore.options.FragmentSize {
		invalid := fmt.Sprintf("data range %d-%d invalid", rangeStart, rangeEnd)
		return "", errors.New(invalid)
	}
	if uint64(len(data)) != fstore.options.FragmentSize {
		invalid := fmt.Sprintf("data length must be equal to fragment size %d", fstore.options.FragmentSize)
		return "", errors.New(invalid)
	}

	cacheFilename := fstore.getCacheFilename(cacheId, rangeStart, rangeEnd)
	storageMap := fstore.GetCurrentStorageMap()
	s, ok := storageMap[pid]
	if !ok || !s.Valid {
		invalid := fmt.Sprintf("pid %s invalid", pid)
		return "", errors.New(invalid)
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

	crc32New := crc32.New(fstore.crc32Table)
	return fmt.Sprintf("%x", crc32New.Sum(nil)), nil
}

// GetUsage implements IFStore.
func (fstore *FStore) GetUsage() (usage *Usage, err error) {
	usage = &Usage{}
	count := 0
	for _, s := range fstore.GetCurrentStorageMap() {
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

// PickAStorage implements IFStore.
func (fstore *FStore) PickAStorage(size uint64) (storage *Storage, cacheId string, err error) {
	storageMap := fstore.GetCurrentStorageMap()
	maxPID := ""
	maxSize := uint64(0)
	requiredSize := size*2 + fstore.options.ThresholdSize
	for pid, s := range storageMap {
		if s.Valid {
			if s := fstore.mount.GetFreeSize(s.APath); s > requiredSize && s > maxSize {
				maxSize = s
				maxPID = pid
			}
		}
	}
	if maxPID != "" {
		cacheId = uuid.NewString()
		if err := fstore.prepareCache(storageMap[maxPID].CPath, cacheId, size); err != nil {
			return nil, "", err
		}
		fstore.allowedCacheIdMap[cacheId] = true
		return storageMap[maxPID], cacheId, nil
	}
	return nil, "", errors.New("no valid storages")
}

// SearchFile implements IFStore.
func (fstore *FStore) SearchFile(filename string, cache ...bool) (apath string, err error) {
	for _, s := range fstore.GetCurrentStorageMap() {
		if s.Valid {
			if apath, err = s.SearchFile(filename, cache...); err == nil {
				return apath, nil
			}
		}
	}
	notFound := fmt.Sprintf("file %s not found", filename)
	return "", errors.New(notFound)
}

// CreateStorage implements IFStore.
func (fstore *FStore) CreateStorage(apath string, replace ...bool) (storage *Storage, err error) {
	cachePath := filepath.Join(apath, fstore.options.CacheFolderName)
	if err := os.MkdirAll(cachePath, os.ModeDir); err != nil {
		return nil, err
	}

	p, err := fstore.mount.FindPartition(apath)
	if err != nil {
		notExist := fmt.Sprintf("partition for %s not exist", apath)
		return nil, errors.New(notExist)
	}

	pid := fmt.Sprintf("%x", crc32.Checksum([]byte(p.Mountpoint), fstore.crc32Table))
	s, ok := fstore.storageMap[pid]
	if !ok { // add new storage
		if err := fstore.loadCacheIds(cachePath); err != nil {
			return nil, err
		}
		fstore.storageMap[pid] = &Storage{
			Partition: p,
			PID:       pid,
			APath:     apath,
			CPath:     cachePath,
			Valid:     true,
		}
		return fstore.storageMap[pid], nil
	}

	if len(replace) == 0 || !replace[0] {
		return nil, errors.New("do you want to replace the exist storage?")
	}

	if err := fstore.loadCacheIds(cachePath); err != nil {
		return nil, err
	}
	s.Partition = p
	s.PID = pid
	s.APath = apath
	s.CPath = cachePath
	s.Valid = true
	return s, nil
}

// GetCurrentStorageMap implements IFStore.
func (fstore *FStore) GetCurrentStorageMap() StorageMap {
	for pid := range fstore.storageMap {
		if fi, err := os.Stat(fstore.storageMap[pid].APath); err != nil || !fi.IsDir() {
			fstore.storageMap[pid].Valid = false
		}
	}
	return fstore.storageMap
}

func NewFStore(mount IMount, options *Options) (fstore *FStore, iFstore IFStore, err error) {
	options, err = NewOptions(options)
	if err != nil {
		return nil, nil, err
	}

	fstore = &FStore{
		mount:             mount,
		options:           options,
		storageMap:        make(StorageMap),
		allowedCacheIdMap: make(map[string]bool),
		crc32Table:        crc32.MakeTable(crc32.Castagnoli),
	}
	return fstore, fstore, nil
}
