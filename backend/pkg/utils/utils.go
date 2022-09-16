package utils

import (
	"log"
	"os"
	"path/filepath"
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
