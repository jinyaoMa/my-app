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

	GetCacheFiles(filename string) (files []*os.File, err error)

	// Upload:Cache
	Cache(filename string, data []byte, rangeStart uint64, rangeEnd uint64, size uint64, forceCache bool) (ok bool, err error)

	// Upload:Cache:Clear
	ClearCache(filename string) (err error)

	// Upload:Persist
	Persist(filename string) (err error)

	// Download:Load
	Load(filename string, rangeStart uint64, rangeEnd uint64) (file fs.File, err error)

	// Upload+Download:Checksum
	// md5:sha512:File.Size => 32 + 128 = 160 hex digits + 2[:] + 20[int64] = 182 (size)
	Checksum(filename string, isCache bool) (checksum string, err error)

	VerifyChecksum(filename string, isCache bool, checksum string) (ok bool, err error)
}
