package rolepermission

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
	return new(RolePermission).Init(res)
}

type RolePermission struct {
	crudbase.Crud[entity.RolePermission, schemas.RolePermissionItem, schemas.RolePermissionSave, schemas.RolePermissionDetail, schemas.RolePermissionSearchItem]
	*apir.Resources
}

func (p *RolePermission) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		p.Crud.RegisterDelete(api),
		p.Crud.RegisterDetail(api),
		p.Crud.RegisterQuery(api, nil),
		p.Crud.RegisterSave(api),
		p.Crud.RegisterSearch(api, func(tx *gorm.DB) (*gorm.DB, error) {
			return tx.Joins("Role").Joins("Permission"), nil
		}),
	)
	return
}

func (p *RolePermission) Init(res *apir.Resources) *RolePermission {
	p.Crud.Init("RolePermission", func(ctx context.Context) (crud.ICrud[entity.RolePermission], context.CancelFunc) {
		return service.NewRolePermissionService(ctx, res.CrudOptions, res.Db)
	}, apir.AuthScheme)
	p.Resources = res
	return p
}
