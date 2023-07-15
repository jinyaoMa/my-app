package storage

import "github.com/shirou/gopsutil/v3/disk"

type MountpointStat struct {
	disk.PartitionStat
	UsedPath string // storage path to use in this mountpoint
}
