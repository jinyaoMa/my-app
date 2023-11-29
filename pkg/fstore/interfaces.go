package fstore

import "os"

type IFileStore interface {
	GetFragmentSize() uint64
	GetCurrentStorageMap() StorageMap
	CreateStorage(apath string, replace ...bool) (storage *Storage, err error)
	RemoveStorage(pid string) (err error)
	SearchFile(filename string, cache ...bool) (apath string, err error)
	SearchAndOpenFile(filename string, flag int, cache ...bool) (file *os.File, err error)
	SearchFileAndGetData(filename string, rangeStart uint64, rangeEnd uint64, cache ...bool) (data []byte, err error)
	PickAStorage(size uint64) (storage *Storage, cacheId string, err error)
	GetUsage() (usage *Usage, err error)
	FillCache(pid string, cacheId string, rangeStart uint64, rangeEnd uint64, data []byte) (checksum string, err error)
	Persist(pid string, cacheId string, ext string) (filename string, err error)
	ClearCache(pid string, cacheId string, progress func(c int, t int)) (err error)
}

type IMount interface {
	FindPartition(apath string) (partition *Partition, err error)
	ContainsPath(apath string) bool
	FindUsage(apath string) (usage *Usage, err error)
	GetFreeSize(apath string) uint64
}
