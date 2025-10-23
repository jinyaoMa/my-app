package userrole

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
	return new(UserRole).Init(res)
}

type UserRole struct {
	crudbase.Crud[entity.UserRole, schemas.UserRoleItem, schemas.UserRoleSave, schemas.UserRoleDetail, schemas.UserRoleSearchItem]
	*apir.Resources
}

func (r *UserRole) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		r.Crud.RegisterDelete(api),
		r.Crud.RegisterDetail(api),
		r.Crud.RegisterQuery(api, nil),
		r.Crud.RegisterSave(api),
		r.Crud.RegisterSearch(api, func(tx *gorm.DB) (*gorm.DB, error) {
			return tx.Joins("User").Joins("Role"), nil
		}),
	)
	return
}

func (r *UserRole) Init(res *apir.Resources) *UserRole {
	r.Crud.Init("UserRole", func(ctx context.Context) (crud.ICrud[entity.UserRole], context.CancelFunc) {
		return service.NewUserRoleService(ctx, res.CrudOptions, res.Db)
	}, apir.AuthScheme)
	r.Resources = res
	return r
}
