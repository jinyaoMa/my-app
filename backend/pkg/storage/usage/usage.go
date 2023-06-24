//go:build !windows

package usage

import "syscall"

// Usage contains usage data and provides user-friendly access methods
type Usage struct {
	stat *syscall.Statfs_t
}

// Available implements Interface.
func (u *Usage) Available() uint64 {
	return u.stat.Bavail * uint64(u.stat.Bsize)
}

// Free implements Interface.
func (u *Usage) Free() uint64 {
	return u.stat.Bfree * uint64(u.stat.Bsize)
}

// Size implements Interface.
func (u *Usage) Size() uint64 {
	return uint64(u.stat.Blocks) * uint64(u.stat.Bsize)
}

// Usage implements Interface.
func (u *Usage) Usage() float32 {
	return float32(u.Used()) / float32(u.Size())
}

// Used implements Interface.
func (u *Usage) Used() uint64 {
	return u.Size() - u.Free()
}

// NewUsage returns an object holding the disk usage of volume path
// or nil in case of error (invalid path, etc)
func NewUsage(path string) Interface {
	var stat syscall.Statfs_t
	syscall.Statfs(path, &stat)
	return &Usage{&stat}
}
