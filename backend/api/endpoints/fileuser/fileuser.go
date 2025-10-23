package fileuser

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
	return new(FileUser).Init(res)
}

type FileUser struct {
	crudbase.Crud[entity.FileUser, schemas.FileUserItem, schemas.FileUserSave, schemas.FileUserDetail, schemas.FileUserSearchItem]
	*apir.Resources
}

func (u *FileUser) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		u.Crud.RegisterDelete(api),
		u.Crud.RegisterDetail(api),
		u.Crud.RegisterQuery(api, nil),
		u.Crud.RegisterSave(api),
		u.Crud.RegisterSearch(api, func(tx *gorm.DB) (*gorm.DB, error) {
			return tx.Joins("File").Joins("User"), nil
		}),
	)
	return
}

func (u *FileUser) Init(res *apir.Resources) *FileUser {
	u.Crud.Init("FileUser", func(ctx context.Context) (crud.ICrud[entity.FileUser], context.CancelFunc) {
		return service.NewFileUserService(ctx, res.CrudOptions, res.Db)
	}, apir.AuthScheme)
	u.Resources = res
	return u
}
