package userrole

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
	return new(UserRole).Init(scheme, db)
}

type UserRole struct {
	crudbase.Crud[entity.UserRole, schemas.UserRoleItem, schemas.UserRoleDetail, schemas.UserRoleSave]
}

func (r *UserRole) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		r.Crud.Register(api)...)
	return
}

func (r *UserRole) Init(scheme string, db *gorm.DB) *UserRole {
	r.Crud.Init("UserRole", db, func(ctx context.Context, db *gorm.DB) (crud.ICrud[entity.UserRole], context.CancelFunc) {
		return service.NewUserRoleService(ctx, db)
	}, scheme)
	return r
}
