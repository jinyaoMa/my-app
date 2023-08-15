package log

import "os"

type FileLogWriter struct {
	*LogWriter
	logFile *os.File
}

func (w *FileLogWriter) Write(p []byte) (n int, err error) {
	if n, err = w.logFile.Write(p); err == nil {
		return
	}
	return w.LogWriter.Write(p)
}

func NewFileLogWriter(filename string, children ...IChainWriter) (w *FileLogWriter, err error) {
	var logFile *os.File
	logFile, err = os.OpenFile(filename, os.O_APPEND|os.O_RDWR|os.O_CREATE, os.ModeAppend)
	if err != nil {
		return
	}
	w = &FileLogWriter{
		LogWriter: NewLogWriter(children...),
		logFile:   logFile,
	}
	return
}

func NewIFileLogWriter(filename string, children ...IChainWriter) (w IChainWriter, err error) {
	return NewFileLogWriter(filename, children...)
}