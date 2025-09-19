package auth

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/api/schemas"
	"majinyao.cn/my-app/backend/internal/service"
	"majinyao.cn/my-app/backend/pkg/api/endpoint/authbase"
	"majinyao.cn/my-app/backend/pkg/api/middlewares/memcache"
	"majinyao.cn/my-app/backend/pkg/flag"
)

func NewVerifier(ctx context.Context, tx *gorm.DB) (authbase.Verifier[schemas.UserData], context.CancelFunc) {
	v := new(Verifier)
	v.ctx = ctx
	userService, cancel := service.NewUserService(ctx, tx)
	v.IUserService = userService
	return v, func() {
		cancel()
	}
}

type Verifier struct {
	ctx context.Context
	service.IUserService
}

func (v *Verifier) VerifyUserData(userdata schemas.UserData, visitorId string) (newUserdata schemas.UserData, err error) {
	if userdata.VisitorId != visitorId {
		return userdata, fmt.Errorf("visitor id not matched")
	}
	return userdata, nil
}

func (v *Verifier) VerifyLogin(input *authbase.LoginInput) (userdata schemas.UserData, err error) {
	user, notFound, err := v.GetByAccountPassword(input.Body.Account, input.Body.Password, "Roles.Permissions")
	if err != nil {
		return
	}
	if notFound {
		err = errors.New("account or password wrong")
		return
	}

	var flag flag.IFlag
	for _, role := range user.Roles {
		for _, perm := range role.Permissions {
			fmt.Printf("\nROLE_PERM: %b\n\n", perm.GetFlag())
			flag = perm.GetFlag().Or(flag)
			fmt.Printf("\nROLE_PERM_FLAG: %b\n\n", flag)
		}
	}

	cache := memcache.GetFromContext(v.ctx)
	cache.Set(fmt.Sprintf("auth_perm_%d", user.Id.Int64()), flag)
	return schemas.UserData{
		VisitorId: input.VisitorId,
		UserId:    user.Id.Int64(),
	}, nil
}
