package fstore

import (
	"errors"
	"fmt"
	"hash/crc32"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type FStore struct {
	mount             IMount
	options           *Options
	storageMap        map[string]*Storage // uid: storage
	allowedCacheIdMap map[string]bool     // cacheId: active bool
	crc32Table        *crc32.Table
}

// ClearCache implements IFStore.
func (fstore *FStore) ClearCache(uid string, cacheId string, progress func(c int, t int)) (err error) {
	active, ok := fstore.allowedCacheIdMap[cacheId]
	if !ok || !active {
		invalid := fmt.Sprintf("cacheId %s invalid", cacheId)
		return errors.New(invalid)
	}

	storageMap := fstore.GetCurrentStorageMap()
	s, ok := storageMap[uid]
	if !ok || !s.Valid {
		invalid := fmt.Sprintf("uid %s invalid", uid)
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
func (fstore *FStore) Persist(uid string, cacheId string, ext string) (filename string, err error) {
	active, ok := fstore.allowedCacheIdMap[cacheId]
	if !ok || !active {
		invalid := fmt.Sprintf("cacheId %s invalid", cacheId)
		return "", errors.New(invalid)
	}

	storageMap := fstore.GetCurrentStorageMap()
	s, ok := storageMap[uid]
	if !ok || !s.Valid {
		invalid := fmt.Sprintf("uid %s invalid", uid)
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
func (fstore *FStore) FillCache(uid string, cacheId string, rangeStart uint64, rangeEnd uint64, data []byte) (checksum string, err error) {
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
	s, ok := storageMap[uid]
	if !ok || !s.Valid {
		invalid := fmt.Sprintf("uid %s invalid", uid)
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
	maxUID := ""
	maxSize := uint64(0)
	requiredSize := size*2 + fstore.options.ThresholdSize
	for uid, s := range storageMap {
		if s.Valid {
			if s := fstore.mount.GetFreeSize(s.APath); s > requiredSize && s > maxSize {
				maxSize = s
				maxUID = uid
			}
		}
	}
	if maxUID != "" {
		cacheId = uuid.NewString()
		if err := fstore.prepareCache(storageMap[maxUID].CPath, cacheId, size); err != nil {
			return nil, "", err
		}
		fstore.allowedCacheIdMap[cacheId] = true
		return storageMap[maxUID], cacheId, nil
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

	uid := fmt.Sprintf("%x", crc32.ChecksumIEEE([]byte(p.Mountpoint)))
	s, ok := fstore.storageMap[uid]
	if !ok { // add new storage
		if err := fstore.loadCacheIds(cachePath); err != nil {
			return nil, err
		}
		fstore.storageMap[uid] = &Storage{
			Partition: p,
			APath:     apath,
			CPath:     cachePath,
			Valid:     true,
		}
		return fstore.storageMap[uid], nil
	}

	if len(replace) == 0 || !replace[0] {
		return nil, errors.New("do you want to replace the exist storage?")
	}

	if err := fstore.loadCacheIds(cachePath); err != nil {
		return nil, err
	}
	s.Partition = p
	s.APath = apath
	s.CPath = cachePath
	s.Valid = true
	return s, nil
}

// GetCurrentStorageMap implements IFStore.
func (fstore *FStore) GetCurrentStorageMap() map[string]*Storage {
	for uid, _ := range fstore.storageMap {
		if fi, err := os.Stat(fstore.storageMap[uid].APath); err != nil || !fi.IsDir() {
			fstore.storageMap[uid].Valid = false
		}
	}
	return fstore.storageMap
}

func NewFStore(mount IMount, options *Options) (fstore *FStore, iFstore IFStore) {
	fstore = &FStore{
		mount:             mount,
		options:           options,
		storageMap:        make(map[string]*Storage),
		allowedCacheIdMap: make(map[string]bool),
		crc32Table:        crc32.MakeTable(crc32.Castagnoli),
	}
	return fstore, fstore
}
