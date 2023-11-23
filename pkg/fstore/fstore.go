package fstore

import (
	"os"
	"path/filepath"
	"slices"
)

type FStore struct {
	options  *FStoreOptions
	storages []*Storage
}

// CreateStorage implements IFStore.
func (fstore *FStore) CreateStorage(mount IMount, apath string) (storage *Storage, ok bool) {
	if fi, err := os.Stat(apath); err != nil || !fi.IsDir() {
		return nil, false
	}
	if partition := mount.FindPartition(apath); partition != nil && !slices.ContainsFunc(fstore.storages, func(s *Storage) bool {
		return s.Mountpoint == partition.Mountpoint
	}) {
		storage = &Storage{
			Partition: *partition,
			APath:     apath,
			CPath:     filepath.Join(apath, fstore.options.CacheFolderName),
		}
		fstore.storages = append(fstore.storages, storage)
		return storage, true
	}
	return nil, false
}

// Storages implements IFStore.
func (fstore *FStore) Storages() []*Storage {
	return fstore.storages
}

func NewFStore(options *FStoreOptions) (fstore *FStore, iFstore IFStore) {
	fstore = &FStore{
		options:  options,
		storages: make([]*Storage, 0),
	}
	return fstore, fstore
}
