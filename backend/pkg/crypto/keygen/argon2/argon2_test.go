package argon2_test

import (
	"runtime"
	"testing"
	"time"

	"majinyao.cn/my-app/backend/pkg/crypto/keygen/argon2"
	"majinyao.cn/my-app/backend/pkg/test"
)

func TestArgon2(t *testing.T) {
	now := time.Now()
	defer test.LogTestingTime(t, now)

	text := "mjy"
	threads := runtime.NumCPU()
	argon2 := argon2.New(argon2.Options{
		Salt:      text,
		Threads:   threads, // 1-254, not greater than num of cpu
		KeyLength: threads, // 2-512 bytes
		Prefix:    argon2.Alg + "_",
	})

	key := argon2.DeriveBase64(text, true)
	ok := argon2.VerifyBase64(key, text, true)
	if !ok {
		t.Fatalf("checksum %s =!> text %s", key, text)
	}
}
