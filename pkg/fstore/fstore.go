package fstore

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type FStore struct {
	mount    IMount
	options  *Options
	storages []*Storage
}

// GetUsage implements IFStore.
func (fstore *FStore) GetUsage() (usage *Usage, err error) {
	usage = &Usage{}
	count := 0
	for _, storage := range fstore.GetCurrentStorages() {
		if u, err := storage.GetUsage(); err != nil {
			return nil, err
		} else {
			usage.Total += u.Total
			usage.Free += u.Free
			usage.Used += u.Used
			usage.UsedPercent += u.UsedPercent
			count++
		}
	}
	if count > 0 {
		usage.UsedPercent /= float64(count)
	}
	return
}

// PickAStorage implements IFStore.
func (fstore *FStore) PickAStorage(size uint64) (storage *Storage, err error) {
	storages := fstore.GetCurrentStorages()
	count := len(storages)
	maxIndex := -1
	maxSize := uint64(0)
	for i := 0; i < count; i++ {
		if storages[i].Valid {
			if s := fstore.mount.GetFreeSize(storages[i].APath); s > size && s > maxSize {
				maxSize = s
				maxIndex = i
			}
		}
	}
	if maxIndex >= 0 {
		return storages[maxIndex], nil
	}
	return nil, errors.New("no valid storages")
}

// SearchFile implements IFStore.
func (fstore *FStore) SearchFile(filename string, cache ...bool) (apath string, err error) {
	for _, storage := range fstore.GetCurrentStorages() {
		if apath, err = storage.SearchFile(filename, cache...); err == nil {
			return
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

	if p, err := fstore.mount.FindPartition(apath); err == nil {
		if i := slices.IndexFunc(fstore.storages, func(s *Storage) bool {
			return strings.HasPrefix(s.APath, p.Mountpoint)
		}); i >= 0 {
			if len(replace) > 0 && replace[0] {
				fstore.storages[i].Partition = p
				fstore.storages[i].APath = apath
				fstore.storages[i].CPath = cachePath
				fstore.storages[i].Valid = true
				return fstore.storages[i], nil
			} else {
				return nil, errors.New("do you want to replace the exist storage?")
			}
		} else {
			storage = &Storage{
				Partition: p,
				APath:     apath,
				CPath:     cachePath,
				Valid:     true,
			}
			fstore.storages = append(fstore.storages, storage)
			return storage, nil
		}
	}
	return nil, errors.New("partition not exist")
}

// GetCurrentStorages implements IFStore.
func (fstore *FStore) GetCurrentStorages() []*Storage {
	count := len(fstore.storages)
	for i := 0; i < count; i++ {
		if fi, err := os.Stat(fstore.storages[i].APath); err != nil || !fi.IsDir() {
			fstore.storages[i].Valid = false
		}
	}
	return fstore.storages
}

func NewFStore(mount IMount, options *Options) (fstore *FStore, iFstore IFStore) {
	fstore = &FStore{
		mount:    mount,
		options:  options,
		storages: make([]*Storage, 0),
	}
	return fstore, fstore
}
