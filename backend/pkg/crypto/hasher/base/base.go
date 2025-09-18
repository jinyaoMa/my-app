package base

import (
	"crypto/hmac"
	"encoding/base64"
	"errors"
	"fmt"
	"hash"
	"sync"

	"majinyao.cn/my-app/backend/pkg/crypto/hasher/base/hash2"
)

type Base struct {
	alg    string
	serial string
	pool   sync.Pool
}

func (b *Base) Get() (h hash2.IHash2) {
	return b.pool.Get().(hash2.IHash2)
}

func (b *Base) Put(h hash2.IHash2) (err error) {
	if h == nil {
		return fmt.Errorf("crypto hasher %s: nil put hash", b.alg)
	}
	if !h.CheckSerial(b.serial) {
		return fmt.Errorf("crypto hasher %s: expected serial %s, got %s", b.alg, b.serial, h.Serial())
	}

	h.Reset()
	b.pool.Put(h)
	return
}

func (b *Base) Using(handle func(h hash2.IHash2)) {
	h := b.Get()
	defer b.Put(h)
	handle(h)
}

func (b *Base) Hash(data []byte) (sum []byte) {
	b.Using(func(h hash2.IHash2) {
		h.Write(data)
		sum = h.Sum()
	})
	return
}

func (b *Base) HashBase64(data string) (checksum string) {
	b.Using(func(h hash2.IHash2) {
		h.Write([]byte(data))
		checksum = h.SumBase64()
	})
	return
}

func (b *Base) Verify(sum []byte, data []byte) (ok bool) {
	b.Using(func(h hash2.IHash2) {
		h.Write(data)
		ok = h.Check(sum)
	})
	return
}

func (b *Base) VerifyBase64(checksum string, data string) (ok bool) {
	b.Using(func(h hash2.IHash2) {
		h.Write([]byte(data))
		ok = h.CheckBase64(checksum)
	})
	return
}

func (c *Base) Init(h func() hash.Hash, alg string, options Options) (*Base, error) {
	key, err := base64.RawURLEncoding.DecodeString(options.Key)
	if err != nil {
		return nil, errors.Join(errors.New("crypto hasher base: invalid key"), err)
	}

	var h2 hash2.IHash2
	if len(key) > 0 {
		h2 = hash2.New(hmac.New(h, key), alg, options.Hash2)
		c.pool = sync.Pool{
			New: func() any {
				return hash2.New(hmac.New(h, key), alg, options.Hash2)
			},
		}
	} else {
		h2 = hash2.New(h(), alg, options.Hash2)
		c.pool = sync.Pool{
			New: func() any {
				return hash2.New(h(), alg, options.Hash2)
			},
		}
	}

	c.alg = h2.Alg()
	c.serial = h2.Serial()
	c.pool.Put(h2)
	return c, nil
}
