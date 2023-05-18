package utility

import (
	"errors"
	"fmt"
	"my-app/backend/pkg/utility/interfaces"

	"github.com/golang-jwt/jwt/v5"
)

type Jwt[TClaims jwt.Claims] struct {
	claims TClaims
	key    string
}

// GetToken implements interfaces.IJwt
func (j *Jwt[TClaims]) GetToken() (token string, err error) {
	ready := jwt.NewWithClaims(jwt.SigningMethodHS512, j.claims)
	token, err = ready.SignedString([]byte(j.key))
	return
}

// ParseToken implements interfaces.IJwt
func (j *Jwt[TClaims]) ParseToken(token string) (claims TClaims, err error) {
	ready, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.key), nil
	})

	var ok bool
	if claims, ok = ready.Claims.(TClaims); ok && ready.Valid {
		return
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		err = errors.New("That's not even a token")
	} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
		// Invalid signature
		err = errors.New("Invalid signature")
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		// Token is either expired or not active yet
		err = errors.New("Token is either expired or not active yet")
	} else {
		err = errors.New(fmt.Sprint("Couldn't handle this token:", err))
	}

	return
}

func NewJwt[TClaims jwt.Claims](claims TClaims, key string) interfaces.IJwt[TClaims] {
	return &Jwt[TClaims]{
		claims: claims,
		key:    key,
	}
}
