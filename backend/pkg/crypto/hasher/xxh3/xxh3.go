package xxh3

import (
	"errors"
	"hash"

	x2 "github.com/zeebo/xxh3"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/base"
)

const Alg string = "xxh3"

func MustNew(options Options) *xxh3 {
	x, err := New(options)
	if err != nil {
		panic(err)
	}
	return x
}

func New(options Options) (*xxh3, error) {
	return new(xxh3).init(options)
}

type xxh3 struct {
	base.Base
}

func (x *xxh3) OutputLength() int {
	return 8
}

func (x *xxh3) init(options Options) (*xxh3, error) {
	_, err := x.Base.Init(func() hash.Hash {
		return x2.New()
	}, Alg, options.Base)
	if err != nil {
		return nil, errors.Join(errors.New("crypto hasher xxh3: init failed"), err)
	}
	return x, nil
}
