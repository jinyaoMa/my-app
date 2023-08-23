package jwt

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type JWT[TClaims jwt.Claims] struct {
	claims TClaims
	key    string
}

// GetToken implements IJWT
func (j *JWT[TClaims]) GetToken() (token string, err error) {
	ready := jwt.NewWithClaims(jwt.SigningMethodHS512, j.claims)
	token, err = ready.SignedString([]byte(j.key))
	return
}

// ParseToken implements IJWT
func (j *JWT[TClaims]) ParseToken(token string) (claims TClaims, err error) {
	ready, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.key), nil
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

func NewJWT[TClaims jwt.Claims](claims TClaims, key string) *JWT[TClaims] {
	return &JWT[TClaims]{
		claims: claims,
		key:    key,
	}
}

func NewIJWT[TClaims jwt.Claims](claims TClaims, key string) IJWT[TClaims] {
	return NewJWT[TClaims](claims, key)
}
