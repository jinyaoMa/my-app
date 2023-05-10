package logger

import (
	"io"
	"log"
	"os"

	"github.com/imdario/mergo"
)

type Options struct {
	Writer         io.Writer
	Tag            string
	PrefixTemplate func(tag string) (prefix string)
	Flags          int
}

func DefaultOptions() *Options {
	return &Options{
		Writer: os.Stderr,
		Tag:    "STD",
		PrefixTemplate: func(tag string) (prefix string) {
			return "[" + tag + "] "
		},
		Flags: log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile,
	}
}

func NewOptions(opts *Options) *Options {
	src := DefaultOptions()

	err := mergo.Merge(opts, *src)
	if err != nil {
		return src
	}

	return opts
}
