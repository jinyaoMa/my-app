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
	"majinyao.cn/my-app/backend/pkg/db"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

func New(scheme string, tx *gorm.DB) endpoint.Register {
	return new(Group).Init(scheme, tx)
}

type Group struct {
	crudbase.Crud[entity.Group, schemas.GroupItem, schemas.GroupDetail, schemas.GroupSave]
}

func (g *Group) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		g.Crud.Register(api)...)
	return
}

func (g *Group) Init(scheme string, tx *gorm.DB) *Group {
	g.Crud.Init("Group", tx, db.DefaultCopierOption, func(ctx context.Context, tx *gorm.DB) (crud.ICrudService[entity.Group], context.CancelFunc) {
		return service.NewGroupService(ctx, tx)
	}, scheme)
	return g
}
