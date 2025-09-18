package xxh3_test

import (
	"testing"
	"time"

	"majinyao.cn/my-app/backend/pkg/crypto/hasher/base"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/base/hash2"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/xxh3"
	"majinyao.cn/my-app/backend/pkg/test"
)

func TestXxh3(t *testing.T) {
	now := time.Now()
	defer test.LogTestingTime(t, now)

	text := "mjy"
	xxh3, err := xxh3.New(xxh3.Options{
		Base: base.Options{
			Hash2: hash2.Options{
				Salt:   text,
				Prefix: xxh3.Alg + "_",
			},
			Key: text,
		},
	})
	if err != nil {
		t.Fatalf("new xxh3 error: %s", err)
	}

	checksum := xxh3.HashBase64(text)
	ok := xxh3.VerifyBase64(checksum, text)
	if !ok {
		t.Fatalf("verify checksum %s =!> text %s", checksum, text)
	}
}
