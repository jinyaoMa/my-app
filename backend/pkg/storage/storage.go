package storage

import (
	"errors"
	"my-app/backend/pkg/utils"
	"strings"
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

// GetMountpointUsage implements Interface.
func (s *Storage) GetMountpointUsage() (u MountpointUsage, err error) {
	return NewMountpointUsage(s.paths...)
}

// AddPaths implements Interface
func (s *Storage) AddPaths(paths ...string) (added int, err error) {
	var u MountpointUsage
	u, err = s.GetMountpointUsage()
	if err != nil {
		return
	}

	for mountpoint, mStat := range u {
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

func New() Interface {
	return &Storage{
		paths: make([]string, 5),
	}
}
