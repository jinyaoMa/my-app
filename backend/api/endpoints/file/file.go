package file

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
	return new(File).Init(scheme, db)
}

type File struct {
	crudbase.Crud[entity.File, schemas.FileItem, schemas.FileDetail, schemas.FileSave]
}

func (f *File) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		f.Crud.Register(api)...)
	return
}

func (f *File) Init(scheme string, db *gorm.DB) *File {
	f.Crud.Init("File", db, func(ctx context.Context, db *gorm.DB) (crud.ICrud[entity.File], context.CancelFunc) {
		return service.NewFileService(ctx, db)
	}, scheme)
	return f
}
