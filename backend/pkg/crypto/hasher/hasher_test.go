package hasher_test

import (
	"strconv"
	"testing"
	"time"

	"majinyao.cn/my-app/backend/pkg/crypto/hasher"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/crc64"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/sha3"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/sm3"
	"majinyao.cn/my-app/backend/pkg/test"
)

func TestHasher(t *testing.T) {
	now := time.Now()
	defer test.LogTestingTime(t, now)

	text := "mjy"
	for _, alg := range []string{
		crc64.Alg,
		sha3.Alg,
		sm3.Alg,
	} {
		h, err := hasher.New(hasher.Options{
			Alg:  alg,
			Salt: text,
			Key:  text, // key is used for HMAC in base64
		})
		if err != nil {
			t.Fatalf("new hasher %s failed: %v", alg, err)
		}

		checksums := make(chan string, 5)
		for i := 0; i < 5; i++ {
			go func(index int) {
				data := text + strconv.Itoa(index)
				checksum := h.HashBase64(data)
				ok := h.VerifyBase64(checksum, data)
				if !ok {
					t.Errorf("hasher %s verify checksum %s =!> text %s", alg, checksum, data)
				}
				checksums <- checksum
			}(i)
		}

		checksumMap := make(map[string]struct{})
		for {
			select {
			case checksum := <-checksums:
				if _, exist := checksumMap[checksum]; exist {
					t.Fatal("there should not be duplicate checksums")
				}
				checksumMap[checksum] = struct{}{}
			default:
				return
			}
		}
	}
}
