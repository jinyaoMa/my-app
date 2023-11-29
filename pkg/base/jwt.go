package base

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwtToken[TClaims jwt.Claims](key string, claims TClaims) (token string, err error) {
	ready := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err = ready.SignedString([]byte(key))
	return
}

func ParseJwtToken[TClaims jwt.Claims](key string, token string) (claims TClaims, err error) {
	ready, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	var ok bool
	if claims, ok = ready.Claims.(TClaims); ok && ready.Valid {
		return
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		err = errors.New("that's not even a token")
	} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
		// Invalid signature
		err = errors.New("invalid signature")
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		// Token is either expired or not active yet
		err = errors.New("token is either expired or not active yet")
	} else {
		err = errors.New(fmt.Sprint("couldn't handle this token:", err))
	}

	return
}
