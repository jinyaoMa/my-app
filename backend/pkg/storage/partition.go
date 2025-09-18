package storage

type Partition struct {
	Device      string   `json:"device"`
	Mountpoint  string   `json:"mountpoint"`
	Fstype      string   `json:"fstype"`
	Opts        []string `json:"opts"`
	Total       uint64   `json:"total"`
	Free        uint64   `json:"free"`
	Used        uint64   `json:"used"`
	UsedPercent float64  `json:"usedPercent"`
}
