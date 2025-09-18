package cipher_test

import (
	"runtime"
	"testing"
	"time"

	"majinyao.cn/my-app/backend/pkg/crypto/cipher"
	"majinyao.cn/my-app/backend/pkg/crypto/cipher/aes"
	"majinyao.cn/my-app/backend/pkg/crypto/cipher/sm4"
	"majinyao.cn/my-app/backend/pkg/crypto/keygen/argon2"
	"majinyao.cn/my-app/backend/pkg/test"
)

func TestCipher(t *testing.T) {
	now := time.Now()
	defer test.LogTestingTime(t, now)

	text := "mjy"
	threads := runtime.NumCPU()

	for _, opt := range []struct {
		Alg       string
		KeyLength int
		IvLength  int
	}{
		{
			Alg:       aes.Alg,
			KeyLength: 32,
			IvLength:  12,
		},
		{
			Alg:       sm4.Alg,
			KeyLength: 16,
			IvLength:  12,
		},
	} {
		argon2Key := argon2.New(argon2.Options{
			Salt:      text,
			Threads:   threads,       // 1-254, not greater than num of cpu
			KeyLength: opt.KeyLength, // 2-512 bytes
		})
		argon2Iv := argon2.New(argon2.Options{
			Salt:      text,
			Threads:   threads,      // 1-254, not greater than num of cpu
			KeyLength: opt.IvLength, // 2-512 bytes
		})
		key := argon2Key.DeriveBase64(text, false)
		iv := argon2Iv.DeriveBase64(text, false)

		c, err := cipher.New(cipher.Options{
			Alg: opt.Alg,
			AAD: text,
			Key: key,
			Iv:  iv,
		})
		if err != nil {
			t.Fatalf("new cipher %s failed: %v", opt.Alg, err)
		}

		ciphertexts := c.EncryptBase64s([]string{text})
		plaintexts, err := c.DecryptBase64s(ciphertexts)
		if err != nil {
			t.Fatalf("decrypt cipher %s text %v failed: %v", opt.Alg, ciphertexts, err)
		}
		if len(plaintexts) != 1 {
			t.Fatalf("cipher %s plaintexts len %v != 1", opt.Alg, len(plaintexts))
		}
		if plaintexts[0] != text {
			t.Fatalf("cipher %s plaintext %v != text %v", opt.Alg, plaintexts[0], text)
		}
	}
}
