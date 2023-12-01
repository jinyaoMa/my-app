package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type IClaims interface {
	jwt.Claims
	GetID() (string, error)
	RefreshExpirationTime() error
}

type IJWT[TClaims IClaims] interface {
	RefreshToken(accessToken string, refreshToken string) (newAccessToken string, newRefreshToken string, err error)
	GenerateToken(claims TClaims) (accessToken string, refreshToken string, err error)
	ParseToken(token string) (claims TClaims, err error)
}
