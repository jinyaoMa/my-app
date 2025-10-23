package permission

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
	return new(Permission).Init(res)
}

type Permission struct {
	crudbase.Crud[entity.Permission, schemas.PermissionItem, schemas.PermissionSave, schemas.PermissionDetail, schemas.PermissionSearchItem]
	*apir.Resources
}

func (p *Permission) Register(api huma.API) (ops []huma.Operation) {
	ops = append(p.Crud.Register(api),
		p.RegisterHas(api),
	)
	return
}

func (p *Permission) Init(res *apir.Resources) *Permission {
	p.Crud.Init("Permission", func(ctx context.Context) (crud.ICrud[entity.Permission], context.CancelFunc) {
		return service.NewPermissionService(ctx, res.CrudOptions, res.Db)
	}, apir.AuthScheme)
	p.Resources = res
	return p
}
