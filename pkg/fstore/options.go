package fstore

const (
	B  uint64 = 1
	KB        = 1024 * B
	MB        = 1024 * KB
	GB        = 1024 * MB
	TB        = 1024 * GB
)

type Options struct {
	CacheFolderName string
	BufferSize      uint64
	FragmentSize    uint64
}

var (
	DefaultOptions = &Options{
		CacheFolderName: ".cache",
		BufferSize:      4 * KB,
		FragmentSize:    4 * MB,
	}
)
