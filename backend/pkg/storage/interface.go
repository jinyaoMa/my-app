package storage

import (
	"os"
)

type Interface interface {
	// AddPaths add scopes to the storage for storing files
	AddPaths(paths ...string) (added int, err error)

	GetMountpointUsage() (u MountpointUsage, err error)

	SearchFile(filename string, isCache bool) (file *os.File, path string, err error)

	GetCacheFiles(filename string) (files []*StorageCacheFile, err error)

	// Upload:Cache
	Cache(filename string, data []byte, rangeStart int64, rangeEnd int64, size int64, forceCache bool) (ok bool, path string, err error)

	// Upload:Cache:Clear
	ClearCache(filename string) (err error)

	// Upload:Persist
	Persist(filename string, cacheFilepaths []string, totalSize int64) (err error)

	// Download:Load
	Load(filename string, rangeStart int64, rangeEnd int64) (data []byte, err error)

	// Upload+Download:Checksum
	// md5:sha512:File.Size => 32 + 128 = 160 hex digits + 2[:] + 20[int64] = 182 (size)
	Checksum(filename string, isCache bool) (checksum string, paths []string, err error)

	VerifyChecksum(filename string, isCache bool, checksum string) (ok bool, paths []string, err error)
}
