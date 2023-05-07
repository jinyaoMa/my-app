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

func Default() (*Snowflake, error) {
	return New(DefaultOptions())
}

func New(options *Options) (*Snowflake, error) {
	options = NewOptions(options)

	var shareBits uint8 = options.NodeBits + options.StepBits
	if shareBits > TotalShareBits {
		return nil, fmt.Errorf("remember, you have a total %d bits to share between node/step", TotalShareBits)
	}

	var nodeMax int64 = -1 ^ (-1 << options.NodeBits)
	if options.NodeNumber < 0 || options.NodeNumber > nodeMax {
		return nil, fmt.Errorf("node number must be between 0 and %s", strconv.FormatInt(nodeMax, 10))
	}

	return &Snowflake{
		epoch:     options.Epoch,
		node:      options.NodeNumber,
		nodeMax:   nodeMax,
		nodeMask:  nodeMax << options.StepBits,
		stepMask:  -1 ^ (-1 << options.StepBits),
		timeShift: shareBits,
		nodeShift: options.StepBits,
	}, nil
}

// Generate creates and returns a unique snowflake ID
// To help guarantee uniqueness
// - Make sure your system is keeping accurate system time
// - Make sure you never have multiple nodes running with the same node ID
func (s *Snowflake) Generate() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Since(s.epoch).Milliseconds()

	if now == s.time {
		s.step = (s.step + 1) & s.stepMask

		if s.step == 0 {
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
