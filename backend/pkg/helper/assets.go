package helper

import (
	"io"
	"io/fs"
	"os"
)

type Assets struct {
	fs.FS
}

func NewAssets(dir string) *Assets {
	return &Assets{
		FS: os.DirFS(dir),
	}
}

func (a *Assets) GetBytes(path string) (data []byte) {
	if file, err := a.Open(path); err == nil {
		defer file.Close()
		if data, err = io.ReadAll(file); err != nil {
			return nil
		}
	}
	return
}
