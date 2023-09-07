package store

import (
	"os"
)

type IStore interface {
	// AddPaths add scopes to the storage for storing files
	AddPaths(paths ...string) (added int, err error)

	GetMountpointUsage() (u MountpointUsage, err error)

	SearchFile(filename string, isCache bool) (file *os.File, path string, err error)

	GetCacheFiles(filename string) (files []*StoreCacheFile, err error)

	// Upload:Cache
	Cache(filename string, data []byte, rangeStart int64, rangeEnd int64, size int64, forceCache bool) (ok bool, path string, err error)

	// Upload:Cache:Clear
	ClearCache(filename string) (err error)

	// Upload:Persist
	Persist(filename string, cacheFilepaths []string, totalSize int64) (ok bool, path string, err error)

	// Download:Load
	Load(filename string, rangeStart int64, rangeEnd int64) (data []byte, err error)

	// Upload+Download:Checksum
	// md5+sha256+crc32+size =>(hex encoded) 128bit/4+256bit/4+32bit/4+64bit/4=32+64+8+16=123bytes
	Checksum(filename string, isCache bool) (checksum string, paths []string, err error)

	VerifyChecksum(filename string, isCache bool, checksum string) (ok bool, paths []string, err error)
}
