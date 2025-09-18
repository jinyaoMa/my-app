package rolepermission

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
	return new(RolePermission).Init(scheme, tx)
}

type RolePermission struct {
	crudbase.Crud[entity.RolePermission, schemas.RolePermissionItem, schemas.RolePermissionDetail, schemas.RolePermissionSave]
}

func (p *RolePermission) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		p.Crud.Register(api)...)
	return
}

func (p *RolePermission) Init(scheme string, tx *gorm.DB) *RolePermission {
	p.Crud.Init("RolePermission", tx, db.DefaultCopierOption, func(ctx context.Context, tx *gorm.DB) (crud.ICrudService[entity.RolePermission], context.CancelFunc) {
		return service.NewRolePermissionService(ctx, tx)
	}, scheme)
	return p
}
