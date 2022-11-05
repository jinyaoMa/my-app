package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
)

const (
	AesSaltHeader       = "Salted_By_My_App__"
	AesSaltHeaderLength = len(AesSaltHeader)
)

type AesWithSalt struct {
	password []byte
}

// NewAesWithSalt salted AES algorithm constructor
func NewAesWithSalt(password string) *AesWithSalt {
	return &AesWithSalt{
		password: []byte(password),
	}
}

// Encrypt encrypt plaintext to hexadecimal encoded ciphertext
func (aws *AesWithSalt) Encrypt(plaintext string) (string, error) {
	// generate random salt, use same length as salt header's
	salt := [AesSaltHeaderLength]byte{}
	_, err := io.ReadFull(rand.Reader, salt[:])
	if err != nil {
		return "", err
	}

	// extract key and iv
	key, iv := aws.extract(salt[:])

	// get padded plaintext
	padded := aws.pkcs7pad([]byte(plaintext))

	// encrypt
	cb, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	cipher.NewCBCEncrypter(cb, iv).
		CryptBlocks(padded, padded)

	// prefix => salt header + salt
	prefix := append([]byte(AesSaltHeader), salt[:]...)

	return hex.EncodeToString(append(prefix, padded...)), nil
}

// Decrypt decrypt hexadecimal encoded ciphertext to plaintext
func (aws *AesWithSalt) Decrypt(ciphertext string) (string, error) {
	ct, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	// get prefix (salt header + salt),
	// salt uses same length as salt header's, so AesSaltHeaderLength*2
	prefix := ct[:AesSaltHeaderLength*2]
	// salt header check
	saltHeader := []byte(prefix[:AesSaltHeaderLength])
	if !bytes.Equal(saltHeader, []byte(AesSaltHeader)) {
		return "", errors.New("check cbc fixed header error")
	}

	// extract key and iv
	salt := []byte(prefix[AesSaltHeaderLength:])
	key, iv := aws.extract(salt)

	// get padded ciphertext (prefix removed)
	padded := ct[AesSaltHeaderLength*2:]
	if len(padded) == 0 || len(padded)%aes.BlockSize != 0 {
		return "", errors.New("data is empty or not block-aligned")
	}

	// decrypt
	cb, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	// decrypt ciphertext only (prefix stays the same)
	cipher.NewCBCDecrypter(cb, iv).
		CryptBlocks(padded, padded)

	// get unpadded plaintext
	unpadded, err := aws.pkcs7unpad(padded)
	if err != nil {
		return "", err
	}

	return string(unpadded), nil
}

// extract key and iv from given password and salt
func (aws *AesWithSalt) extract(salt []byte) (key []byte, iv []byte) {
	// AES 256 KEY length => 32 bytes => 2 * block size => 2 * md5 checksum length
	keyLength := aes.BlockSize * 2
	// IV legth => 16 bytes => 1 * block size => 1 * md5 checksum length
	ivLength := aes.BlockSize
	// md5 filling length should be multiple of md5 checksum size in order to fill KEY and IV,
	md5Filling := make([]byte, keyLength+ivLength)

	var prevSum [16]byte // 16 => md5 checksum length
	// len(md5Filling)/16 => number of md5 checksums to fill 32-byte key and 16-byte iv
	for i := 0; i < len(md5Filling)/16; i++ {
		// next checksum = previous md5 checksum (empty at 1st time) + password + salt
		prevSum = md5.Sum(append(append(prevSum[:], aws.password...), salt...))
		copy(md5Filling[i*16:], prevSum[:])
	}

	return md5Filling[:keyLength], md5Filling[keyLength:]
}

// pkcs7pad add pkcs7 padding
func (aws *AesWithSalt) pkcs7pad(data []byte) []byte {
	padLen := aes.BlockSize - len(data)%aes.BlockSize
	padding := bytes.Repeat([]byte{byte(padLen)}, padLen)
	return append(data, padding...)
}

// pkcs7unpad remove pkcs7 padding
func (aws *AesWithSalt) pkcs7unpad(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("pkcs7: data is empty")
	}
	if length%aes.BlockSize != 0 {
		return nil, errors.New("pkcs7: data is not block-aligned")
	}
	padLen := int(data[length-1])
	ref := bytes.Repeat([]byte{byte(padLen)}, padLen)
	if padLen > aes.BlockSize || padLen == 0 || !bytes.HasSuffix(data, ref) {
		return nil, errors.New("pkcs7: invalid padding")
	}
	return data[:length-padLen], nil
}
