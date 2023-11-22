package fstore

import (
	"strings"

	"github.com/shirou/gopsutil/v3/disk"
)

type Mount struct {
	partitions []*Partition
	usage      *Usage
}

// ContainsPath implements IMount.
func (mount *Mount) ContainsPath(path string) bool {
	return mount.FindPartition(path) != nil
}

// FindPartition implements IMount.
func (mount *Mount) FindPartition(path string) *Partition {
	count := len(mount.partitions)
	for i := 0; i < count; i++ {
		if strings.HasPrefix(path, mount.partitions[i].Mountpoint) {
			return mount.partitions[i]
		}
	}
	return nil
}

// Refresh implements IMount.
func (mount *Mount) Refresh() (iMount IMount, err error) {
	partitionStats, err := disk.Partitions(false)
	if err != nil {
		return mount, err
	}

	mount.partitions = make([]*Partition, 0)
	for _, partitionStat := range partitionStats {
		usageStat, err := disk.Usage(partitionStat.Mountpoint)
		if err != nil {
			return mount, err
		}
		mount.partitions = append(mount.partitions, &Partition{
			Mountpoint: partitionStat.Mountpoint,
			FsType:     partitionStat.Fstype,
			Opts:       partitionStat.Opts,
			Usage: Usage{
				Total:       usageStat.Total,
				Free:        usageStat.Free,
				Used:        usageStat.Used,
				UsedPercent: usageStat.UsedPercent,
			},
		})
	}

	mount.usage = &Usage{}
	for _, partition := range mount.partitions {
		mount.usage.Total += partition.Total
		mount.usage.Free += partition.Free
		mount.usage.Used += partition.Used
		mount.usage.UsedPercent += partition.UsedPercent
	}
	mount.usage.UsedPercent /= float64(len(mount.partitions))

	return mount, nil
}

// Usage implements IMount.
func (mount *Mount) Usage() *Usage {
	return mount.usage
}

// Partitions implements IMount.
func (mount *Mount) Partitions() []*Partition {
	return mount.partitions
}

func NewMount() (mount *Mount, iMount IMount, err error) {
	mount = new(Mount)
	iMount, err = mount.Refresh()
	if err != nil {
		return nil, nil, err
	}
	return
}
