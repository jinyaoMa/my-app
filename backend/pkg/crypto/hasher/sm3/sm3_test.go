package sm3_test

import (
	"testing"
	"time"

	"majinyao.cn/my-app/backend/pkg/crypto/hasher/base"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/base/hash2"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/sm3"
	"majinyao.cn/my-app/backend/pkg/test"
)

func TestSm3(t *testing.T) {
	now := time.Now()
	defer test.LogTestingTime(t, now)

	text := "mjy"
	sm3, err := sm3.New(sm3.Options{
		Base: base.Options{
			Hash2: hash2.Options{
				Salt:   text,
				Prefix: sm3.Alg + "_",
			},
			Key: text,
		},
	})
	if err != nil {
		t.Fatalf("new sm3 error: %s", err)
	}

	checksum := sm3.HashBase64(text)
	ok := sm3.VerifyBase64(checksum, text)
	if !ok {
		t.Fatalf("verify checksum %s =!> text %s", checksum, text)
	}
}
