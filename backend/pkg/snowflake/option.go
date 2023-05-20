package snowflake

import (
	"time"

	"github.com/imdario/mergo"
)

const (
	TotalShareBits uint8 = 22 // total 22 bits to share between node/step
)

type Option struct {
	Epoch      time.Time // started timestamp, store into 41 bits
	NodeBits   uint8     // bits to store nodes ids
	StepBits   uint8     // bits to store increment ids within a millisecond
	NodeNumber int64     // current node id/number
}

// DefaultOption return default Option with node id 1023
func DefaultOption() *Option {
	var epoch time.Time = time.Date(2023, 5, 7, 23, 23, 23, 233, time.UTC) // 2023-05-07 23:23:23.233
	var nodeBits uint8 = 10                                                // max 1024 nodes
	var stepBits uint8 = TotalShareBits - nodeBits                         // max 4096 ids/ms

	return &Option{
		Epoch:      epoch,
		NodeBits:   nodeBits,
		StepBits:   stepBits,
		NodeNumber: 0, // min node id
	}
}

// NewOption override the default Option of Snowflake
func NewOption(dst *Option) *Option {
	src := DefaultOption()

	err := mergo.Merge(dst, *src)
	if err != nil {
		return src
	}

	return dst
}
