package storage

import (
	"strings"

	"github.com/shirou/gopsutil/v3/disk"
)

type MountpointStat struct {
	PartitionStat *disk.PartitionStat
	UsageStat     *disk.UsageStat
	UsedPath      string // storage path to use in this mountpoint
}

type MountpointUsage map[string]*MountpointStat

func (u MountpointUsage) AvailableMountPoints() (mountpoints []string) {
	for mountpoint, mStat := range u {
		if mStat.UsedPath != "" {
			mountpoints = append(mountpoints, mountpoint)
		}
	}
	return
}

func (u MountpointUsage) TotalSize() (total uint64) {
	for _, mStat := range u {
		total += mStat.UsageStat.Total
	}
	return
}

func (u MountpointUsage) TotalFree() (free uint64) {
	for _, mStat := range u {
		free += mStat.UsageStat.Free
	}
	return
}

func (u MountpointUsage) TotalUsed() (used uint64) {
	for _, mStat := range u {
		used += mStat.UsageStat.Used
	}
	return
}

func (u MountpointUsage) TotalUsedPercent() (usedPercent float64) {
	usedPercent = float64(u.TotalUsed()) / float64(u.TotalSize()) * 100
	return
}

func NewMountpointUsage(paths ...string) (u MountpointUsage, err error) {
	var pStats []disk.PartitionStat
	pStats, err = disk.Partitions(false)
	if err != nil {
		return
	}

	u = make(MountpointUsage)
	for _, pStat := range pStats {
		var uStats *disk.UsageStat
		uStats, err = disk.Usage(pStat.Mountpoint)
		if err != nil {
			return nil, err
		}

		u[pStat.Mountpoint] = &MountpointStat{
			PartitionStat: &pStat,
			UsageStat:     uStats,
		}

		for _, path := range paths {
			if strings.HasPrefix(path, pStat.Mountpoint) {
				u[pStat.Mountpoint].UsedPath = path
			}
		}
	}
	return
}
