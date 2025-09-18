package cipher

import (
	"fmt"

	"majinyao.cn/my-app/backend/pkg/crypto/cipher/aes"
	"majinyao.cn/my-app/backend/pkg/crypto/cipher/sm4"
)

type ICipher interface {
	Encrypt(plaindata []byte) (cipherdata []byte)
	EncryptBase64(plaintext string) (ciphertext string)
	EncryptBase64s(plaintexts []string) (ciphertexts []string)
	Decrypt(cipherdata []byte) (plaindata []byte, err error)
	DecryptBase64(ciphertext string) (plaintext string, err error)
	DecryptBase64s(ciphertexts []string) (plaintexts []string, err error)
}

func MustNew(options Options) ICipher {
	c, err := New(options)
	if err != nil {
		panic(err)
	}
	return c
}

func New(options Options) (c ICipher, err error) {
	switch options.Alg {
	case aes.Alg:
		c, err = aes.New(options.toAes())
	case sm4.Alg:
		c, err = sm4.New(options.toSm4())
	default:
		err = fmt.Errorf("cipher: unknown algorithm %s", options.Alg)
	}
	return
}
