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

var _utils = &utils{}

type utils struct {
	once          sync.Once
	executableDir string  // the folder that application executable located
	panicLogger   *Logger // handle app initializing panics
}

// utils resources for global use
func Utils() *utils {
	_utils.once.Do(func() {
		// get executable directory
		exe, err := os.Executable()
		if err != nil {
			panic("failed to get executable path")
		}
		_utils.executableDir = filepath.Dir(exe)

		// initialize panic logger
		panicLogPath := filepath.Join(_utils.executableDir, "MyApp.panic")
		panicFile, err := os.OpenFile(
			panicLogPath,
			os.O_CREATE|os.O_WRONLY|os.O_APPEND,
			0666,
		)
		if err != nil {
			panic("failed to open panic log: " + panicLogPath)
		}
		_utils.panicLogger = NewFileLogger("", panicFile)
	})
	return _utils
}

func (u *utils) Panic(v ...any) {
	u.panicLogger.Panicln(v...)
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
