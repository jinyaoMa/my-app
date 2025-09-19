package authbase

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/fwt"
)

type Verifier[T fwt.IdentityGetter] interface {
	// verify and update userdata
	VerifyUserData(userdata T, visitorId string) (newUserdata T, err error)
	// verify login and return userdata
	VerifyLogin(input *LoginInput) (userdata T, err error)
}

type verifier[T fwt.IdentityGetter] struct{}

// VerifyLogin implements Verifier.
func (v *verifier[T]) VerifyLogin(input *LoginInput) (userdata T, err error) {
	return userdata, errors.New("unimplemented")
}

// VerifyUserData implements Verifier.
func (v *verifier[T]) VerifyUserData(userdata T, visitorId string) (newUserdata T, err error) {
	return newUserdata, errors.New("unimplemented")
}

func newVerifier[T fwt.IdentityGetter](ctx context.Context, tx *gorm.DB) (Verifier[T], context.CancelFunc) {
	return &verifier[T]{}, func() {}
}
