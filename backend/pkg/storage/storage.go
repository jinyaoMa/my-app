package storage

import (
	"my-app/backend/pkg/storage/usage"
	"my-app/backend/pkg/utils"
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

// TotalAvailable implements Interface.
func (s *Storage) TotalAvailable() (size uint64) {
	s.forUsage(func(u usage.Interface) {
		size += u.Available()
	})
	return
}

// TotalSize implements Interface.
func (s *Storage) TotalSize() (size uint64) {
	s.forUsage(func(u usage.Interface) {
		size += u.Size()
	})
	return
}

// AddPaths implements Interface
func (s *Storage) AddPaths(paths ...string) (added int) {
	for _, path := range paths {
		if utils.CheckIfDirectoryExists(path) {
			s.paths = append(s.paths, path)
			added += 1
		}
	}
	return
}

func (s *Storage) forUsage(callback func(u usage.Interface)) {
	for _, path := range s.paths {
		u := usage.NewUsage(path)
		callback(u)
	}
}

func New() Interface {
	return &Storage{
		paths: make([]string, 5),
	}
}
