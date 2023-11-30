package base

import "dario.cat/mergo"

type IOptions interface {
	HasMerged() bool
	SetMerged(merged bool)
}

type Options struct {
	Merged bool
}

// GetMerged implements IOptions.
func (options *Options) HasMerged() bool {
	return options.Merged
}

// SetMerged implements IOptions.
func (options *Options) SetMerged(merged bool) {
	options.Merged = merged
}

func NewOptions() (*Options, IOptions) {
	options := new(Options)
	return options, options
}

func MergeOptions[T IOptions](src T, dst T) (T, error) {
	if dst.HasMerged() {
		return dst, nil
	}
	err := mergo.Merge(dst, src)
	if err != nil {
		return src, err
	}
	dst.SetMerged(true)
	return dst, nil
}
