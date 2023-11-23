package fstore

import "github.com/shirou/gopsutil/v3/disk"

type Partition struct {
	Mountpoint string   `json:"mountpoint"`
	Fstype     string   `json:"fstype"`
	Opts       []string `json:"opts"`
}

func (partition *Partition) GetUsage() (usage *Usage, err error) {
	usageStat, err := disk.Usage(partition.Mountpoint)
	if err != nil {
		return nil, err
	}
	return &Usage{
		Total:       usageStat.Total,
		Free:        usageStat.Free,
		Used:        usageStat.Used,
		UsedPercent: usageStat.UsedPercent,
	}, nil
}
