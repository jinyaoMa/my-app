package grouprole

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
	return new(GroupRole).Init(res)
}

type GroupRole struct {
	crudbase.Crud[entity.GroupRole, schemas.GroupRoleItem, schemas.GroupRoleSave, schemas.GroupRoleDetail, schemas.GroupRoleSearchItem]
	*apir.Resources
}

func (r *GroupRole) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		r.Crud.RegisterDelete(api),
		r.Crud.RegisterDetail(api),
		r.Crud.RegisterQuery(api, nil),
		r.Crud.RegisterSave(api),
		r.Crud.RegisterSearch(api, func(tx *gorm.DB) (*gorm.DB, error) {
			return tx.Joins("Group").Joins("Role"), nil
		}),
	)
	return
}

func (r *GroupRole) Init(res *apir.Resources) *GroupRole {
	r.Crud.Init("GroupRole", func(ctx context.Context) (crud.ICrud[entity.GroupRole], context.CancelFunc) {
		return service.NewGroupRoleService(ctx, res.CrudOptions, res.Db)
	}, apir.AuthScheme)
	r.Resources = res
	return r
}
