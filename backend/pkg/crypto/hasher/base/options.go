package base

import "majinyao.cn/my-app/backend/pkg/crypto/hasher/base/hash2"

type Options struct {
	Hash2 hash2.Options `json:"hash2"`
	Key   string        `json:"key"` // key is used for HMAC in base64
}
