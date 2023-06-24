package storage

type Interface interface {
	// AddPaths add scopes to the storage for storing files
	AddPaths(paths ...string) (added int)

	// TotalSize returns total size of the file system
	TotalSize() (size uint64)
}
