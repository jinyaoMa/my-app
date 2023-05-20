package utils

import (
	"os"
	"path/filepath"
)

// the file name (without extension) of application executable
func executableFilenameWithoutExtension() (string, error) {
	// get executable directory
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	return GetFilenameWithoutExtension(exe), nil
}

// the folder that application executable located
func executableDirectory() (string, error) {
	// get executable directory
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(exe), nil
}
