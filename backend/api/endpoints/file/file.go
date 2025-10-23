package file

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
	return new(File).Init(res)
}

type File struct {
	crudbase.Crud[entity.File, schemas.FileItem, schemas.FileSave, schemas.FileDetail, schemas.FileSearchItem]
	*apir.Resources
}

func (f *File) Register(api huma.API) (ops []huma.Operation) {
	ops = append(f.Crud.Register(api),
		f.RegisterData(api),
	)
	return
}

func (f *File) Init(res *apir.Resources) *File {
	f.Crud.Init("File", func(ctx context.Context) (crud.ICrud[entity.File], context.CancelFunc) {
		return service.NewFileService(ctx, res.CrudOptions, res.Db, res.Store)
	}, apir.AuthScheme)
	f.Resources = res
	return f
}
