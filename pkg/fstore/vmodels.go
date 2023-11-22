package fstore

type FStoreOptions struct {
	CacheFolderName string
}

type Storage struct {
	APath string `json:"apath"` // the absolute path in partition to use for storage
	CPath string `json:"cpath"` // the absolute cache path of storage
	Partition
}

type Partition struct {
	Mountpoint string   `json:"mountpoint"`
	FsType     string   `json:"fsType"`
	Opts       []string `json:"opts"`
	Usage
}

type Usage struct {
	Total       uint64  `json:"total"`       // total space of partition
	Free        uint64  `json:"free"`        // free space of partition
	Used        uint64  `json:"used"`        // used space of partition
	UsedPercent float64 `json:"usedPercent"` // used percent of partition
}
