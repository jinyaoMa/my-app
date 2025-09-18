package keygen_test

import (
	"runtime"
	"testing"
	"time"

	"majinyao.cn/my-app/backend/pkg/crypto/keygen"
	"majinyao.cn/my-app/backend/pkg/crypto/keygen/argon2"
	"majinyao.cn/my-app/backend/pkg/test"
)

func TestKeygen(t *testing.T) {
	now := time.Now()
	defer test.LogTestingTime(t, now)

	text := "mjy"
	threads := runtime.NumCPU()

	for _, alg := range []string{
		argon2.Alg,
	} {
		k, err := keygen.New(keygen.Options{
			Alg:       alg,
			Salt:      text,
			Threads:   threads,
			KeyLength: threads,
		})
		if err != nil {
			t.Fatal(err)
		}

		key := k.DeriveBase64(text, true)
		ok := k.VerifyBase64(key, text, true)
		if !ok {
			t.Fatalf("keygen %s checksum %s =!> text %s", alg, key, text)
		}
	}
}
