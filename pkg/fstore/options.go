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
	ThresholdSize   uint64 // space tried to keep per storage
	BufferSize      uint64
	FragmentSize    uint64
}

var (
	DefaultOptions = &Options{
		CacheFolderName: ".cache",
		ThresholdSize:   8 * GB,
		BufferSize:      8 * KB,
		FragmentSize:    8 * MB,
	}
)
