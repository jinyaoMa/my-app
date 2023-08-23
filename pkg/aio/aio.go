package aio

import (
	"encoding/json"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

type AIO struct {
	fs.FS
	root string
}

// GetBytes implements IAIO.
func (a *AIO) GetBytes(paths ...string) (data []byte) {
	if file, err := a.Open(filepath.Join(paths...)); err == nil {
		defer file.Close()
		if data, err = io.ReadAll(file); err != nil {
			return nil
		}
	}
	return
}

// LoadJSON implements IAIO.
func (a *AIO) LoadJSON(v interface{}, paths ...string) (ok bool) {
	data := a.GetBytes(paths...)
	if data == nil {
		return false
	}
	err := json.Unmarshal(data, v)
	return err == nil
}

// WalkDir implements IAIO.
func (a *AIO) WalkDir(callback func(path string, isDir bool, entry fs.DirEntry) (err error), paths ...string) (err error) {
	err = fs.WalkDir(a, filepath.Join(paths...), func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		return callback(path, d.IsDir(), d)
	})
	return
}

func (a *AIO) Root() string {
	return a.root
}

func New(root string) *AIO {
	return &AIO{
		FS:   os.DirFS(root),
		root: root,
	}
}

func NewIAIO(root string) IAIO {
	return New(root)
}
