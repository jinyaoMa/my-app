package role

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
	return new(Role).Init(res)
}

type Role struct {
	crudbase.Crud[entity.Role, schemas.RoleItem, schemas.RoleSave, schemas.RoleDetail, schemas.RoleSearchItem]
	*apir.Resources
}

func (r *Role) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops, r.Crud.Register(api)...)
	return
}

func (r *Role) Init(res *apir.Resources) *Role {
	r.Crud.Init("Role", func(ctx context.Context) (crud.ICrud[entity.Role], context.CancelFunc) {
		return service.NewRoleService(ctx, res.CrudOptions, res.Db)
	}, apir.AuthScheme)
	r.Resources = res
	return r
}
