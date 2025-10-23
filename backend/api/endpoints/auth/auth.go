package auth

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"majinyao.cn/my-app/backend/api/apir"
	"majinyao.cn/my-app/backend/internal/service"
	"majinyao.cn/my-app/backend/internal/view"
	"majinyao.cn/my-app/backend/pkg/api/endpoint"
	"majinyao.cn/my-app/backend/pkg/api/endpoint/authbase"
)

func New(res *apir.Resources) endpoint.Register {
	return new(Auth).Init(res)
}

type Auth struct {
	authbase.Auth[view.UserData]
	*apir.Resources
}

func (a *Auth) Register(api huma.API) (ops []huma.Operation) {
	ops = append(a.Auth.Register(api),
		a.RegisterRole(api),
		a.RegisterLoginTemporary(api),
	)
	return
}

func (a *Auth) Init(res *apir.Resources) *Auth {
	a.Auth.Init(apir.AuthScheme, func(ctx context.Context) (authbase.Verifier[view.UserData], context.CancelFunc) {
		return service.NewUserService(ctx, res.CrudOptions, res.Db, res.Store)
	})
	a.Resources = res
	return a
}
