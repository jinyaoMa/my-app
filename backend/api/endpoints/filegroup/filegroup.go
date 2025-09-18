package filegroup

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
	return new(FileGroup).Init(scheme, tx)
}

type FileGroup struct {
	crudbase.Crud[entity.FileGroup, schemas.FileGroupItem, schemas.FileGroupDetail, schemas.FileGroupSave]
}

func (g *FileGroup) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		g.Crud.Register(api)...)
	return
}

func (g *FileGroup) Init(scheme string, tx *gorm.DB) *FileGroup {
	g.Crud.Init("FileGroup", tx, db.DefaultCopierOption, func(ctx context.Context, tx *gorm.DB) (crud.ICrudService[entity.FileGroup], context.CancelFunc) {
		return service.NewFileGroupService(ctx, tx)
	}, scheme)
	return g
}
