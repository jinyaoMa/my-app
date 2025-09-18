package keygen

import "majinyao.cn/my-app/backend/pkg/crypto/keygen/argon2"

type Options struct {
	Alg       string `json:"alg"`
	Salt      string `json:"salt"`
	Threads   int    `json:"threads"`
	KeyLength int    `json:"keyLength"`
}

func (o *Options) toArgon2() argon2.Options {
	return argon2.Options{
		Salt:      o.Salt,
		Threads:   o.Threads,
		KeyLength: o.KeyLength,
		Prefix:    argon2.Alg + "_",
	}
}
