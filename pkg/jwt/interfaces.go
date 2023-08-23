package jwt

import "github.com/golang-jwt/jwt/v5"

type IJWT[TClaims jwt.Claims] interface {
	GetToken() (token string, err error)
	ParseToken(token string) (claims TClaims, err error)
}
