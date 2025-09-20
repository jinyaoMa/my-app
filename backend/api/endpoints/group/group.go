package group

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
	return new(Group).Init(scheme, db)
}

type Group struct {
	crudbase.Crud[entity.Group, schemas.GroupItem, schemas.GroupDetail, schemas.GroupSave]
}

func (g *Group) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		g.Crud.Register(api)...)
	return
}

func (g *Group) Init(scheme string, db *gorm.DB) *Group {
	g.Crud.Init("Group", db, func(ctx context.Context, db *gorm.DB) (crud.ICrud[entity.Group], context.CancelFunc) {
		return service.NewGroupService(ctx, db)
	}, scheme)
	return g
}
