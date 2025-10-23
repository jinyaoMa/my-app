package storage

type Partition struct {
	Device      string   `json:"device"`
	Mountpoint  string   `json:"mountpoint"`
	Fstype      string   `json:"fstype"`
	Opts        []string `json:"opts"`
	Total       int64    `json:"total"`
	Free        int64    `json:"free"`
	Used        int64    `json:"used"`
	UsedPercent float64  `json:"usedPercent"`
}
