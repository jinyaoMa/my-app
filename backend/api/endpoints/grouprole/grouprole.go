package grouprole

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
	return new(GroupRole).Init(scheme, tx)
}

type GroupRole struct {
	crudbase.Crud[entity.GroupRole, schemas.GroupRoleItem, schemas.GroupRoleDetail, schemas.GroupRoleSave]
}

func (r *GroupRole) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		r.Crud.Register(api)...)
	return
}

func (r *GroupRole) Init(scheme string, tx *gorm.DB) *GroupRole {
	r.Crud.Init("GroupRole", tx, db.DefaultCopierOption, func(ctx context.Context, tx *gorm.DB) (crud.ICrudService[entity.GroupRole], context.CancelFunc) {
		return service.NewGroupRoleService(ctx, tx)
	}, scheme)
	return r
}
