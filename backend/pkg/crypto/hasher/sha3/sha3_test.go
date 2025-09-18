package sha3_test

import (
	"testing"
	"time"

	"majinyao.cn/my-app/backend/pkg/crypto/hasher/base"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/base/hash2"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/sha3"
	"majinyao.cn/my-app/backend/pkg/test"
)

func TestSha3(t *testing.T) {
	now := time.Now()
	defer test.LogTestingTime(t, now)

	text := "mjy"
	sha3, err := sha3.New(sha3.Options{
		Base: base.Options{
			Hash2: hash2.Options{
				Salt:   text,
				Prefix: sha3.Alg + "_",
			},
			Key: text,
		},
		BitLength: 256,
	})
	if err != nil {
		t.Fatalf("new sha3 error: %s", err)
	}

	checksum := sha3.HashBase64(text)
	ok := sha3.VerifyBase64(checksum, text)
	if !ok {
		t.Fatalf("verify checksum %s =!> text %s", checksum, text)
	}
}
