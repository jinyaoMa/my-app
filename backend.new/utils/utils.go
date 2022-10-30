package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"os"
	"path/filepath"
	"sync"
)

const (
	Copyright = "© 2022 jinyaoMa"
)

var (
	instance *utils
	once     sync.Once
)

type utils struct {
	executableDir string
}

// Utils get utils global instance
func Utils() *utils {
	once.Do(func() {
		// get executable directory
		exe, err := os.Executable()
		if err != nil {
			panic("failed to get executable path")
		}
		executableDir := filepath.Dir(exe)

		// initialize utils
		instance = &utils{
			executableDir: executableDir,
		}
	})
	return instance
}

// GetExecutablePath get the path started from application's executable directory
func (u *utils) GetExecutablePath(elem ...string) string {
	if len(elem) == 0 {
		return u.executableDir
	}
	return filepath.Join(append([]string{u.executableDir}, elem...)...)
}

// HasDir check if the directory path exists
func (u *utils) HasDir(elem ...string) bool {
	fi, _ := os.Stat(filepath.Join(elem...))
	return fi != nil && fi.IsDir()
}

// HasExecutableDir check if the directory path started from application's executable directory exists
func (u *utils) HasExecutableDir(elem ...string) bool {
	return u.HasDir(u.GetExecutablePath(elem...))
}

// SHA1 sha1 hash, str => string to hash, return hexadecimal encoded hashed string
func (u *utils) SHA1(str string) string {
	hashed := sha1.Sum([]byte(str))
	return hex.EncodeToString(hashed[:])
}
