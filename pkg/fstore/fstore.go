package fstore

import (
	"path/filepath"
)

type FStore struct {
	mount   IMount
	options *FStoreOptions
}

// AddStorage implements IFStore.
func (fstore *FStore) CreateStorage(apath string) (storage *Storage, ok bool) {
	if partition, ok := fstore.mount.AssignStorage(apath); ok {
		return &Storage{
			Partition: partition,
			CachePath: filepath.Join(partition.StoragePath, fstore.options.CacheFolderName),
		}, ok
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
