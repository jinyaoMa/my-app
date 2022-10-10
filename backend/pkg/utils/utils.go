package utils

import (
	"os"
	"path/filepath"
)

var (
	executablePath string
)

func init() {
	exe, err := os.Executable()
	if err != nil {
		panic("failed to get executable path")
	}
	executablePath = filepath.Dir(exe)
}

func GetExecutablePath(elem ...string) string {
	return filepath.Join(append([]string{executablePath}, elem...)...)
}
