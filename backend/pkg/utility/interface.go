package utility

type Interface interface {
	// GetExecutableName get the filename with the same name as application executable
	// but specify a different extension
	GetExecutableFileName(ext string) string

	// GetExecutablePath get the path started from application executable's directory
	GetExecutablePath(elem ...string) string
}
