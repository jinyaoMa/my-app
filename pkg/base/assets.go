package base

import (
	"io/fs"
	"os"
)

type Assets struct {
	fs.FS
}

func (assets *Assets) ReadBytes(path string) (bytes []byte, err error) {
	f, err := assets.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	if _, err := f.Read(bytes); err != nil {
		return nil, err
	}
	return
}

func (assets *Assets) Sub(dir string) (subAssets *Assets, err error) {
	sub, err := fs.Sub(assets, dir)
	if err != nil {
		return nil, err
	}
	return &Assets{
		FS: sub,
	}, nil
}

func (assets *Assets) WalkFiles(dir string, fn func(path string, filename string, fileType fs.FileMode) error) (err error) {
	return fs.WalkDir(assets, dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		return fn(path, d.Name(), d.Type())
	})
}

func NewAssets(apath string) (assets *Assets) {
	assets = &Assets{
		FS: os.DirFS(apath),
	}
	return assets
}
