package auth

import (
	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/api/schemas"
	"majinyao.cn/my-app/backend/pkg/api/endpoint"
	"majinyao.cn/my-app/backend/pkg/api/endpoint/authbase"
)

func New(scheme string, db *gorm.DB) endpoint.Register {
	return new(Auth).Init(scheme, db)
}

type Auth struct {
	authbase.Auth[schemas.UserData]
}

func (a *Auth) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		a.Auth.Register(api)...)
	return
}

func (a *Auth) Init(scheme string, db *gorm.DB) *Auth {
	a.Auth.Init(scheme, db, NewVerifier)
	return a
}
