package helper

import (
	"io"
	"os"
	"path/filepath"
)

// CheckIfFileExists check if a path exists, and it is a file
func CheckIfFileExists(path string) (exists bool) {
	fileInfo, err := os.Stat(path)
	if err == nil && !fileInfo.IsDir() {
		return true
	}
	return false
}

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
	filename, err = ExecutableFilenameWithoutExtension()
	if err != nil {
		return
	}

	filename += "." + ext
	return
}

// GetPathStartedFromExecutable get the path started from application executable's directory
func GetPathStartedFromExecutable(elem ...string) (path string, err error) {
	path, err = ExecutableDirectory()
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

// the file name (without extension) of application executable
func ExecutableFilenameWithoutExtension() (string, error) {
	// get executable directory
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	return GetFilenameWithoutExtension(exe), nil
}

// the folder that application executable located
func ExecutableDirectory() (string, error) {
	// get executable directory
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(exe), nil
}
