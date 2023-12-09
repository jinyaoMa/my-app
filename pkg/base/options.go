package base

import "dario.cat/mergo"

type Options struct {
	merged bool
}

// GetMerged implements IOptions.
func (options *Options) HasMerged() bool {
	return options.merged
}

// SetMerged implements IOptions.
func (options *Options) SetMerged(merged bool) {
	options.merged = merged
}

func NewOptions() (options *Options, iOptions IOptions) {
	options = new(Options)
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
