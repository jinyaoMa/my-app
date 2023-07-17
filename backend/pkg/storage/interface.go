package storage

type Interface interface {
	// AddPaths add scopes to the storage for storing files
	AddPaths(paths ...string) (added int, err error)

	GetMountpointUsage() (u MountpointUsage, err error)
}
