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
	paths map[string]usage.Interface
}

// TotalAvailable implements Interface.
func (s *Storage) TotalAvailable() (size uint64) {
	for _, u := range s.paths {
		size += u.Available()
	}
	return
}

// TotalSize implements Interface.
func (s *Storage) TotalSize() (size uint64) {
	for _, u := range s.paths {
		size += u.Size()
	}
	return
}

// AddPaths implements Interface
func (s *Storage) AddPaths(paths ...string) (added int) {
	for _, path := range paths {
		if utils.CheckIfDirectoryExists(path) {
			s.paths[path] = usage.NewUsage(path)
			added += 1
		}
	}
	return
}

func New() Interface {
	return &Storage{
		paths: make(map[string]usage.Interface),
	}
}
