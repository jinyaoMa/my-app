package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/md5"
	"errors"
)

// extract key and iv from given password and salt
func (saltedAES *SaltedAES) extract(salt []byte) (key []byte, iv []byte) {
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
		prevSum = md5.Sum(append(append(prevSum[:], saltedAES.password...), salt...))
		copy(md5Filling[i*16:], prevSum[:])
	}

	return md5Filling[:keyLength], md5Filling[keyLength:]
}

// pkcs7pad add pkcs7 padding
func (saltedAES *SaltedAES) pkcs7pad(data []byte) []byte {
	padLen := aes.BlockSize - len(data)%aes.BlockSize
	padding := bytes.Repeat([]byte{byte(padLen)}, padLen)
	return append(data, padding...)
}

// pkcs7unpad remove pkcs7 padding
func (saltedAES *SaltedAES) pkcs7unpad(data []byte) ([]byte, error) {
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
