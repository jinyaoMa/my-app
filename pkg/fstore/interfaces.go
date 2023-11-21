package fstore

type IFStore interface {
	CreateStorage(apath string) (storage *Storage, ok bool)
}

type IMount interface {
	Refresh() (iMount IMount, err error)
	Partitions() []*Partition
	Usage() *Usage
	FindPartition(path string) *Partition
	FindAvailablePartition(path string) *Partition
	ContainsPath(path string) bool
	ContainsAvailablePath(path string) bool
	AssignStorage(apath string) (*Partition, bool)
}
