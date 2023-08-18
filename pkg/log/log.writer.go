package log

type LogWriter struct {
	children []ITreeWriter // write next until nil
}

// Add implements IChainWriter.
func (w *LogWriter) Add(child ITreeWriter) {
	w.children = append(w.children, child)
}

// Next implements IChainWriter.
func (w *LogWriter) Children() []ITreeWriter {
	return w.children
}

// Write implements io.Writer.
func (w *LogWriter) Write(p []byte) (n int, err error) {
	if children := w.Children(); len(children) > 0 {
		for _, child := range children {
			if n_, err_ := child.Write(p); err_ != nil {
				n, err = n_, err_
			}
		}
		return
	}
	return
}

func NewLogWriter(children ...ITreeWriter) *LogWriter {
	return &LogWriter{
		children: children,
	}
}

func NewILogWriter(children ...ITreeWriter) ITreeWriter {
	return NewLogWriter(children...)
}
