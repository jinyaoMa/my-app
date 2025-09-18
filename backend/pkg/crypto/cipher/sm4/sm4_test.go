package sm4_test

import (
	"runtime"
	"testing"
	"time"

	"majinyao.cn/my-app/backend/pkg/crypto/cipher/base"
	"majinyao.cn/my-app/backend/pkg/crypto/cipher/sm4"
	"majinyao.cn/my-app/backend/pkg/crypto/keygen/argon2"

	"majinyao.cn/my-app/backend/pkg/test"
)

func TestSm4(t *testing.T) {
	now := time.Now()
	defer test.LogTestingTime(t, now)

	text := "mjy"
	threads := runtime.NumCPU()
	argon2Key := argon2.New(argon2.Options{
		Salt:      text,
		Threads:   threads, // 1-254, not greater than num of cpu
		KeyLength: 16,      // 2-512 bytes
	})
	argon2Iv := argon2.New(argon2.Options{
		Salt:      text,
		Threads:   threads, // 1-254, not greater than num of cpu
		KeyLength: 12,      // 2-512 bytes
	})
	key := argon2Key.DeriveBase64(text, false)
	iv := argon2Iv.DeriveBase64(text, false)

	sm4, err := sm4.New(sm4.Options{
		Base: base.Options{
			AAD:    text,          // additional authenticated data
			Key:    key,           // in base64 w/o prefix
			Iv:     iv,            // in base64 w/o prefix
			Prefix: sm4.Alg + "_", // prefix of base64 key
		},
	})
	if err != nil {
		t.Fatalf("new sm4 failed: %v", err)
	}

	ciphertexts := sm4.EncryptBase64s([]string{text})
	plaintexts, err := sm4.DecryptBase64s(ciphertexts)
	if err != nil {
		t.Fatalf("decrypt ciphertexts %v failed: %v", ciphertexts, err)
	}
	if len(plaintexts) != 1 {
		t.Fatalf("plaintexts len %v != 1", len(plaintexts))
	}
	if plaintexts[0] != text {
		t.Fatalf("plaintext %v != text %v", plaintexts[0], text)
	}
}
