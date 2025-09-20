package permission

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
	return new(Permission).Init(scheme, db)
}

type Permission struct {
	crudbase.Crud[entity.Permission, schemas.PermissionItem, schemas.PermissionDetail, schemas.PermissionSave]
}

func (p *Permission) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		p.Crud.Register(api)...)
	return
}

func (p *Permission) Init(scheme string, db *gorm.DB) *Permission {
	p.Crud.Init("Permission", db, func(ctx context.Context, db *gorm.DB) (crud.ICrud[entity.Permission], context.CancelFunc) {
		return service.NewPermissionService(ctx, db)
	}, scheme)
	return p
}
