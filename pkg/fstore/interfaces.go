package fstore

type IFStore interface {
	CreateStorage(apath string) (storage *Storage, ok bool)
}

type IMount interface {
	Partitions() []*Partition
	Usage() *Usage
	Refresh() (iMount IMount, err error)
	FindPartition(path string) *Partition
	ContainsPath(path string) bool
}
