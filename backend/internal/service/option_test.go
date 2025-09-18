package service_test

import (
	"context"
	"runtime"
	"testing"
	"time"

	"gorm.io/gorm/logger"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/internal/service"
	"majinyao.cn/my-app/backend/pkg/crypto/cipher"
	"majinyao.cn/my-app/backend/pkg/crypto/cipher/sm4"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/sm3"
	"majinyao.cn/my-app/backend/pkg/crypto/keygen"
	"majinyao.cn/my-app/backend/pkg/crypto/keygen/argon2"
	"majinyao.cn/my-app/backend/pkg/db"
	"majinyao.cn/my-app/backend/pkg/snowflake"
	"majinyao.cn/my-app/backend/pkg/test"
)

func TestOptionService(t *testing.T) {
	now := time.Now()
	defer test.LogTestingTime(t, now)

	epoch := time.Now()
	text := "mjy"
	threads := runtime.NumCPU()
	keyGen, _ := keygen.New(keygen.Options{
		Alg:       argon2.Alg,
		Salt:      text,
		Threads:   threads,
		KeyLength: 16,
	})
	ivGen, _ := keygen.New(keygen.Options{
		Alg:       argon2.Alg,
		Salt:      text,
		Threads:   threads,
		KeyLength: 12,
	})
	key := keyGen.DeriveBase64(text, false)
	iv := ivGen.DeriveBase64(text, false)

	tx, err := db.Open([]any{
		&entity.Option{},
	}, db.Options{
		Driver:   db.DrvSqlite,
		Dsn:      ":memory:?_pragma=foreign_keys(1)",
		LogLevel: logger.Info,
		Snowflake: snowflake.Options{
			Epoch:    epoch,
			NodeBits: 7,
			StepBits: 14,
			NodeId:   1,
		},
		Keygen: keygen.Options{
			Alg:       argon2.Alg,
			Salt:      text,
			Threads:   threads,
			KeyLength: 32,
		},
		Hasher: hasher.Options{
			Alg:  sm3.Alg,
			Salt: text,
			Key:  text,
		},
		Cipher: cipher.Options{
			Alg: sm4.Alg,
			AAD: text,
			Key: key,
			Iv:  iv,
		},
		AutoMigrate: true,
	})
	if err != nil {
		t.Fatal(err)
	}

	optionService, cancel := service.NewOptionService(context.Background(), tx)
	defer cancel()

	optTest := entity.MustNewOption("test", 1.0)
	err = optionService.LoadOrCreateByKey(optTest)
	if err != nil || optTest.GetFloat64() != 1.0 {
		t.Fatal(err)
	}

	optTest.SetBool(true)
	optionService.Update(optTest)
	if optTest.GetBool() != true {
		t.Fatal("update failed")
	}

	optTest2 := entity.MustNewOption("test", 1.0)
	err = optionService.LoadOrCreateByKey(optTest2)
	if err != nil || optTest2.GetBool() != true || optTest.Id != optTest2.Id {
		t.Fatal(err)
	}
}
