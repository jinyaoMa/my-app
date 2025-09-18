package role

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
	return new(Role).Init(scheme, tx)
}

type Role struct {
	crudbase.Crud[entity.Role, schemas.RoleItem, schemas.RoleDetail, schemas.RoleSave]
}

func (r *Role) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		r.Crud.Register(api)...)
	return
}

func (r *Role) Init(scheme string, tx *gorm.DB) *Role {
	r.Crud.Init("Role", tx, db.DefaultCopierOption, func(ctx context.Context, tx *gorm.DB) (crud.ICrudService[entity.Role], context.CancelFunc) {
		return service.NewRoleService(ctx, tx)
	}, scheme)
	return r
}
