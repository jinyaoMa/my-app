package storage

import (
	"io/fs"
	"os"
)

type Interface interface {
	// AddPaths add scopes to the storage for storing files
	AddPaths(paths ...string) (added int, err error)

	GetMountpointUsage() (u MountpointUsage, err error)

	SearchFile(filename string, isCache bool) (file *os.File, err error)

	// Upload:Cache
	Cache(filename string, data []byte, rangeStart uint64, rangeEnd uint64, totalSize uint64) (err error)

	// Upload:Cache:Clear
	ClearCache(filename string) (err error)

	// Upload:Persist
	Persist(filename string) (err error)

	// Download:Load
	Load(filename string, rangeStart uint64, rangeEnd uint64) (file fs.File, err error)

	// Upload+Download:Checksum
	Checksum(filename string, checksum string, isCache bool) (ok bool)
}
