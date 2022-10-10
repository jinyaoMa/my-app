package utils

import (
	"embed"
	"encoding/json"
	"io/fs"
	"path/filepath"
)

type EmbedFs struct {
	fs   embed.FS
	root string
}

// get ready to read data from an embed fs with a root path specified
func NewEmbedFs(fs embed.FS, rootPaths ...string) *EmbedFs {
	root := filepath.Join(rootPaths...)
	root = filepath.ToSlash(root)
	return &EmbedFs{
		fs:   fs,
		root: root,
	}
}

// get bytes of a file from embed fs
func (e *EmbedFs) GetFileBytes(filePaths ...string) (data []byte, err error) {
	name := filepath.Join(append([]string{e.root}, filePaths...)...)
	name = filepath.ToSlash(name)
	data, err = e.fs.ReadFile(name)
	return
}

// load JSON file to a go struct instance from embed fs
func (e *EmbedFs) LoadJSON(v any, filePaths ...string) error {
	data, err := e.GetFileBytes(filePaths...)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

// wallk through all files in a directory of embed fs
func (e *EmbedFs) WalkDir(callback func(path string, isDir bool, f fs.DirEntry) error, dirPaths ...string) error {
	dir := filepath.Join(append([]string{e.root}, dirPaths...)...)
	dir = filepath.ToSlash(dir)
	files, err := e.fs.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, f := range files {
		if err := callback(dir+"/"+f.Name(), f.IsDir(), f); err != nil {
			return err
		}
	}
	return nil
}
