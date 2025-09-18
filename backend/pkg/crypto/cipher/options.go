package cipher

import (
	"majinyao.cn/my-app/backend/pkg/crypto/cipher/aes"
	"majinyao.cn/my-app/backend/pkg/crypto/cipher/base"
	"majinyao.cn/my-app/backend/pkg/crypto/cipher/sm4"
)

type Options struct {
	Alg string `json:"alg"`
	AAD string `json:"aad"`
	Key string `json:"key"`
	Iv  string `json:"iv"`
}

func (o *Options) toAes() aes.Options {
	return aes.Options{
		Base: base.Options{
			AAD:    o.AAD,
			Key:    o.Key,
			Iv:     o.Iv,
			Prefix: aes.Alg + "_",
		},
	}
}

func (o *Options) toSm4() sm4.Options {
	return sm4.Options{
		Base: base.Options{
			AAD:    o.AAD,
			Key:    o.Key,
			Iv:     o.Iv,
			Prefix: sm4.Alg + "_",
		},
	}
}
