package server

import "github.com/imdario/mergo"

type Options struct {
	UseHttps bool
}

func DefaultOptions() *Options {
	return &Options{
		UseHttps: true,
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
