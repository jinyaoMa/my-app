package authbase

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/fwt"
)

type Auth[T fwt.IdentityGetter] struct {
	Scheme      string
	Db          *gorm.DB
	NewVerifier func(ctx context.Context, db *gorm.DB) (Verifier[T], context.CancelFunc)
}

func (a *Auth[T]) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		a.RegisterLogin(api),
		a.RegisterRefresh(api),
	)
	return
}

func (a *Auth[T]) Init(scheme string, db *gorm.DB, getVerifier func(ctx context.Context, db *gorm.DB) (Verifier[T], context.CancelFunc)) *Auth[T] {
	a.Scheme = scheme
	a.Db = db
	if getVerifier == nil {
		a.NewVerifier = newVerifier[T]
	} else {
		a.NewVerifier = getVerifier
	}
	return a
}
