package filegroup

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
	return new(FileGroup).Init(res)
}

type FileGroup struct {
	crudbase.Crud[entity.FileGroup, schemas.FileGroupItem, schemas.FileGroupSave, schemas.FileGroupDetail, schemas.FileGroupSearchItem]
	*apir.Resources
}

func (g *FileGroup) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		g.Crud.RegisterDelete(api),
		g.Crud.RegisterDetail(api),
		g.Crud.RegisterQuery(api, nil),
		g.Crud.RegisterSave(api),
		g.Crud.RegisterSearch(api, func(tx *gorm.DB) (*gorm.DB, error) {
			return tx.Joins("File").Joins("Group"), nil
		}),
	)
	return
}

func (g *FileGroup) Init(res *apir.Resources) *FileGroup {
	g.Crud.Init("FileGroup", func(ctx context.Context) (crud.ICrud[entity.FileGroup], context.CancelFunc) {
		return service.NewFileGroupService(ctx, res.CrudOptions, res.Db)
	}, apir.AuthScheme)
	g.Resources = res
	return g
}
