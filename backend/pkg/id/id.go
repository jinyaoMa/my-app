package id

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type ID struct {
	config *Config

	mu    sync.Mutex
	epoch time.Time
	time  int64
	node  int64
	step  int64

	nodeMax   int64
	nodeMask  int64
	stepMask  int64
	timeShift uint
	nodeShift uint
}

// Generate implements IID
func (id *ID) Generate() int64 {
	id.mu.Lock()
	defer id.mu.Unlock()

	now := time.Since(id.epoch).Milliseconds()

	if now == id.time {
		id.step = (id.step + 1) & id.stepMask

		if id.step == 0 { // step bits go over, move to next timestamp
			for now <= id.time {
				now = time.Since(id.epoch).Milliseconds()
			}
		}
	} else {
		id.step = 0
	}

	id.time = now

	r := int64((now)<<id.timeShift |
		(id.node << id.nodeShift) |
		(id.step),
	)

	return r
}

// New return Snowflake Id generator with custom options
func New(cfg *Config) (*ID, error) {
	cfg = NewConfig(cfg)

	var shareBits uint = cfg.NodeBits + cfg.StepBits
	if shareBits > TotalShareBits {
		return nil, fmt.Errorf("remember, you have a total %d bits to share between node/step", TotalShareBits)
	}

	var nodeMax int64 = -1 ^ (-1 << cfg.NodeBits)
	if cfg.NodeNumber < 0 || cfg.NodeNumber > nodeMax {
		return nil, fmt.Errorf("node number must be between 0 and %s", strconv.FormatInt(nodeMax, 10))
	}

	return &ID{
		config:    cfg,
		epoch:     cfg.Epoch,
		node:      cfg.NodeNumber,
		nodeMax:   nodeMax,
		nodeMask:  nodeMax << cfg.StepBits,
		stepMask:  -1 ^ (-1 << cfg.StepBits),
		timeShift: shareBits,
		nodeShift: cfg.StepBits,
	}, nil
}

// Default return Snowflake Id generator with default options
func Default() (*ID, error) {
	return New(DefaultConfig())
}

func NewIID(cfg *Config) (IID, error) {
	return New(cfg)
}
