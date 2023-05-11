package snowflake

import (
	"fmt"
	"my-app/backend/pkg/snowflake/interfaces"
	"my-app/backend/pkg/snowflake/options"
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

// Default return Snowflake Id generator with default options
func Default() (interfaces.ISnowflake, error) {
	return NewSnowflake(options.DefaultOSnowflake())
}

// NewSnowflake return Snowflake Id generator with custom options
func NewSnowflake(opts *options.OSnowflake) (interfaces.ISnowflake, error) {
	opts = options.NewOSnowflake(opts)

	var shareBits uint8 = opts.NodeBits + opts.StepBits
	if shareBits > options.TotalShareBits {
		return nil, fmt.Errorf("remember, you have a total %d bits to share between node/step", options.TotalShareBits)
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

// Generate implements IUtility
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
