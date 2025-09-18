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
	"majinyao.cn/my-app/backend/pkg/db"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

func New(scheme string, tx *gorm.DB) endpoint.Register {
	return new(UserRole).Init(scheme, tx)
}

type UserRole struct {
	crudbase.Crud[entity.UserRole, schemas.UserRoleItem, schemas.UserRoleDetail, schemas.UserRoleSave]
}

func (r *UserRole) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		r.Crud.Register(api)...)
	return
}

func (r *UserRole) Init(scheme string, tx *gorm.DB) *UserRole {
	r.Crud.Init("UserRole", tx, db.DefaultCopierOption, func(ctx context.Context, tx *gorm.DB) (crud.ICrudService[entity.UserRole], context.CancelFunc) {
		return service.NewUserRoleService(ctx, tx)
	}, scheme)
	return r
}
