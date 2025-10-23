package authbase

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"majinyao.cn/my-app/backend/pkg/fwt"
)

type Auth[T fwt.IdentityGetter] struct {
	Scheme      string
	NewVerifier func(ctx context.Context) (Verifier[T], context.CancelFunc)
}

func (a *Auth[T]) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		a.RegisterLogin(api),
		a.RegisterRefresh(api),
	)
	return
}

func (a *Auth[T]) Init(scheme string, getVerifier func(ctx context.Context) (Verifier[T], context.CancelFunc)) *Auth[T] {
	a.Scheme = scheme
	a.NewVerifier = getVerifier
	return a
}
