package authbase

import (
	"majinyao.cn/my-app/backend/pkg/fwt"
)

type Verifier[T fwt.IdentityGetter] interface {
	// verify and update userdata
	VerifyUserData(userdata T, visitorId string) (newUserdata T, err error)
	// verify login and return userdata
	VerifyLogin(input *LoginInput) (userdata T, err error)
}
