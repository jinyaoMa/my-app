package fstore

import (
	"path/filepath"
	"slices"
)

type FStore struct {
	mount    IMount
	options  *FStoreOptions
	storages []*Storage
}

// CreateStorage implements IFStore.
func (fstore *FStore) CreateStorage(apath string) (storage *Storage, ok bool) {
	if _, err := fstore.mount.Refresh(); err != nil {
		return nil, false
	}
	if partition := fstore.mount.FindPartition(apath); partition != nil && !slices.ContainsFunc(fstore.storages, func(s *Storage) bool {
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

func NewFStore(mount IMount, options *FStoreOptions) (fstore *FStore, iFstore IFStore) {
	fstore = &FStore{
		mount:   mount,
		options: options,
	}
	return fstore, fstore
}
