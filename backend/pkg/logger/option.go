package logger

import (
	"io"
	"log"
	"os"

	"github.com/imdario/mergo"
)

type Option struct {
	Writer         io.Writer
	Tag            string
	PrefixTemplate func(tag string) (prefix string)
	Flags          int
}

func DefaultOption() *Option {
	return &Option{
		Writer: os.Stderr,
		Tag:    "STD",
		PrefixTemplate: func(tag string) (prefix string) {
			return "[" + tag + "]"
		},
		Flags: log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile,
	}
}

func NewOption(dst *Option) *Option {
	src := DefaultOption()

	err := mergo.Merge(dst, *src)
	if err != nil {
		return src
	}

	return dst
}
