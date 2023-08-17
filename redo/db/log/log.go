package log

import (
	"my-app/redo/log"
	"os"
)

type DbLogWriter struct {
	*log.LogWriter
	logFile *os.File
}

func (w *DbLogWriter) Write(p []byte) (n int, err error) {
	if n, err = w.logFile.Write(p); err == nil {
		return
	}
	return w.LogWriter.Write(p)
}

func NewFileLogWriter(filename string, children ...log.ITreeWriter) (w *DbLogWriter, err error) {
	var logFile *os.File
	logFile, err = os.OpenFile(filename, os.O_APPEND|os.O_RDWR|os.O_CREATE, os.ModeAppend)
	if err != nil {
		return
	}
	w = &DbLogWriter{
		LogWriter: log.NewLogWriter(children...),
		logFile:   logFile,
	}
	return
}

func NewIFileLogWriter(filename string, children ...log.ITreeWriter) (w log.ITreeWriter, err error) {
	return NewFileLogWriter(filename, children...)
}
