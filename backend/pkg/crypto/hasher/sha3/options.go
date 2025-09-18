package sha3

import "majinyao.cn/my-app/backend/pkg/crypto/hasher/base"

type Options struct {
	Base      base.Options `json:"base"`
	BitLength int          `json:"bitLength"` // 224, 256, 384, 512 (default)
}
