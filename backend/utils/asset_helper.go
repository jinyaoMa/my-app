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
	// Walk walk through all files shallowly (not drill down to folders) in a directory of FS
	Walk(callback func(path string, isDir bool, f fs.DirEntry) error, dirPaths ...string) error
	// Extract extract root directory to desinatiion directory
	Extract(dst ...string) error
}

type EmbedFS struct {
	fs   embed.FS
	root string
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
	// current directory
	cd := filepath.Join(dirPaths...)
	cd = filepath.ToSlash(cd)

	dir := filepath.Join(ef.root, cd)
	dir = filepath.ToSlash(dir)
	files, err := ef.fs.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, f := range files {
		if err := callback(cd+"/"+f.Name(), f.IsDir(), f); err != nil {
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
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	return ef._extract(files, ef.root, dir)
}

// _extract drill down the directory entry and extract all folders and files
func (ef *EmbedFS) _extract(files []fs.DirEntry, cd string, ecd string) error {
	for _, f := range files {
		// path of current directory/file name
		_cd := cd + "/" + f.Name()
		// path of extracted current directory/file name
		_ecd := filepath.Join(ecd, f.Name())

		if f.IsDir() { // extract the folder
			_files, err := ef.fs.ReadDir(_cd)
			if err != nil {
				return err
			}

			err = os.MkdirAll(_ecd, os.ModePerm)
			if err != nil {
				return err
			}

			err = ef._extract(_files, _cd, _ecd)
			if err != nil {
				return err
			}

		} else { // extract the file
			fileContent, err := ef.fs.ReadFile(_cd)
			if err != nil {
				return err
			}

			if err := os.WriteFile(_ecd, fileContent, 0666); err != nil {
				return err
			}
		}
	}
	return nil
}

type DirFS struct {
	fs fs.FS
}

// NewDirFS get ready to read data from a directory FS with a root path specified
func NewDirFS(dirRoot ...string) AssetHelper {
	root := filepath.Join(dirRoot...)
	root = filepath.ToSlash(root)
	return &DirFS{
		fs: os.DirFS(root),
	}
}

// GetFileBytes implements MyFS
func (df *DirFS) GetFileBytes(filePaths ...string) (data []byte, err error) {
	name := filepath.Join(filePaths...)
	name = filepath.ToSlash(name)
	data, err = fs.ReadFile(df.fs, name)
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
	// current directory
	cd := filepath.Join(dirPaths...)
	cd = filepath.ToSlash(cd)

	dir := filepath.Join(".", cd)
	dir = filepath.ToSlash(dir)
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
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	return df._extract(files, ".", dir)
}

// _extract drill down the directory entry and extract all folders and files
func (df *DirFS) _extract(files []fs.DirEntry, cd string, ecd string) error {
	for _, f := range files {
		// path of current directory/file name
		_cd := cd + "/" + f.Name()
		// path of extracted current directory/file name
		_ecd := filepath.Join(ecd, f.Name())

		if f.IsDir() { // extract the folder
			_files, err := fs.ReadDir(df.fs, _cd)
			if err != nil {
				return err
			}

			err = os.MkdirAll(_ecd, os.ModePerm)
			if err != nil {
				return err
			}

			err = df._extract(_files, _cd, _ecd)
			if err != nil {
				return err
			}

		} else { // extract the file
			fileContent, err := fs.ReadFile(df.fs, _cd)
			if err != nil {
				return err
			}

			if err := os.WriteFile(_ecd, fileContent, 0666); err != nil {
				return err
			}
		}
	}
	return nil
}
