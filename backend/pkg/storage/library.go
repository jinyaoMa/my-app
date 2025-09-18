package storage

type Library struct {
	Mountpoint string `json:"mountpoint"` // mount point
	Directory  string `json:"directory"`  // absolute path inside mountpoint for storage library
	Disabled   bool   `json:"disabled"`   // disable this storage library
}
