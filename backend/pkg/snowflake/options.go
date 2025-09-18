package snowflake

import "time"

type Options struct {
	// started timestamp in millisecond, store into at least 42 bits (~139 years in range, 70 years with negative signed)
	Epoch time.Time `json:"epoch"`

	// bits to store nodes ids, e.g. 10 => max 1024 nodes (max 22 bits to share between node/step)
	NodeBits int `json:"nodeBits"`

	// bits to store increment ids, e.g. 12 => max 4096 ids (max 22 bits to share between node/step)
	StepBits int `json:"stepBits"`

	// current node id/number, e.g. if max 4096 ids/moment, then min node id is 0 and max node id is 1023
	NodeId int `json:"nodeId"`
}
