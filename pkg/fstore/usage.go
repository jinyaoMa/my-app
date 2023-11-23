package fstore

type Usage struct {
	Total       uint64  `json:"total"`       // total space of partition
	Free        uint64  `json:"free"`        // free space of partition
	Used        uint64  `json:"used"`        // used space of partition
	UsedPercent float64 `json:"usedPercent"` // used percent of partition
}
