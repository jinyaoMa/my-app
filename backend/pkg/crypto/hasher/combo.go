package hasher

import (
	"errors"
	"fmt"

	"majinyao.cn/my-app/backend/pkg/crypto/hasher/base/hash2"
)

type ICombo interface {
	OutputLengths() []int
	Get() (hs hash2.Hash2s)
	Put(hs hash2.Hash2s) (err error)
	Using(handle func(hs hash2.Hash2s))
	Hash(data []byte) (sums [][]byte)
	HashBase64(data string) (checksums []string)
	Verify(sums [][]byte, data []byte) (ok bool)
	VerifyBase64(checksums []string, data string) (ok bool)
}

func NewCombo(options ...Options) (c ICombo, err error) {
	hs := make(combo, 0, len(options))
	for _, opts := range options {
		h, errNew := New(opts)
		if errNew != nil {
			return nil, errors.Join(errors.New("crypto hasher combo: new failed"), errNew)
		}
		hs = append(hs, h)
	}

	c = hs
	return
}

type combo []IHasher

func (c combo) OutputLengths() []int {
	ls := make([]int, 0, len(c))
	for _, h := range c {
		ls = append(ls, h.OutputLength())
	}
	return ls
}

func (c combo) Get() (hs hash2.Hash2s) {
	hs = make(hash2.Hash2s, 0, len(c))
	for _, h := range c {
		hs = append(hs, h.Get())
	}
	return
}

func (c combo) Put(hs hash2.Hash2s) (err error) {
	if len(c) != len(hs) {
		return fmt.Errorf("crypto hasher combo: put expected %d hashes, got %d", len(c), len(hs))
	}
	for i := range c {
		err = c[i].Put(hs[i])
		if err != nil {
			return fmt.Errorf("crypto hasher combo: put failed, index %d, err %w", i, err)
		}
	}
	return
}

func (c combo) Using(handle func(hs hash2.Hash2s)) {
	hs := c.Get()
	defer c.Put(hs)
	handle(hs)
}

func (c combo) Hash(data []byte) (sums [][]byte) {
	c.Using(func(hs hash2.Hash2s) {
		hs.Write(data)
		sums = hs.Sum()
	})
	return
}

func (c combo) HashBase64(data string) (checksums []string) {
	c.Using(func(hs hash2.Hash2s) {
		hs.Write([]byte(data))
		checksums = hs.SumBase64()
	})
	return
}

func (c combo) Verify(sums [][]byte, data []byte) (ok bool) {
	c.Using(func(hs hash2.Hash2s) {
		hs.Write(data)
		ok = hs.Check(sums)
	})
	return
}

func (c combo) VerifyBase64(checksums []string, data string) (ok bool) {
	c.Using(func(hs hash2.Hash2s) {
		hs.Write([]byte(data))
		ok = hs.CheckBase64(checksums)
	})
	return
}
