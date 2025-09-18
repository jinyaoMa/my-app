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
	"majinyao.cn/my-app/backend/pkg/db"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

func New(scheme string, tx *gorm.DB) endpoint.Register {
	return new(FileCategory).Init(scheme, tx)
}

type FileCategory struct {
	crudbase.Crud[entity.FileCategory, schemas.FileCategoryItem, schemas.FileCategoryDetail, schemas.FileCategorySave]
}

func (c *FileCategory) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		c.Crud.Register(api)...)
	return
}

func (c *FileCategory) Init(scheme string, tx *gorm.DB) *FileCategory {
	c.Crud.Init("FileCategory", tx, db.DefaultCopierOption, func(ctx context.Context, tx *gorm.DB) (crud.ICrudService[entity.FileCategory], context.CancelFunc) {
		return service.NewFileCategoryService(ctx, tx)
	}, scheme)
	return c
}
