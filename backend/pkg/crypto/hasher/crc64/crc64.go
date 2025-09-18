package crc64

import (
	"errors"
	"hash"
	c2 "hash/crc64"

	"majinyao.cn/my-app/backend/pkg/crypto/hasher/base"
)

const Alg string = "crc64"

func MustNew(options Options) *crc64 {
	c, err := New(options)
	if err != nil {
		panic(err)
	}
	return c
}

func New(options Options) (*crc64, error) {
	return new(crc64).init(options)
}

type crc64 struct {
	base.Base
}

func (c *crc64) OutputLength() int {
	return c2.Size
}

func (c *crc64) init(options Options) (*crc64, error) {
	table := c2.MakeTable(c2.ISO)
	_, err := c.Base.Init(func() hash.Hash {
		return c2.New(table)
	}, Alg, options.Base)
	if err != nil {
		return nil, errors.Join(errors.New("crypto hasher crc64: init failed"), err)
	}
	return c, nil
}
