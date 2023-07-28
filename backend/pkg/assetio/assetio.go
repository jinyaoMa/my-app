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

func (a *Assetio) Root() string {
	return a.root
}

func NewAssetio(root string) (a Interface, err error) {
	return &Assetio{
		FS:   os.DirFS(root),
		root: root,
	}, nil
}
