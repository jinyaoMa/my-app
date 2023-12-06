package ctick

import (
	"my-app/pkg/base"
	"time"
)

type CodeTickerOptions struct {
	base.Options
	Expiration time.Duration
	Size       uint
	Chars      []string
}

func DefaultCodeTickerOptions() *CodeTickerOptions {
	return &CodeTickerOptions{
		Expiration: 10 * time.Minute,
		Size:       6,
		Chars:      base.CodeDigits,
	}
}

func NewCodeTickerOptions(dst *CodeTickerOptions) (*CodeTickerOptions, error) {
	return base.MergeOptions(DefaultCodeTickerOptions(), dst)
}
