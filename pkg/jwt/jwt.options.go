package jwt

import (
	"my-app/pkg/base"
	"time"
)

type JWTOptions struct {
	base.Options
	AccessTokenKey        string // key for access token
	RefreshTokenSalt      string // salt for refresh token
	RefreshTokenExtension int64  // extended time since expired to allow refresh, unit: second
	NumberOfUsers         int    // limit number of users to generate token at the same time
}

func (jwtOptions *JWTOptions) GetAccessTokenKey() []byte {
	return []byte(jwtOptions.AccessTokenKey)
}

func (jwtOptions *JWTOptions) GetRefreshTokenExtension() time.Duration {
	return time.Duration(jwtOptions.RefreshTokenExtension) * time.Second
}

func DefaultJWTOptions() *JWTOptions {
	return &JWTOptions{
		AccessTokenKey:        "jinyaoMa",
		RefreshTokenSalt:      "1234567890",
		RefreshTokenExtension: 10800, // can refresh access token 3 hours after it expired
		NumberOfUsers:         100,
	}
}

func NewJWTOptions(dst *JWTOptions) (*JWTOptions, error) {
	return base.MergeOptions(DefaultJWTOptions(), dst)
}
