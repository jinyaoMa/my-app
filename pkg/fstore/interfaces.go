package fstore

type IFStore interface {
	GetCurrentStorages() []*Storage
	CreateStorage(apath string, replace ...bool) (storage *Storage, err error)
	SearchFile(filename string, cache ...bool) (apath string, err error)
	PickAStorage(size uint64) (storage *Storage, err error)
	GetUsage() (usage *Usage, err error)
}

type IMount interface {
	FindPartition(apath string) (partition *Partition, err error)
	ContainsPath(apath string) bool
	FindUsage(apath string) (usage *Usage, err error)
	GetFreeSize(apath string) uint64
}
