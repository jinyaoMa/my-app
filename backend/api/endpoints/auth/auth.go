package auth

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/api/schemas"
	"majinyao.cn/my-app/backend/internal/service"
	"majinyao.cn/my-app/backend/pkg/api/endpoint"
	"majinyao.cn/my-app/backend/pkg/api/endpoint/authbase"
)

func New(scheme string, tx *gorm.DB) endpoint.Register {
	return new(Auth).Init(scheme, tx)
}

type Auth struct {
	authbase.Auth[schemas.UserData]
}

func (a *Auth) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		a.Auth.Register(api)...)
	return
}

func (a *Auth) Init(scheme string, tx *gorm.DB) *Auth {
	a.Auth.Init(scheme, tx, func(ctx context.Context, tx *gorm.DB) (authbase.Verifier[schemas.UserData], context.CancelFunc) {
		return service.NewUserService(ctx, tx)
	})
	return a
}
