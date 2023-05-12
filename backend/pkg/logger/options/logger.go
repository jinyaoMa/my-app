package options

import (
	"io"
	"log"
	"os"

	"github.com/imdario/mergo"
)

type OLogger struct {
	Writer         io.Writer
	Tag            string
	PrefixTemplate func(tag string) (prefix string)
	Flags          int
}

func DefaultOLogger() *OLogger {
	return &OLogger{
		Writer: os.Stderr,
		Tag:    "STD",
		PrefixTemplate: func(tag string) (prefix string) {
			return "[" + tag + "]"
		},
		Flags: log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile,
	}
}

func NewOLogger(dst *OLogger) *OLogger {
	src := DefaultOLogger()

	err := mergo.Merge(dst, *src)
	if err != nil {
		return src
	}

	return dst
}
