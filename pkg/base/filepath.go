package base

import "path/filepath"

func GetFilenameWithoutExtension(path string) string {
	var filename = filepath.Base(path)
	for i := len(filename) - 1; i >= 0; i-- {
		if filename[i] == '.' {
			return filename[:i]
		}
	}
	return filename
}

func GetFilepathWithoutExtension(path string) string {
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '.' {
			return path[:i]
		}
	}
	return path
}
