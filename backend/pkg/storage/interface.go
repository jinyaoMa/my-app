package storage

type Interface interface {
	// AddPaths add scopes to the storage for storing files
	AddPaths(paths ...string) (added int, err error)

	// TotalSize returns total size of the file system
	TotalSize() (size uint64)

	// TotalAvailable returns total available bytes on file system to an unprivileged user
	TotalAvailable() (size uint64)

	MountpointUsage() (usage map[string]*MountpointStat, err error)

	AvailableMountPoints() (mountpoints []string, err error)
}
