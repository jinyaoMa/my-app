package userpassword

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/api/apir"
	"majinyao.cn/my-app/backend/api/schemas"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/internal/service"
	"majinyao.cn/my-app/backend/pkg/api/endpoint"
	"majinyao.cn/my-app/backend/pkg/api/endpoint/crudbase"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

func New(res *apir.Resources) endpoint.Register {
	return new(UserPassword).Init(res)
}

type UserPassword struct {
	crudbase.Crud[entity.UserPassword, schemas.UserPasswordItem, schemas.UserPasswordSave, schemas.UserPasswordDetail, schemas.UserPasswordSearchItem]
	*apir.Resources
}

func (p *UserPassword) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		p.Crud.RegisterDelete(api),
		p.Crud.RegisterDetail(api),
		p.Crud.RegisterQuery(api, nil),
		p.Crud.RegisterSave(api),
		p.Crud.RegisterSearch(api, func(tx *gorm.DB) (*gorm.DB, error) {
			return tx.Joins("User"), nil
		}),
	)
	return
}

func (p *UserPassword) Init(res *apir.Resources) *UserPassword {
	p.Crud.Init("UserPassword", func(ctx context.Context) (crud.ICrud[entity.UserPassword], context.CancelFunc) {
		return service.NewUserPasswordService(ctx, res.CrudOptions, res.Db)
	}, apir.AuthScheme)
	p.Resources = res
	return p
}
