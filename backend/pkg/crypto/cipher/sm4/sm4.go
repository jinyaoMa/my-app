package sm4

import (
	"crypto/cipher"
	"errors"

	s2 "github.com/emmansun/gmsm/sm4"
	"majinyao.cn/my-app/backend/pkg/crypto/cipher/base"
)

const Alg string = "sm4"

func MustNew(options Options) *sm4 {
	s, err := New(options)
	if err != nil {
		panic(err)
	}
	return s
}

func New(options Options) (*sm4, error) {
	return new(sm4).init(options)
}

type sm4 struct {
	base.Base
}

func (s *sm4) init(options Options) (*sm4, error) {
	_, err := s.Base.Init(func(key []byte) (cipher.Block, error) {
		return s2.NewCipher(key)
	}, Alg, options.Base)
	if err != nil {
		return nil, errors.Join(errors.New("crypto cipher sm4: init failed"), err)
	}
	return s, nil
}
