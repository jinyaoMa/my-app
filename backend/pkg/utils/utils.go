package utils

import (
	"log"
	"os"
	"path/filepath"
)

const (
	Copyright = "© 2022 jinyaoMa"
)

var (
	executablePath string
)

func init() {
	var err error

	executablePath, err = os.Executable()
	if err != nil {
		log.Fatalf("fail to get executable path: %+v\n", err)
	}
	executablePath = filepath.Dir(executablePath)
}

func GetExecutablePath(elem ...string) string {
	return filepath.Join(append([]string{executablePath}, elem...)...)
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func GetDialogDirectoryPath(path string) string {
	if IsDir(path) {
		return path
	}
	next := filepath.Dir(path)
	if next == path {
		return path
	}
	return GetDialogDirectoryPath(path)
}
