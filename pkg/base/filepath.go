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
