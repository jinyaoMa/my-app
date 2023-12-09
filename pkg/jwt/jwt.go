package jwt

import (
	"errors"
	"fmt"
	"my-app/pkg/crypto"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT[TClaims IClaims] struct {
	options *JWTOptions
	crypto  crypto.ICrypto // refresh token encryption
}

// GenerateToken implements IJWT.
func (j *JWT[TClaims]) GenerateToken(claims TClaims) (accessToken string, refreshToken string, err error) {
	id, err := claims.GetID()
	if err != nil {
		return "", "", err
	}
	expirationTime, err := claims.GetExpirationTime()
	if err != nil {
		return "", "", err
	}
	ready := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	accessToken, err = ready.SignedString(j.options.GetAccessTokenKey())
	if err != nil {
		return "", "", err
	}
	refreshToken_ := NewRefreshToken(id, expirationTime.Add(j.options.GetRefreshTokenExtension()))
	refreshToken, err = j.crypto.Encrypt(string(refreshToken_))
	if err != nil {
		return "", "", err
	}
	return
}

// ParseToken implements IJWT.
func (j *JWT[TClaims]) ParseToken(token string) (claims TClaims, err error) {
	ready, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return j.options.GetAccessTokenKey(), nil
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

// RefreshToken implements IJWT.
func (j *JWT[TClaims]) RefreshToken(accessToken string, refreshToken string) (newAccessToken string, newRefreshToken string, err error) {
	claims, err := j.ParseToken(accessToken)
	if err != nil {
		return "", "", err
	}
	id, err := claims.GetID()
	if err != nil {
		return "", "", err
	}
	refreshToken, err = j.crypto.Decrypt(refreshToken)
	if err != nil {
		return "", "", err
	}
	id_, expirationTime, err := RefreshToken(refreshToken).Split()
	if err != nil {
		return "", "", err
	}
	if id != id_ {
		e := fmt.Sprintf("refresh token %s is invalid for access token %s", refreshToken, accessToken)
		return "", "", errors.New(e)
	}
	if expirationTime.Before(time.Now()) {
		e := fmt.Sprintf("refresh token %s has been expired at %s", refreshToken, expirationTime.String())
		return "", "", errors.New(e)
	}
	err = claims.RefreshExpirationTime()
	if err != nil {
		return "", "", err
	}
	return j.GenerateToken(claims)
}

func NewJWT[TClaims IClaims](crypto crypto.ICrypto, options *JWTOptions) (jwt *JWT[TClaims], iJWT IJWT[TClaims]) {
	jwt = &JWT[TClaims]{
		options: options,
		crypto:  crypto,
	}
	return jwt, jwt
}
