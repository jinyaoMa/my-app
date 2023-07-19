package storage

import (
	"io/fs"
)

type Interface interface {
	// AddPaths add scopes to the storage for storing files
	AddPaths(paths ...string) (added int, err error)

	GetMountpointUsage() (u MountpointUsage, err error)

	// Upload:Cache
	Cache(file fs.File, rangeStart uint64, rangeEnd uint64) (err error)

	// Upload:Cache:Clear
	ClearCache(filename string) (err error)

	// Upload:Persist
	Persist(filename string) (err error)

	// Download:Load
	Load(filename string, rangeStart uint64, rangeEnd uint64) (file fs.File, err error)

	// Upload+Download:Checksum
	Checksum(filename string, checksum string, isCache bool) (ok bool)
}
