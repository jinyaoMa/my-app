package usage

type Interface interface {
	// Free returns total free bytes on file system
	Free() uint64

	// Available returns total available bytes on file system to an unprivileged user
	Available() uint64

	// Size returns total size of the file system
	Size() uint64

	// Used returns total bytes used in file system
	Used() uint64

	// Usage returns percentage of use on the file system
	Usage() float32
}
