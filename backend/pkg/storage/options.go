package storage

import "majinyao.cn/my-app/backend/pkg/crypto/hasher"

const (
	B  uint64 = 1
	KB uint64 = 1024 * B
	MB uint64 = 1024 * KB
	GB uint64 = 1024 * MB
	TB uint64 = 1024 * GB
	PB uint64 = 1024 * TB
	EB uint64 = 1024 * PB
)

// windows MaxPathLength may be 254 bytes
type Options struct {
	Libraries     []Library        `json:"libraries"`
	Temporary     string           `json:"temporary"`     // temporary folder name
	BufferSize    uint64           `json:"bufferSize"`    // internal file io buffer size
	Hashers       []hasher.Options `json:"hashers"`       // hashers combo options
	MaxPathLength int              `json:"maxPathLength"` // max path length, default 2 * (hashers output length sum + 16)
}
