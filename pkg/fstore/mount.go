package fstore

import (
	"errors"
	"fmt"
	"strings"

	"github.com/shirou/gopsutil/v3/disk"
)

type Mount struct{}

// GetFreeSize implements IMount.
func (mount *Mount) GetFreeSize(apath string) uint64 {
	if u, err := mount.FindUsage(apath); err == nil {
		return u.Free
	}
	return 0
}

// FindUsage implements IMount.
func (*Mount) FindUsage(apath string) (usage *Usage, err error) {
	uStat, err := disk.Usage(apath)
	if err != nil {
		return nil, err
	}
	return &Usage{
		Total:       uStat.Total,
		Free:        uStat.Free,
		Used:        uStat.Used,
		UsedPercent: uStat.UsedPercent,
	}, nil
}

// ContainsPath implements IMount.
func (mount *Mount) ContainsPath(apath string) bool {
	_, err := mount.FindPartition(apath)
	return err == nil
}

// FindPartition implements IMount.
func (*Mount) FindPartition(apath string) (partition *Partition, err error) {
	pStats, err := disk.Partitions(false)
	if err != nil {
		return nil, err
	}
	for _, pStat := range pStats {
		if strings.HasPrefix(apath, pStat.Mountpoint) {
			return &Partition{
				Mountpoint: pStat.Mountpoint,
				Fstype:     pStat.Fstype,
				Opts:       pStat.Opts,
			}, nil
		}
	}
	e := fmt.Sprintf("partition of path %s not found", apath)
	return nil, errors.New(e)
}

func NewMount() (mount *Mount, iMount IMount) {
	mount = new(Mount)
	return mount, mount
}
