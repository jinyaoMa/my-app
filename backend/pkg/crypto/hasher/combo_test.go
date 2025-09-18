package hasher_test

import (
	"testing"
	"time"

	"majinyao.cn/my-app/backend/pkg/crypto/hasher"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/crc64"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/sha3"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/sm3"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/xxh3"
	"majinyao.cn/my-app/backend/pkg/test"
)

func TestCombo(t *testing.T) {
	now := time.Now()
	defer test.LogTestingTime(t, now)

	text := "mjy"

	c, err := hasher.NewCombo(hasher.Options{
		Alg:  crc64.Alg,
		Salt: text,
		Key:  text, // key is used for HMAC in base64
	}, hasher.Options{
		Alg:  sha3.Alg,
		Salt: text,
		Key:  text, // key is used for HMAC in base64
	}, hasher.Options{
		Alg:  sm3.Alg,
		Salt: text,
		Key:  text, // key is used for HMAC in base64
	}, hasher.Options{
		Alg:  xxh3.Alg,
		Salt: text,
		Key:  text, // key is used for HMAC in base64
	})
	if err != nil {
		t.Fatalf("new combo failed: %v", err)
	}

	checksums := c.HashBase64(text)
	ok := c.VerifyBase64(checksums, text)
	if !ok {
		t.Fatalf("combo verify checksum %v =!> text %v", checksums, text)
	}
}
