package filecategory

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
	return new(FileCategory).Init(res)
}

type FileCategory struct {
	crudbase.Crud[entity.FileCategory, schemas.FileCategoryItem, schemas.FileCategorySave, schemas.FileCategoryDetail, schemas.FileCategorySearchItem]
	*apir.Resources
}

func (c *FileCategory) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops, c.Crud.Register(api)...)
	return
}

func (c *FileCategory) Init(res *apir.Resources) *FileCategory {
	c.Crud.Init("FileCategory", func(ctx context.Context) (crud.ICrud[entity.FileCategory], context.CancelFunc) {
		return service.NewFileCategoryService(ctx, res.CrudOptions, res.Db)
	}, apir.AuthScheme)
	c.Resources = res
	return c
}
