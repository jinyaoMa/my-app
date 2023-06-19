package storage

import (
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
	paths     []string
	totalSize uint64
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

func New() Interface {
	return &Storage{}
}
