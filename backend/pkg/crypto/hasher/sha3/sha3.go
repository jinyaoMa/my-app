package sha3

import (
	s2 "crypto/sha3"
	"errors"
	"hash"

	"majinyao.cn/my-app/backend/pkg/crypto/hasher/base"
)

const Alg string = "sha3"

func MustNew(options Options) *sha3 {
	s, err := New(options)
	if err != nil {
		panic(err)
	}
	return s
}

func New(options Options) (*sha3, error) {
	return new(sha3).init(options)
}

type sha3 struct {
	base.Base
	bitLength int
}

func (s *sha3) OutputLength() int {
	return s.bitLength / 8
}

func (s *sha3) init(options Options) (*sha3, error) {
	if options.BitLength != 0 &&
		options.BitLength != 224 &&
		options.BitLength != 256 &&
		options.BitLength != 384 &&
		options.BitLength != 512 {
		return nil, errors.Join(errors.New("crypto hasher sha3: init failed, bit length must be 224, 256, 384 or 512"))
	}

	s.bitLength = options.BitLength
	_, err := s.Base.Init(func() hash.Hash {
		switch s.bitLength {
		case 224:
			return s2.New224()
		case 256:
			return s2.New256()
		case 384:
			return s2.New384()
		}
		return s2.New512()
	}, Alg, options.Base)
	if err != nil {
		return nil, errors.Join(errors.New("crypto hasher sha3: init failed"), err)
	}
	return s, nil
}
