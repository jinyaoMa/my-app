package id

import (
	"time"

	"dario.cat/mergo"
)

const (
	TotalShareBits uint8 = 22 // total 22 bits to share between node/step
)

type Config struct {
	Epoch      time.Time // started timestamp, store into 41 bits
	NodeBits   uint8     // bits to store nodes ids
	StepBits   uint8     // bits to store increment ids within a millisecond
	NodeNumber int64     // current node id/number
}

// DefaultConfig return default Config with node id 0
func DefaultConfig() *Config {
	var epoch time.Time = time.Date(2023, 5, 7, 23, 23, 23, 233, time.UTC) // 2023-05-07 23:23:23.233
	var nodeBits uint8 = 10                                                // max 1024 nodes
	var stepBits uint8 = TotalShareBits - nodeBits                         // max 4096 ids/ms

	return &Config{
		Epoch:      epoch,
		NodeBits:   nodeBits,
		StepBits:   stepBits,
		NodeNumber: 0, // min node id
	}
}

// NewConfig override the default Config of ID generator
func NewConfig(dst *Config) *Config {
	src := DefaultConfig()
	err := mergo.Merge(dst, *src)
	if err != nil {
		return src
	}
	return dst
}
