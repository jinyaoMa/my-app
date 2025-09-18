package cflog

import (
	"errors"
	"io"
	"os"
)

func NewWriter(path string, enableConsole bool) (io.Writer, error) {
	return new(Writer).init(path, enableConsole)
}

type Writer struct {
	enableConsole bool
	file          *os.File
}

func (w *Writer) Write(p []byte) (n int, err error) {
	if w.enableConsole {
		_, err = os.Stdout.Write(p)
		if err != nil {
			return 0, err
		}
	}
	return w.file.Write(p)
}

func (w *Writer) init(path string, enableConsole bool) (*Writer, error) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, errors.Join(errors.New("failed to init cflog writer"), err)
	}
	w.file = file
	w.enableConsole = enableConsole
	return w, nil
}
