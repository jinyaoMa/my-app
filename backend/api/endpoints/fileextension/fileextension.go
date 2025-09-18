package fileextension

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
	return new(FileExtension).Init(scheme, tx)
}

type FileExtension struct {
	crudbase.Crud[entity.FileExtension, schemas.FileExtensionItem, schemas.FileExtensionDetail, schemas.FileExtensionSave]
}

func (e *FileExtension) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		e.Crud.Register(api)...)
	return
}

func (e *FileExtension) Init(scheme string, tx *gorm.DB) *FileExtension {
	e.Crud.Init("FileExtension", tx, db.DefaultCopierOption, func(ctx context.Context, tx *gorm.DB) (crud.ICrudService[entity.FileExtension], context.CancelFunc) {
		return service.NewFileExtensionService(ctx, tx)
	}, scheme)
	return e
}
