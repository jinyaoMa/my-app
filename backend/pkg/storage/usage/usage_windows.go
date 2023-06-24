package usage

import (
	"syscall"
	"unsafe"
)

type Usage struct {
	freeBytes  int64
	totalBytes int64
	availBytes int64
}

// Available implements Interface.
func (u *Usage) Available() uint64 {
	return uint64(u.availBytes)
}

// Free implements Interface.
func (u *Usage) Free() uint64 {
	return uint64(u.freeBytes)
}

// Size implements Interface.
func (u *Usage) Size() uint64 {
	return uint64(u.totalBytes)
}

// Usage implements Interface.
func (u *Usage) Usage() float32 {
	return float32(u.Used()) / float32(u.Size())
}

// Used implements Interface.
func (u *Usage) Used() uint64 {
	return u.Size() - u.Free()
}

// NewUsages returns an object holding the disk usage of volume path
// or nil in case of error (invalid path, etc)
func NewUsage(path string) Interface {

	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")

	ptr, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return nil
	}

	u := &Usage{}

	c.Call(
		uintptr(unsafe.Pointer(ptr)),
		uintptr(unsafe.Pointer(&u.freeBytes)),
		uintptr(unsafe.Pointer(&u.totalBytes)),
		uintptr(unsafe.Pointer(&u.availBytes)))

	return u
}
