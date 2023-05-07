package snowflake

import (
	"time"

	"github.com/imdario/mergo"
)

const (
	TotalShareBits uint8 = 22
)

type Options struct {
	Epoch      time.Time // started timestamp, store into 41 bits
	NodeBits   uint8     // bits to store nodes ids
	StepBits   uint8     // bits to store increment ids within a millisecond
	NodeNumber int64     // current node id/number
}

func DefaultOptions() *Options {
	var epoch time.Time = time.Date(2023, 5, 7, 23, 23, 23, 233, time.UTC) // 2023-05-07 23:23:23.233
	var nodeBits uint8 = 10                                                // max 1024 nodes
	var stepBits uint8 = TotalShareBits - nodeBits                         // max 4096 ids/ms

	return &Options{
		Epoch:      epoch,
		NodeBits:   nodeBits,
		StepBits:   stepBits,
		NodeNumber: 1024, // max node id
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
