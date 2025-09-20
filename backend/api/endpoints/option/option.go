package option

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/api/schemas"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/internal/service"
	"majinyao.cn/my-app/backend/pkg/api/endpoint"
	"majinyao.cn/my-app/backend/pkg/api/endpoint/crudbase"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

func New(scheme string, db *gorm.DB) endpoint.Register {
	return new(Option).Init(scheme, db)
}

type Option struct {
	crudbase.Crud[entity.Option, schemas.OptionItem, schemas.OptionDetail, schemas.OptionSave]
}

func (o *Option) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		o.Crud.Register(api)...)
	return
}

func (o *Option) Init(scheme string, db *gorm.DB) *Option {
	o.Crud.Init("Option", db, func(ctx context.Context, db *gorm.DB) (crud.ICrud[entity.Option], context.CancelFunc) {
		return service.NewOptionService(ctx, db)
	}, scheme)
	return o
}
