package fstore

import (
	"os"
	"strings"

	"github.com/shirou/gopsutil/v3/disk"
)

type Mount struct {
	partitions []*Partition
	usage      *Usage
}

// AssignStorage implements IMount.
func (mount *Mount) AssignStorage(apath string) (*Partition, bool) {
	if fi, err := os.Stat(apath); err != nil || !fi.IsDir() {
		return nil, false
	}

	if partition := mount.FindAvailablePartition(apath); partition != nil {
		partition.StoragePath = apath
		return partition, true
	}
	return nil, false
}

// ContainsAvailablePath implements IMount.
func (mount *Mount) ContainsAvailablePath(path string) bool {
	return mount.FindAvailablePartition(path) != nil
}

// ContainsPath implements IMount.
func (mount *Mount) ContainsPath(path string) bool {
	return mount.FindPartition(path) != nil
}

// FindAvailablePartition implements IMount.
func (mount *Mount) FindAvailablePartition(path string) *Partition {
	count := len(mount.partitions)
	for i := 0; i < count; i++ {
		if mount.partitions[i].StoragePath == "" && strings.HasPrefix(path, mount.partitions[i].Mountpoint) {
			return mount.partitions[i]
		}
	}
	return nil
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

// Usage implements IMount.
func (mount *Mount) Usage() *Usage {
	return mount.usage
}

// List implements IMount.
func (mount *Mount) Partitions() []*Partition {
	return mount.partitions
}

// Refresh implements IMount.
func (mount *Mount) Refresh() (iMount IMount, err error) {
	if mount.partitions != nil {
		mount.partitions = nil
	}

	partitionStats, err := disk.Partitions(false)
	if err != nil {
		return mount, err
	}

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

func NewMount() (mount *Mount, iMount IMount, err error) {
	mount = new(Mount)
	iMount, err = mount.Refresh()
	if err != nil {
		return nil, nil, err
	}
	return
}
