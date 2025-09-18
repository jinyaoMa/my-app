package base

import (
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

type Base struct {
	aad          []byte
	iv           []byte
	prefix       []byte
	prefixLength int
	aead         cipher.AEAD
}

func (b *Base) Encrypt(plaindata []byte) (cipherdata []byte) {
	return b.aead.Seal(nil, b.iv, plaindata, b.aad)
}

func (b *Base) EncryptBase64(plaintext string) (ciphertext string) {
	cipherdata := b.Encrypt([]byte(plaintext))
	return base64.RawURLEncoding.EncodeToString(append(b.prefix, cipherdata...))
}

func (b *Base) EncryptBase64s(plaintexts []string) (ciphertexts []string) {
	for _, plaintext := range plaintexts {
		ciphertexts = append(ciphertexts, b.EncryptBase64(plaintext))
	}
	return
}

func (b *Base) Decrypt(cipherdata []byte) (plaindata []byte, err error) {
	return b.aead.Open(nil, b.iv, cipherdata, b.aad)
}

func (b *Base) DecryptBase64(ciphertext string) (plaintext string, err error) {
	cipherdata, err := base64.RawURLEncoding.DecodeString(ciphertext)
	if err != nil {
		return
	}
	plaindata, err := b.Decrypt(cipherdata[b.prefixLength:])
	if err != nil {
		return
	}
	return string(plaindata), nil
}

func (b *Base) DecryptBase64s(ciphertexts []string) (plaintexts []string, err error) {
	for _, ciphertext := range ciphertexts {
		plaintext, errDecrypt := b.DecryptBase64(ciphertext)
		if errDecrypt == nil {
			plaintexts = append(plaintexts, plaintext)
		} else {
			plaintexts = append(plaintexts, ciphertext)
			err = errors.Join(errDecrypt, err)
		}
	}
	return
}

func (b *Base) Init(h func(key []byte) (cipher.Block, error), alg string, options Options) (*Base, error) {
	key, err := base64.RawURLEncoding.DecodeString(options.Key)
	if err != nil {
		return nil, errors.Join(errors.New("crypto cipher base: invalid key"), err)
	}

	iv, err := base64.RawURLEncoding.DecodeString(options.Iv)
	if err != nil {
		return nil, errors.Join(errors.New("crypto cipher base: invalid iv"), err)
	}

	block, err := h(key)
	if err != nil {
		return nil, errors.Join(errors.New("crypto cipher base: new block failed"), err)
	}

	b.aead, err = cipher.NewGCMWithNonceSize(block, len(iv))
	if err != nil {
		return nil, errors.Join(errors.New("crypto cipher base: new aead failed"), err)
	}

	prefix := []byte(options.Prefix)
	if len(prefix) == 0 {
		prefix = []byte(alg + "_")
	}

	b.aad = []byte(options.AAD)
	b.iv = iv
	b.prefix = prefix
	b.prefixLength = len(prefix)
	return b, nil
}
