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
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

func New(scheme string, db *gorm.DB) endpoint.Register {
	return new(RolePermission).Init(scheme, db)
}

type RolePermission struct {
	crudbase.Crud[entity.RolePermission, schemas.RolePermissionItem, schemas.RolePermissionDetail, schemas.RolePermissionSave]
}

func (p *RolePermission) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		p.Crud.Register(api)...)
	return
}

func (p *RolePermission) Init(scheme string, db *gorm.DB) *RolePermission {
	p.Crud.Init("RolePermission", db, func(ctx context.Context, db *gorm.DB) (crud.ICrud[entity.RolePermission], context.CancelFunc) {
		return service.NewRolePermissionService(ctx, db)
	}, scheme)
	return p
}
