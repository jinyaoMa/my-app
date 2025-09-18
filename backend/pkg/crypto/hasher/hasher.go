package hasher

import (
	"fmt"

	"majinyao.cn/my-app/backend/pkg/crypto/hasher/base/hash2"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/crc64"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/sha3"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/sm3"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/xxh3"
)

type IHasher interface {
	OutputLength() int
	Get() (h hash2.IHash2)
	Put(h hash2.IHash2) (err error)
	Using(handle func(h hash2.IHash2))
	Hash(data []byte) (sum []byte)
	HashBase64(data string) (checksum string)
	Verify(sum []byte, data []byte) (ok bool)
	VerifyBase64(checksum string, data string) (ok bool)
}

func MustNew(options Options) IHasher {
	h, err := New(options)
	if err != nil {
		panic(err)
	}
	return h
}

func New(options Options) (h IHasher, err error) {
	switch options.Alg {
	case crc64.Alg:
		h, err = crc64.New(options.toCrc64())
	case sha3.Alg:
		h, err = sha3.New(options.toSha3())
	case sm3.Alg:
		h, err = sm3.New(options.toSm3())
	case xxh3.Alg:
		h, err = xxh3.New(options.toXxh3())
	default:
		err = fmt.Errorf("hasher: unknown algorithm %s", options.Alg)
	}
	return
}
