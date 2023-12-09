package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
)

type SaltedAES struct {
	password         []byte
	saltHeader       string
	saltHeaderLength int
}

// Decrypt implements ICrypto
func (saltedAES *SaltedAES) Decrypt(ciphertext string) (string, error) {
	ct, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	// get prefix (salt header + salt),
	// salt uses same length as salt header's, so saltedAES.saltHeaderLength*2
	prefix := ct[:saltedAES.saltHeaderLength*2]
	// salt header check
	saltHeader := []byte(prefix[:saltedAES.saltHeaderLength])
	if !bytes.Equal(saltHeader, []byte(saltedAES.saltHeader)) {
		return "", errors.New("check cbc fixed header error")
	}

	// extract key and iv
	salt := []byte(prefix[saltedAES.saltHeaderLength:])
	key, iv := saltedAES.extract(salt)

	// get padded ciphertext (prefix removed)
	padded := ct[saltedAES.saltHeaderLength*2:]
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
	unpadded, err := saltedAES.pkcs7unpad(padded)
	if err != nil {
		return "", err
	}

	return string(unpadded), nil
}

// Encrypt implements ICrypto
func (saltedAES *SaltedAES) Encrypt(plaintext string) (string, error) {
	// generate random salt, use same length as salt header's
	salt := make([]byte, saltedAES.saltHeaderLength)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return "", err
	}

	// extract key and iv
	key, iv := saltedAES.extract(salt)

	// get padded plaintext
	padded := saltedAES.pkcs7pad([]byte(plaintext))

	// encrypt
	cb, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	cipher.NewCBCEncrypter(cb, iv).
		CryptBlocks(padded, padded)

	// prefix => salt header + salt
	prefix := append([]byte(saltedAES.saltHeader), salt...)

	return hex.EncodeToString(append(prefix, padded...)), nil
}

func NewAesWithSalt(password string, saltHeaders ...string) (saltedAES *SaltedAES, iCrypto ICrypto) {
	if len(saltHeaders) == 0 {
		saltHeaders = append(saltHeaders, "Salted_By_My_App__")
	}
	saltedAES = &SaltedAES{
		password:         []byte(password),
		saltHeader:       saltHeaders[0],
		saltHeaderLength: len([]byte(saltHeaders[0])),
	}
	return saltedAES, saltedAES
}
