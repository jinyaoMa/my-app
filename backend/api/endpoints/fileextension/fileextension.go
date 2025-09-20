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
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

func New(scheme string, db *gorm.DB) endpoint.Register {
	return new(FileExtension).Init(scheme, db)
}

type FileExtension struct {
	crudbase.Crud[entity.FileExtension, schemas.FileExtensionItem, schemas.FileExtensionDetail, schemas.FileExtensionSave]
}

func (e *FileExtension) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		e.Crud.Register(api)...)
	return
}

func (e *FileExtension) Init(scheme string, db *gorm.DB) *FileExtension {
	e.Crud.Init("FileExtension", db, func(ctx context.Context, db *gorm.DB) (crud.ICrud[entity.FileExtension], context.CancelFunc) {
		return service.NewFileExtensionService(ctx, db)
	}, scheme)
	return e
}
