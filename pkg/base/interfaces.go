package base

import "io/fs"

type IAssets interface {
	ReadBytes(path string) (bytes []byte, err error)
	Sub(dir string) (subAssets IAssets, err error)
	WalkFiles(dir string, fn func(path string, filename string, fileType fs.FileMode) error) (err error)
}

type ICommand interface {
	Open()
	Close()
}
