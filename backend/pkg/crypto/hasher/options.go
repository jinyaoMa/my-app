package hasher

import (
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/base"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/base/hash2"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/crc64"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/sha3"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/sm3"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/xxh3"
)

type Options struct {
	Alg       string `json:"alg"`
	Salt      string `json:"salt"`
	Key       string `json:"key"`       // key is used for HMAC in base64
	BitLength int    `json:"bitLength"` // 224, 256, 384, 512 (default)
}

func (o *Options) toCrc64() crc64.Options {
	return crc64.Options{
		Base: base.Options{
			Hash2: hash2.Options{
				Salt:   o.Salt,
				Prefix: crc64.Alg + "_",
			},
			Key: o.Key,
		},
	}
}

func (o *Options) toSha3() sha3.Options {
	return sha3.Options{
		Base: base.Options{
			Hash2: hash2.Options{
				Salt:   o.Salt,
				Prefix: sha3.Alg + "_",
			},
			Key: o.Key,
		},
		BitLength: o.BitLength,
	}
}

func (o *Options) toSm3() sm3.Options {
	return sm3.Options{
		Base: base.Options{
			Hash2: hash2.Options{
				Salt:   o.Salt,
				Prefix: sm3.Alg + "_",
			},
			Key: o.Key,
		},
	}
}

func (o *Options) toXxh3() xxh3.Options {
	return xxh3.Options{
		Base: base.Options{
			Hash2: hash2.Options{
				Salt:   o.Salt,
				Prefix: xxh3.Alg + "_",
			},
			Key: o.Key,
		},
	}
}
