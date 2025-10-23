package group

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
	return new(Group).Init(res)
}

type Group struct {
	crudbase.Crud[entity.Group, schemas.GroupItem, schemas.GroupSave, schemas.GroupDetail, schemas.GroupSearchItem]
	*apir.Resources
}

func (g *Group) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops, g.Crud.Register(api)...)
	return
}

func (g *Group) Init(res *apir.Resources) *Group {
	g.Crud.Init("Group", func(ctx context.Context) (crud.ICrud[entity.Group], context.CancelFunc) {
		return service.NewGroupService(ctx, res.CrudOptions, res.Db)
	}, apir.AuthScheme)
	g.Resources = res
	return g
}
