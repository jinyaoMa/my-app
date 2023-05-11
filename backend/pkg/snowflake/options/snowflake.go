package options

import (
	"time"

	"github.com/imdario/mergo"
)

const (
	TotalShareBits uint8 = 22 // total 22 bits to share between node/step
)

type OSnowflake struct {
	Epoch      time.Time // started timestamp, store into 41 bits
	NodeBits   uint8     // bits to store nodes ids
	StepBits   uint8     // bits to store increment ids within a millisecond
	NodeNumber int64     // current node id/number
}

// DefaultOSnowflake return default OSnowflake with node id 1023
func DefaultOSnowflake() *OSnowflake {
	var epoch time.Time = time.Date(2023, 5, 7, 23, 23, 23, 233, time.UTC) // 2023-05-07 23:23:23.233
	var nodeBits uint8 = 10                                                // max 1024 nodes
	var stepBits uint8 = TotalShareBits - nodeBits                         // max 4096 ids/ms

	return &OSnowflake{
		Epoch:      epoch,
		NodeBits:   nodeBits,
		StepBits:   stepBits,
		NodeNumber: 1023, // max node id
	}
}

// NewOSnowflake override the default OSnowflake of Snowflake
func NewOSnowflake(dst *OSnowflake) *OSnowflake {
	src := DefaultOSnowflake()

	err := mergo.Merge(dst, *src)
	if err != nil {
		return src
	}

	return dst
}
