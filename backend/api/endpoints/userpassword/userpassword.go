package userpassword

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/api/schemas"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/internal/service"
	"majinyao.cn/my-app/backend/pkg/api/endpoint"
	"majinyao.cn/my-app/backend/pkg/api/endpoint/crudbase"
	"majinyao.cn/my-app/backend/pkg/db"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

func New(scheme string, tx *gorm.DB) endpoint.Register {
	return new(UserPassword).Init(scheme, tx)
}

type UserPassword struct {
	crudbase.Crud[entity.UserPassword, schemas.UserPasswordItem, schemas.UserPasswordDetail, schemas.UserPasswordSave]
}

func (p *UserPassword) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		p.Crud.Register(api)...)
	return
}

func (p *UserPassword) Init(scheme string, tx *gorm.DB) *UserPassword {
	p.Crud.Init("UserPassword", tx, db.DefaultCopierOption, func(ctx context.Context, tx *gorm.DB) (crud.ICrudService[entity.UserPassword], context.CancelFunc) {
		return service.NewUserPasswordService(ctx, tx)
	}, scheme)
	return p
}
