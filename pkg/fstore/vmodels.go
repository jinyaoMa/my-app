package fstore

type FStoreOptions struct {
	CacheFolderName string
}

type Storage struct {
	*Partition
	CachePath string `json:"cachePath"` // the absolute cache path of storage
}

type Partition struct {
	Mountpoint  string   `json:"mountpoint"`
	FsType      string   `json:"fsType"`
	Opts        []string `json:"opts"`
	StoragePath string   `json:"storagePath"` // the absolute path in partition to use for storage
	Usage
}

type Usage struct {
	Total       uint64  `json:"total"`       // total space of partition
	Free        uint64  `json:"free"`        // free space of partition
	Used        uint64  `json:"used"`        // used space of partition
	UsedPercent float64 `json:"usedPercent"` // used percent of partition
}
