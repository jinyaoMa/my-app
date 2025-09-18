package snowflake

import (
	"fmt"
	"sync"
	"time"
)

type ISnowflake interface {
	Generate() int64
}

func MustNew(options Options) ISnowflake {
	s, err := New(options)
	if err != nil {
		panic(err)
	}
	return s
}

func New(options Options) (ISnowflake, error) {
	return new(snowflake).init(options)
}

type snowflake struct {
	mutex     sync.Mutex
	epoch     time.Time
	timeShift int
	nodeBit   int64
	stepMask  int64
	time      int64
	step      int64
}

func (s *snowflake) Generate() int64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()

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
	return (s.time << s.timeShift) | s.nodeBit | s.step
}

func (s *snowflake) init(options Options) (*snowflake, error) {
	maxShareBits := 22
	shareBits := options.NodeBits + options.StepBits
	if shareBits > maxShareBits {
		return nil, fmt.Errorf("showflake has only a max %d bits to share between node/step", maxShareBits)
	}

	nodeMax := -1 ^ (-1 << options.NodeBits)
	if options.NodeId < 0 || options.NodeId > nodeMax {
		return nil, fmt.Errorf("node id must be between 0 and %d", nodeMax)
	}

	s.epoch = options.Epoch
	s.timeShift = shareBits
	s.nodeBit = int64(options.NodeId) << options.StepBits
	s.stepMask = -1 ^ (-1 << options.StepBits)
	return s, nil
}
