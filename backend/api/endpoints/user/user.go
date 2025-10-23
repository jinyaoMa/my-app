package user

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"majinyao.cn/my-app/backend/api/apir"
	"majinyao.cn/my-app/backend/api/schemas"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/internal/service"
	"majinyao.cn/my-app/backend/pkg/api/endpoint"
	"majinyao.cn/my-app/backend/pkg/api/endpoint/crudbase"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

func New(res *apir.Resources) endpoint.Register {
	return new(User).Init(res)
}

type User struct {
	crudbase.Crud[entity.User, schemas.UserItem, schemas.UserSave, schemas.UserDetail, schemas.UserSearchItem]
	*apir.Resources
}

func (u *User) Register(api huma.API) (ops []huma.Operation) {
	ops = append(u.Crud.Register(api),
		u.RegisterCurrent(api),
	)
	return
}

func (u *User) Init(res *apir.Resources) *User {
	u.Crud.Init("User", func(ctx context.Context) (crud.ICrud[entity.User], context.CancelFunc) {
		return service.NewUserService(ctx, res.CrudOptions, res.Db, res.Store)
	}, apir.AuthScheme)
	u.Resources = res
	return u
}
