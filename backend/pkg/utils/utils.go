package utils

import (
	"crypto/sha1"
	"encoding/hex"
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

// sha1 hash, str => string to hash, return hexadecimal encoded hashed string
func SHA1(str string) string {
	hashed := sha1.Sum([]byte(str))
	return hex.EncodeToString(hashed[:])
}
