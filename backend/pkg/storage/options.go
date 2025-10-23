package storage

import "majinyao.cn/my-app/backend/pkg/crypto/hasher"

// windows MaxPathLength may be 254 bytes
type Options struct {
	Libraries     []Library        `json:"libraries"`
	Temporary     string           `json:"temporary"`     // temporary folder name
	BufferSize    int64            `json:"bufferSize"`    // internal file io buffer size
	Hashers       []hasher.Options `json:"hashers"`       // hashers combo options
	MaxPathLength int              `json:"maxPathLength"` // max path length, default 2 * (hashers output length sum + 16)
}
