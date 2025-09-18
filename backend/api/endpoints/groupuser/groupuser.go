package groupuser

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
	return new(GroupUser).Init(scheme, tx)
}

type GroupUser struct {
	crudbase.Crud[entity.GroupUser, schemas.GroupUserItem, schemas.GroupUserDetail, schemas.GroupUserSave]
}

func (u *GroupUser) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		u.Crud.Register(api)...)
	return
}

func (u *GroupUser) Init(scheme string, tx *gorm.DB) *GroupUser {
	u.Crud.Init("GroupUser", tx, db.DefaultCopierOption, func(ctx context.Context, tx *gorm.DB) (crud.ICrudService[entity.GroupUser], context.CancelFunc) {
		return service.NewGroupUserService(ctx, tx)
	}, scheme)
	return u
}
