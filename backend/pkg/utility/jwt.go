package utility

import (
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
func (*Jwt[TClaims]) ParseToken(token string) (claims TClaims, err error) {
	panic("unimplemented")
}

func NewJwt[TClaims jwt.Claims](claims TClaims, key string) interfaces.IJwt[TClaims] {
	return &Jwt[TClaims]{
		claims: claims,
		key:    key,
	}
}
