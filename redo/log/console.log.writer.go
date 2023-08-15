package log

import "os"

type ConsoleLogWriter struct {
	*LogWriter
}

func (w *ConsoleLogWriter) Write(p []byte) (n int, err error) {
	if n, err = os.Stderr.Write(p); err == nil {
		return
	}
	return w.LogWriter.Write(p)
}

func NewConsoleLogWriter(children ...IChainWriter) *ConsoleLogWriter {
	return &ConsoleLogWriter{
		LogWriter: NewLogWriter(children...),
	}
}

func NewIConsoleLogWriter(children ...IChainWriter) IChainWriter {
	return NewConsoleLogWriter(children...)
}
