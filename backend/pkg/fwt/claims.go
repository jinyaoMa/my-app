package fwt

import (
	"errors"
	"time"
)

var (
	ErrClaimsWrongIssuer   = errors.New("fwt claims was issued by others")
	ErrClaimsWrongSubject  = errors.New("fwt claims was for another subject")
	ErrClaimsWrongIssuedAt = errors.New("fwt claims was issued before service epoch")
	ErrClaimsExpired       = errors.New("fwt claims has been expired")
)

type IdentityGetter interface {
	// identity could be user id, device id, or any other unique identifier, or mixed ids like `{userid}_{deviceid}`,
	// normally represents a user with a client, or a user with a device, so that one token for one client/device
	GetIdentity() string
}

type Claims[T IdentityGetter] struct {
	Issuer    string
	Subject   string
	IssuedAt  time.Time // time that claims generated/issued
	ExpiredAt time.Time // time that claims expired
	Data      T
}

func (c *Claims[T]) validate(issuer string, subject string, epoch time.Time) error {
	if c.Issuer != issuer {
		return ErrClaimsWrongIssuer
	}
	if c.Subject != subject {
		return ErrClaimsWrongSubject
	}
	if c.IssuedAt.Before(epoch) {
		return ErrClaimsWrongIssuedAt
	}
	if c.ExpiredAt.Before(time.Now()) {
		return ErrClaimsExpired
	}
	return nil
}
