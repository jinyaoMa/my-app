package crc64_test

import (
	"testing"
	"time"

	"majinyao.cn/my-app/backend/pkg/crypto/hasher/base"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/base/hash2"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/crc64"
	"majinyao.cn/my-app/backend/pkg/test"
)

func TestCrc64(t *testing.T) {
	now := time.Now()
	defer test.LogTestingTime(t, now)

	text := "mjy"
	crc64, err := crc64.New(crc64.Options{
		Base: base.Options{
			Hash2: hash2.Options{
				Salt:   text,
				Prefix: crc64.Alg + "_",
			},
			Key: text,
		},
	})
	if err != nil {
		t.Fatalf("new crc64 error: %s", err)
	}

	checksum := crc64.HashBase64(text)
	ok := crc64.VerifyBase64(checksum, text)
	if !ok {
		t.Fatalf("verify checksum %s =!> text %s", checksum, text)
	}
}
