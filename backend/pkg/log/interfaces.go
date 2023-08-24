package log

import "io"

type ITreeWriter interface {
	io.Writer

	Children() []ITreeWriter

	Add(child ITreeWriter)
}
