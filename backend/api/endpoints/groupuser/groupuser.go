package groupuser

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
	return new(GroupUser).Init(res)
}

type GroupUser struct {
	crudbase.Crud[entity.GroupUser, schemas.GroupUserItem, schemas.GroupUserSave, schemas.GroupUserDetail, schemas.GroupUserSearchItem]
	*apir.Resources
}

func (u *GroupUser) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		u.Crud.RegisterDelete(api),
		u.Crud.RegisterDetail(api),
		u.Crud.RegisterQuery(api, nil),
		u.Crud.RegisterSave(api),
		u.Crud.RegisterSearch(api, func(tx *gorm.DB) (*gorm.DB, error) {
			return tx.Joins("Group").Joins("User"), nil
		}),
	)
	return
}

func (u *GroupUser) Init(res *apir.Resources) *GroupUser {
	u.Crud.Init("GroupUser", func(ctx context.Context) (crud.ICrud[entity.GroupUser], context.CancelFunc) {
		return service.NewGroupUserService(ctx, res.CrudOptions, res.Db)
	}, apir.AuthScheme)
	u.Resources = res
	return u
}
