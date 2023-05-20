package storage

import (
	"math/big"
	"my-app/backend/pkg/utils"
)

type Storage struct {
	paths     []string
	totalSize big.Int
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
