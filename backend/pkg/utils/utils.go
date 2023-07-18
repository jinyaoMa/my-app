package utils

import (
	"io"
	"os"
	"path/filepath"
)

// CheckIfDirectoryIsEmpty check if a directory is empty
func CheckIfDirectoryIsEmpty(path string) (isEmpty bool) {
	f, err := os.Open(path)
	if err != nil {
		return false
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	return err == io.EOF
}

// CheckIfDirectoryExists check if a path exists, and it is a directory
func CheckIfDirectoryExists(path string) (exists bool) {
	fileInfo, err := os.Stat(path)
	if err == nil && fileInfo.IsDir() {
		return true
	}
	return false
}

// GetFilenameSameAsExecutable get the filename with the same name as
// application executable but specify a different extension
func GetFilenameSameAsExecutable(ext string) (filename string, err error) {
	filename, err = executableFilenameWithoutExtension()
	if err != nil {
		return
	}

	filename += "." + ext
	return
}

// GetPathStartedFromExecutable get the path started from application executable's directory
func GetPathStartedFromExecutable(elem ...string) (path string, err error) {
	path, err = executableDirectory()
	if err != nil || len(elem) == 0 {
		return
	}

	path = filepath.Join(append([]string{path}, elem...)...)
	return
}

// GetFilenameWithoutExtension get the filename without extension from path
func GetFilenameWithoutExtension(path string) string {
	for i := len(path) - 1; i >= 0 && !os.IsPathSeparator(path[i]); i-- {
		if path[i] == '.' {
			return path[:i]
		}
	}
	return filepath.Base(path)
}
