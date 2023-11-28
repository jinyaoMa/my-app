package base

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func IsFileExists(path string) (exists bool) {
	fileInfo, err := os.Stat(path)
	if err == nil && !fileInfo.IsDir() {
		return true
	}
	return false
}

func IsDirectoryExists(path string) (exists bool) {
	fileInfo, err := os.Stat(path)
	if err == nil && fileInfo.IsDir() {
		return true
	}
	return false
}

func IsDirectoryEmpty(path string) (isEmpty bool, err error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return false, err
	}
	if !fi.IsDir() {
		e := fmt.Sprintf("path %s is not a directory", path)
		return false, errors.New(e)
	}
	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	return err == io.EOF, nil
}

func GetExecutable() (directory string, filename string, err error) {
	exe, err := os.Executable()
	if err != nil {
		return "", "", err
	}
	return filepath.Dir(exe), filepath.Base(exe), nil
}
