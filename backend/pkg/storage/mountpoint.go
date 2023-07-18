package storage

import (
	"strings"

	"github.com/shirou/gopsutil/v3/disk"
)

type MountpointStat struct {
	PartitionStat *disk.PartitionStat
	UsageStat     *disk.UsageStat
	StoragePath   *StoragePath // storage path to use in this mountpoint
}

type MountpointUsage map[string]*MountpointStat

func (u MountpointUsage) PickAPath(needSize uint64) (path *StoragePath) {
	for _, mStat := range u {
		if mStat.StoragePath != nil && mStat.UsageStat.Free > needSize {
			return mStat.StoragePath
		}
	}
	return
}

func (u MountpointUsage) AvailableMountPoints() (mountpoints []string) {
	for mountpoint, mStat := range u {
		if mStat.StoragePath != nil {
			mountpoints = append(mountpoints, mountpoint)
		}
	}
	return
}

func (u MountpointUsage) TotalSize() (total uint64) {
	for _, mStat := range u {
		if mStat.StoragePath != nil {
			total += mStat.UsageStat.Total
		}
	}
	return
}

func (u MountpointUsage) TotalFree() (free uint64) {
	for _, mStat := range u {
		if mStat.StoragePath != nil {
			free += mStat.UsageStat.Free
		}
	}
	return
}

func (u MountpointUsage) TotalUsed() (used uint64) {
	for _, mStat := range u {
		if mStat.StoragePath != nil {
			used += mStat.UsageStat.Used
		}
	}
	return
}

func (u MountpointUsage) TotalUsedPercent() (usedPercent float64) {
	if u.TotalSize() == 0 {
		return 0
	}
	usedPercent = float64(u.TotalUsed()) / float64(u.TotalSize()) * 100
	return
}

func NewMountpointUsage(paths []*StoragePath) (u MountpointUsage, err error) {
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
			if strings.HasPrefix(path.Dir, pStat.Mountpoint) {
				u[pStat.Mountpoint].StoragePath = path
			}
		}
	}
	return
}
