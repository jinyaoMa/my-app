package option

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"majinyao.cn/my-app/backend/api/apir"
	"majinyao.cn/my-app/backend/api/schemas"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/internal/service"
	"majinyao.cn/my-app/backend/pkg/api/endpoint"
	"majinyao.cn/my-app/backend/pkg/api/endpoint/crudbase"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

func New(res *apir.Resources) endpoint.Register {
	return new(Option).Init(res)
}

type Option struct {
	crudbase.Crud[entity.Option, schemas.OptionItem, schemas.OptionSave, schemas.OptionDetail, schemas.OptionSearchItem]
	*apir.Resources
}

func (o *Option) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops, o.Crud.Register(api)...)
	return
}

func (o *Option) Init(res *apir.Resources) *Option {
	o.Crud.Init("Option", func(ctx context.Context) (crud.ICrud[entity.Option], context.CancelFunc) {
		return service.NewOptionService(ctx, res.CrudOptions, res.Db)
	}, apir.AuthScheme)
	o.Resources = res
	return o
}
