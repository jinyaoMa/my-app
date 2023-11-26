package fstore

type IFStore interface {
	GetCurrentStorageMap() map[string]*Storage
	CreateStorage(apath string, replace ...bool) (storage *Storage, err error)
	SearchFile(filename string, cache ...bool) (apath string, err error)
	PickAStorage(size uint64) (storage *Storage, cacheId string, err error)
	GetUsage() (usage *Usage, err error)
	FillCache(uid string, cacheId string, rangeStart uint64, rangeEnd uint64, data []byte) (checksum string, err error)
	Persist(uid string, cacheId string, ext string) (filename string, err error)
	ClearCache(uid string, cacheId string, progress func(c int, t int)) (err error)
}

type IMount interface {
	FindPartition(apath string) (partition *Partition, err error)
	ContainsPath(apath string) bool
	FindUsage(apath string) (usage *Usage, err error)
	GetFreeSize(apath string) uint64
}
