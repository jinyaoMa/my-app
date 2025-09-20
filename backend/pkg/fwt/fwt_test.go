package fwt_test

import (
	"testing"
	"time"

	"github.com/apache/fory/go/fory"
	"majinyao.cn/my-app/backend/pkg/codegen"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher/sha3"
	"majinyao.cn/my-app/backend/pkg/fwt"
	"majinyao.cn/my-app/backend/pkg/snowflake"
	"majinyao.cn/my-app/backend/pkg/test"
	"majinyao.cn/my-app/backend/pkg/utils"
)

type UserData struct {
	UserId int64
	Nums   []int64
}

func (u UserData) GetIdentity() string {
	return utils.ConvertInt64ToB36(u.UserId)
}

func (u UserData) GetAgentId() string {
	return "test"
}

func TestFwt(t *testing.T) {
	now := time.Now()
	defer test.LogTestingTime(t, now)

	key := "mjy"
	f, err := fwt.New[UserData](fwt.Options{
		Snowflake: snowflake.Options{
			Epoch:    now,
			NodeBits: 7,
			StepBits: 14,
			NodeId:   1,
		},
		Hasher: hasher.Options{
			Alg:  sha3.Alg,
			Salt: key,
			Key:  key,
		},
		Codegen: codegen.Options{
			Characters: codegen.Digits + codegen.Letters + codegen.ULetters,
		},
		Issuer:        "mjy",
		Subject:       "test",
		Epoch:         now,
		ExpiredAge:    5,
		RefreshAge:    10,
		RefreshLength: 32,
	}, func(f *fory.Fory) error {
		f.RegisterTagType("UserData", UserData{})
		return nil
	})
	if err != nil {
		t.Fatalf("New failed: %v", err)
	}

	userdata := UserData{
		UserId: 123,
		Nums:   []int64{4, 5, 6},
	}
	accessToken, refreshToken, expiredAt, err := f.Generate(userdata)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}
	if expiredAt.Before(now.Add(time.Second)) {
		t.Fatalf("expiredAt.Before(now.Add(time.Second))")
	}

	userClaims, err := f.ParseAccessToken(accessToken)
	if err != nil {
		t.Fatalf("ParseAccessToken failed: %v", err)
	}
	if userClaims.Issuer != "mjy" || userClaims.Subject != "test" || userClaims.Data.UserId != 123 || userClaims.Data.Nums[1] != 5 {
		t.Fatalf("claims.Issuer != \"mjy\" || claims.Subject != \"test\" || claims.UserData.UserId != 123 || userClaims.Data.Nums[1] != 5")
	}

	err = f.ValidateClaims(userClaims)
	if err != nil {
		t.Fatalf("ValidateClaims failed: %v", err)
	}

	newAccessToken, newRefreshToken, expiredAt, err := f.Refresh(userClaims, refreshToken, func(data UserData) (newData UserData, err error) {
		data.Nums[1] = 7
		return data, nil
	})
	if err != nil {
		t.Fatalf("RefreshToken failed: %v", err)
	}
	if expiredAt.Before(now.Add(time.Second)) {
		t.Fatalf("expiredAt.Before(now.Add(time.Second))")
	}
	if newRefreshToken == refreshToken {
		t.Fatalf("newRefreshToken == refreshToken")
	}

	userClaims, err = f.ParseAccessToken(newAccessToken)
	if err != nil {
		t.Fatalf("ParseAccessToken failed: %v", err)
	}
	if userClaims.Issuer != "mjy" || userClaims.Subject != "test" || userClaims.Data.UserId != 123 || userClaims.Data.Nums[1] != 7 {
		t.Fatalf("claims.Issuer != \"mjy\" || claims.Subject != \"test\" || claims.UserData.UserId != 123 || userClaims.Data.Nums[1] != 7")
	}
}
