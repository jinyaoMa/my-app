package base

import "dario.cat/mergo"

type IOptions interface {
	SetMerged(merged bool)
}

type Options struct {
	Merged bool
}

// SetMerged implements IOptions.
func (options *Options) SetMerged(merged bool) {
	options.Merged = merged
}

func NewOptions() (*Options, IOptions) {
	options := new(Options)
	return options, options
}

func SimpleMerge[T IOptions](src T, dst T) T {
	err := mergo.Merge(dst, src)
	if err != nil {
		return src
	}
	dst.SetMerged(true)
	return dst
}
