package user

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
	return new(User).Init(scheme, tx)
}

type User struct {
	crudbase.Crud[entity.User, schemas.UserItem, schemas.UserDetail, schemas.UserSave]
}

func (u *User) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		u.Crud.Register(api)...)
	return
}

func (u *User) Init(scheme string, tx *gorm.DB) *User {
	u.Crud.Init("User", tx, db.DefaultCopierOption, func(ctx context.Context, tx *gorm.DB) (crud.ICrudService[entity.User], context.CancelFunc) {
		return service.NewUserService(ctx, tx)
	}, scheme)
	return u
}
