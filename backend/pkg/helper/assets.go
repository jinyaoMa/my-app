package helper

import (
	"io"
	"io/fs"
	"os"
)

type IAssets interface {
	fs.FS
	GetBytes(path string) (data []byte)
}

type Assets struct {
	fs.FS
}

// GetBytes implements IAssets.
func (a *Assets) GetBytes(path string) (data []byte) {
	if file, err := a.Open(path); err == nil {
		defer file.Close()
		if data, err = io.ReadAll(file); err != nil {
			return nil
		}
	}
	return
}

func NewAssets(dir string) IAssets {
	return &Assets{
		FS: os.DirFS(dir),
	}
}
