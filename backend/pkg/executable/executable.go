package executable

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type IExecutable interface {
	GetPath() string
	GetDir() string
	GetBase() string
	GetExt() string
	GetName() string
	JoinDir(elem ...string) string
	GetPathWithName(name string) string
	GetBaseWithExt(ext string) string
	GetPathWithExt(ext string) string
}

func MustNew() IExecutable {
	e, err := New()
	if err != nil {
		panic(err)
	}
	return e
}

func New() (IExecutable, error) {
	return new(executable).init()
}

type executable struct {
	path string // executable absolute path
	dir  string // executable absolute directory
	base string // executable filename include extension
	ext  string // executable extension
	name string // executable filename exclude extension
}

// GetPathWithExt implements IExecutable.
func (e *executable) GetPathWithExt(ext string) string {
	return filepath.Join(e.dir, e.GetBaseWithExt(ext))
}

// GetBaseWithExt implements IExecutable.
func (e *executable) GetBaseWithExt(ext string) string {
	if ext == "" {
		return e.name
	}

	if !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}
	return e.name + ext
}

// GetPathWithName implements IExecutable.
func (e *executable) GetPathWithName(name string) string {
	return filepath.Join(e.dir, name+e.ext)
}

// JoinDir implements IExecutable.
func (e *executable) JoinDir(elem ...string) string {
	return filepath.Join(append([]string{e.dir}, elem...)...)
}

// GetBase implements IExecutable.
func (e *executable) GetBase() string {
	return e.base
}

// GetDir implements IExecutable.
func (e *executable) GetDir() string {
	return e.dir
}

// GetExt implements IExecutable.
func (e *executable) GetExt() string {
	return e.ext
}

// GetName implements IExecutable.
func (e *executable) GetName() string {
	return e.name
}

// GetPath implements IExecutable.
func (e *executable) GetPath() string {
	return e.path
}

func (e *executable) init() (*executable, error) {
	// read executable path
	path, err := os.Executable()
	if err != nil {
		return e, errors.Join(errors.New("executable: failed to read executable path"), err)
	}

	e.path = path
	e.dir = filepath.Dir(path)
	e.base = filepath.Base(path)
	e.ext = filepath.Ext(path)
	e.name = e.base[:len(e.base)-len(e.ext)]
	return e, nil
}
