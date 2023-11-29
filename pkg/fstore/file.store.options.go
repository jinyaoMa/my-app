package fstore

import (
	"my-app/pkg/base"
)

const (
	B  uint64 = 1
	KB        = 1024 * B
	MB        = 1024 * KB
	GB        = 1024 * MB
	TB        = 1024 * GB
)

type FileStoreOptions struct {
	base.Options
	CacheFolderName string
	ThresholdSize   uint64 // space tried to keep per storage
	BufferSize      uint64 // memory used to read data in iteration
	FragmentSize    uint64 // cache file fragment size
}

func DefaultFileStoreOptions() *FileStoreOptions {
	return &FileStoreOptions{
		CacheFolderName: ".cache",
		ThresholdSize:   8 * GB,
		BufferSize:      8 * KB,
		FragmentSize:    8 * MB,
	}
}

func NewFileStoreOptions(dst *FileStoreOptions) (*FileStoreOptions, error) {
	return base.SimpleMerge(DefaultFileStoreOptions(), dst)
}
