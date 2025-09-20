package filecategory

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
	return new(FileCategory).Init(scheme, db)
}

type FileCategory struct {
	crudbase.Crud[entity.FileCategory, schemas.FileCategoryItem, schemas.FileCategoryDetail, schemas.FileCategorySave]
}

func (c *FileCategory) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		c.Crud.Register(api)...)
	return
}

func (c *FileCategory) Init(scheme string, db *gorm.DB) *FileCategory {
	c.Crud.Init("FileCategory", db, func(ctx context.Context, db *gorm.DB) (crud.ICrud[entity.FileCategory], context.CancelFunc) {
		return service.NewFileCategoryService(ctx, db)
	}, scheme)
	return c
}
