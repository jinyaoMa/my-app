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
	"majinyao.cn/my-app/backend/pkg/db"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

func New(scheme string, tx *gorm.DB) endpoint.Register {
	return new(File).Init(scheme, tx)
}

type File struct {
	crudbase.Crud[entity.File, schemas.FileItem, schemas.FileDetail, schemas.FileSave]
}

func (f *File) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		f.Crud.Register(api)...)
	return
}

func (f *File) Init(scheme string, tx *gorm.DB) *File {
	f.Crud.Init("File", tx, db.DefaultCopierOption, func(ctx context.Context, tx *gorm.DB) (crud.ICrudService[entity.File], context.CancelFunc) {
		return service.NewFileService(ctx, tx)
	}, scheme)
	return f
}
