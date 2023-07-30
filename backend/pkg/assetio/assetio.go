package assetio

import (
	"encoding/json"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

type Assetio struct {
	fs.FS
	root string
}

// GetBytes implements Interface.
func (a *Assetio) GetBytes(paths ...string) (data []byte) {
	if file, err := a.Open(filepath.Join(paths...)); err == nil {
		defer file.Close()
		if data, err = io.ReadAll(file); err != nil {
			return nil
		}
	}
	return
}

// LoadJSON implements Interface.
func (a *Assetio) LoadJSON(v interface{}, paths ...string) (ok bool) {
	data := a.GetBytes(paths...)
	if data == nil {
		return false
	}
	err := json.Unmarshal(data, v)
	return err == nil
}

// WalkDir implements Interface.
func (a *Assetio) WalkDir(callback func(path string, isDir bool, entry fs.DirEntry) (err error), paths ...string) (err error) {
	err = fs.WalkDir(a, filepath.Join(paths...), func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		return callback(path, d.IsDir(), d)
	})
	return
}

func (a *Assetio) Root() string {
	return a.root
}

func NewAssetio(root string) Interface {
	return &Assetio{
		FS:   os.DirFS(root),
		root: root,
	}
}
