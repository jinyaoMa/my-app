package utility

import (
	"os"
	"path/filepath"
	"strings"
)

type IUtility interface {
	// GetExecutableName get the filename with the same name as application executable
	// but specify a different extension
	GetExecutableFileName(ext string) string

	// GetExecutablePath get the path started from application executable's directory
	GetExecutablePath(elem ...string) string
}

type Utility struct {
	executableName string // the file name (without extension) of application executable
	executableDir  string // the folder that application executable located
}

func NewUtility() (IUtility, error) {
	// get executable directory
	exe, err := os.Executable()
	if err != nil {
		return nil, err
	}
	executableName := strings.SplitN(filepath.Base(exe), ".", 2)[0]
	executableDir := filepath.Dir(exe)

	return &Utility{
		executableName: executableName,
		executableDir:  executableDir,
	}, nil
}

// GetExecutableFileName implements IUtility
func (u *Utility) GetExecutableFileName(ext string) string {
	return u.executableName + "." + ext
}

// GetExecutablePath implements IUtility
func (u *Utility) GetExecutablePath(elem ...string) string {
	if len(elem) == 0 {
		return u.executableDir
	}
	return filepath.Join(append([]string{u.executableDir}, elem...)...)
}
