package aes

import (
	a2 "crypto/aes"
	"crypto/cipher"
	"errors"

	"majinyao.cn/my-app/backend/pkg/crypto/cipher/base"
)

const Alg string = "aes"

func MustNew(options Options) *aes {
	a, err := New(options)
	if err != nil {
		panic(err)
	}
	return a
}

func New(options Options) (*aes, error) {
	return new(aes).init(options)
}

type aes struct {
	base.Base
}

func (a *aes) init(options Options) (*aes, error) {
	_, err := a.Base.Init(func(key []byte) (cipher.Block, error) {
		return a2.NewCipher(key)
	}, Alg, options.Base)
	if err != nil {
		return nil, errors.Join(errors.New("crypto cipher aes: init failed"), err)
	}
	return a, nil
}
