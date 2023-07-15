package storage

import (
	"errors"
	"my-app/backend/pkg/utils"
	"strings"

	"github.com/shirou/gopsutil/v3/disk"
)

const (
	B  uint64 = 1
	KB        = 1024 * B
	MB        = 1024 * KB
	GB        = 1024 * MB
	TB        = 1024 * GB
)

type Storage struct {
	paths []string
}

// AvailableMountPoints implements Interface.
func (s *Storage) AvailableMountPoints() (mountpoints []string, err error) {
	var usage map[string]*MountpointStat
	usage, err = s.MountpointUsage()
	if err != nil {
		return
	}

	for mountpoint, mStat := range usage {
		if mStat.UsedPath == "" {
			mountpoints = append(mountpoints, mountpoint)
		}
	}
	return
}

// TotalAvailable implements Interface.
func (s *Storage) TotalAvailable() (size uint64) {
	return
}

// TotalSize implements Interface.
func (s *Storage) TotalSize() (size uint64) {
	return
}

// AddPaths implements Interface
func (s *Storage) AddPaths(paths ...string) (added int, err error) {
	var usage map[string]*MountpointStat
	usage, err = s.MountpointUsage()
	if err != nil {
		return
	}

	for mountpoint, mStat := range usage {
		for _, path := range paths {
			if mStat.UsedPath == "" &&
				strings.HasPrefix(path, mountpoint) &&
				utils.CheckIfDirectoryExists(path) {
				s.paths = append(s.paths, path)
				mStat.UsedPath = path
				added += 1
			}
		}
	}

	if added != len(paths) {
		err = errors.New("some paths cannot be added due to occupied mountpoints")
		return
	}

	return
}

// MountPointUsage implements Interface.
func (s *Storage) MountpointUsage() (usage map[string]*MountpointStat, err error) {
	var pStats []disk.PartitionStat
	pStats, err = disk.Partitions(false)
	if err != nil {
		return
	}

	usage = make(map[string]*MountpointStat)
	for _, pStat := range pStats {
		usage[pStat.Mountpoint] = &MountpointStat{
			PartitionStat: pStat,
		}

		for _, path := range s.paths {
			if strings.HasPrefix(path, pStat.Mountpoint) {
				usage[pStat.Mountpoint].UsedPath = path
			}
		}
	}
	return
}

func New() Interface {
	return &Storage{
		paths: make([]string, 5),
	}
}
