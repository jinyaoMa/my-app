package id

import (
	"time"

	"dario.cat/mergo"
)

const (
	TotalShareBits uint = 22 // total 22 bits to share between node/step
)

type Config struct {
	Epoch      time.Time `comment:"started timestamp, store into 41 bits"`
	NodeBits   uint      `comment:"bits to store nodes ids, e.g. 10 => max 1024 nodes"`
	StepBits   uint      `comment:"bits to store increment ids within a millisecond, StepBits equals to TotalShareBits (total 22 bits to share between node/step) minus NodeBits, e.g. 22 - 10 = 12 => max 4096 ids/ms"`
	NodeNumber int64     `comment:"current node id/number, e.g. if max 4096 ids/ms, then min node id is 0 and max node id is 4095"`
}

// DefaultConfig return default Config with node id 0
func DefaultConfig() *Config {
	var epoch time.Time = time.Date(2023, 5, 7, 23, 23, 23, 233, time.UTC) // 2023-05-07 23:23:23.233
	var nodeBits uint = 10                                                 // max 1024 nodes
	var stepBits uint = TotalShareBits - nodeBits                          // max 4096 ids/ms

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
