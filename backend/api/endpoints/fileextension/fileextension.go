package fileextension

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
	return new(FileExtension).Init(res)
}

type FileExtension struct {
	crudbase.Crud[entity.FileExtension, schemas.FileExtensionItem, schemas.FileExtensionSave, schemas.FileExtensionDetail, schemas.FileExtensionSearchItem]
	*apir.Resources
}

func (e *FileExtension) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops, e.Crud.Register(api)...)
	return
}

func (e *FileExtension) Init(res *apir.Resources) *FileExtension {
	e.Crud.Init("FileExtension", func(ctx context.Context) (crud.ICrud[entity.FileExtension], context.CancelFunc) {
		return service.NewFileExtensionService(ctx, res.CrudOptions, res.Db)
	}, apir.AuthScheme)
	e.Resources = res
	return e
}
