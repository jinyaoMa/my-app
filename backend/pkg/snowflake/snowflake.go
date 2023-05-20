package snowflake

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Snowflake struct {
	mu    sync.Mutex
	epoch time.Time
	time  int64
	node  int64
	step  int64

	nodeMax   int64
	nodeMask  int64
	stepMask  int64
	timeShift uint8
	nodeShift uint8
}

// Generate implements Interface
func (s *Snowflake) Generate() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Since(s.epoch).Milliseconds()

	if now == s.time {
		s.step = (s.step + 1) & s.stepMask

		if s.step == 0 { // step bits go over, move to next timestamp
			for now <= s.time {
				now = time.Since(s.epoch).Milliseconds()
			}
		}
	} else {
		s.step = 0
	}

	s.time = now

	r := int64((now)<<s.timeShift |
		(s.node << s.nodeShift) |
		(s.step),
	)

	return r
}

// Default return Snowflake Id generator with default options
func Default() (Interface, error) {
	return New(DefaultOption())
}

// New return Snowflake Id generator with custom options
func New(opts *Option) (Interface, error) {
	opts = NewOption(opts)

	var shareBits uint8 = opts.NodeBits + opts.StepBits
	if shareBits > TotalShareBits {
		return nil, fmt.Errorf("remember, you have a total %d bits to share between node/step", TotalShareBits)
	}

	var nodeMax int64 = -1 ^ (-1 << opts.NodeBits)
	if opts.NodeNumber < 0 || opts.NodeNumber > nodeMax {
		return nil, fmt.Errorf("node number must be between 0 and %s", strconv.FormatInt(nodeMax, 10))
	}

	return &Snowflake{
		epoch:     opts.Epoch,
		node:      opts.NodeNumber,
		nodeMax:   nodeMax,
		nodeMask:  nodeMax << opts.StepBits,
		stepMask:  -1 ^ (-1 << opts.StepBits),
		timeShift: shareBits,
		nodeShift: opts.StepBits,
	}, nil
}
