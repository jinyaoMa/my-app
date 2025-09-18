package keygen

import (
	"fmt"

	"majinyao.cn/my-app/backend/pkg/crypto/keygen/argon2"
)

type IKeygen interface {
	Derive(password []byte) (key []byte)
	DeriveBase64(password string, hasPrefix ...bool) (key string)
	Verify(key []byte, password []byte) bool
	VerifyBase64(key string, password string, hasPrefix ...bool) bool
}

func MustNew(options Options) IKeygen {
	k, err := New(options)
	if err != nil {
		panic(err)
	}
	return k
}

func New(options Options) (k IKeygen, err error) {
	switch options.Alg {
	case argon2.Alg:
		k = argon2.New(options.toArgon2())
	default:
		err = fmt.Errorf("keygen: unknown algorithm %s", options.Alg)
	}
	return
}
