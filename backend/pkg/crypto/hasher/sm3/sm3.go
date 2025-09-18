package sm3

import (
	"errors"
	"hash"

	s2 "github.com/emmansun/gmsm/sm3"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/base"
)

const Alg string = "sm3"

func MustNew(options Options) *sm3 {
	s, err := New(options)
	if err != nil {
		panic(err)
	}
	return s
}

func New(options Options) (*sm3, error) {
	return new(sm3).init(options)
}

type sm3 struct {
	base.Base
}

func (s *sm3) OutputLength() int {
	return s2.Size
}

func (s *sm3) init(options Options) (*sm3, error) {
	_, err := s.Base.Init(func() hash.Hash {
		return s2.New()
	}, Alg, options.Base)
	if err != nil {
		return nil, errors.Join(errors.New("crypto hasher sm3: init failed"), err)
	}
	return s, nil
}
