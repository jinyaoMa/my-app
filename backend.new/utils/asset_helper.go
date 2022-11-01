package utils

import (
	"embed"
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
)

// AssetHelper access files shallowly in a directory/embed.FS
type AssetHelper interface {
	// GetFileBytes get bytes of a file from FS
	GetFileBytes(filePaths ...string) (data []byte, err error)
	// LoadJSON load JSON file to a go struct instance from FS
	LoadJSON(v any, filePaths ...string) error
	// Walk walk through all files shallowly in a directory of FS
	Walk(callback func(path string, isDir bool, f fs.DirEntry) error, dirPaths ...string) error
	// Extract extract root directory to desinatiion directory
	Extract(dst ...string) error
}

type EmbedFS struct {
	fs   embed.FS
	root string
}

// GetFileBytes implements MyFS
func (ef *EmbedFS) GetFileBytes(filePaths ...string) (data []byte, err error) {
	name := filepath.Join(append([]string{ef.root}, filePaths...)...)
	name = filepath.ToSlash(name)
	data, err = ef.fs.ReadFile(name)
	return
}

// LoadJSON implements MyFS
func (ef *EmbedFS) LoadJSON(v any, filePaths ...string) error {
	data, err := ef.GetFileBytes(filePaths...)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

// Walk implements MyFS
func (ef *EmbedFS) Walk(callback func(path string, isDir bool, f fs.DirEntry) error, dirPaths ...string) error {
	dir := filepath.Join(append([]string{ef.root}, dirPaths...)...)
	dir = filepath.ToSlash(dir)
	files, err := ef.fs.ReadDir(dir)
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

// Extract implements AssetHelper
func (ef *EmbedFS) Extract(dst ...string) error {
	files, err := ef.fs.ReadDir(ef.root)
	if err != nil {
		return err
	}

	dir := filepath.Join(dst...)
	os.MkdirAll(dir, os.ModePerm)

	for _, f := range files {
		fileContent, err := ef.fs.ReadFile(ef.root + "/" + f.Name())
		if err != nil {
			return err
		}

		filename := filepath.Join(dir, f.Name())
		if err := os.WriteFile(filename, fileContent, 0666); err != nil {
			return err
		}
	}
	return nil
}

// NewEmbedFs get ready to read data from an embed FS with a root path specified
func NewEmbedFS(fs embed.FS, dirRoot ...string) AssetHelper {
	root := filepath.Join(dirRoot...)
	root = filepath.ToSlash(root)
	return &EmbedFS{
		fs:   fs,
		root: root,
	}
}

type DirFS struct {
	fs   fs.FS
	root string
}

// GetFileBytes implements MyFS
func (df *DirFS) GetFileBytes(filePaths ...string) (data []byte, err error) {
	data, err = fs.ReadFile(df.fs, filepath.Join(filePaths...))
	return
}

// LoadJSON implements MyFS
func (df *DirFS) LoadJSON(v any, filePaths ...string) error {
	data, err := df.GetFileBytes(filePaths...)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

// Walk implements MyFS
func (df *DirFS) Walk(callback func(path string, isDir bool, f fs.DirEntry) error, dirPaths ...string) error {
	dir := filepath.Join(dirPaths...)
	files, err := fs.ReadDir(df.fs, dir)
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

// Extract implements AssetHelper
func (df *DirFS) Extract(dst ...string) error {
	files, err := fs.ReadDir(df.fs, ".")
	if err != nil {
		return err
	}

	dir := filepath.Join(dst...)
	os.MkdirAll(dir, os.ModePerm)

	for _, f := range files {
		fileContent, err := fs.ReadFile(df.fs, f.Name())
		if err != nil {
			return err
		}

		filename := filepath.Join(dir, f.Name())
		if err := os.WriteFile(filename, fileContent, 0666); err != nil {
			return err
		}
	}
	return nil
}

// NewDirFS get ready to read data from a directory FS with a root path specified
func NewDirFS(dirRoot ...string) AssetHelper {
	root := filepath.Join(dirRoot...)
	return &DirFS{
		fs:   os.DirFS(root),
		root: root,
	}
}
