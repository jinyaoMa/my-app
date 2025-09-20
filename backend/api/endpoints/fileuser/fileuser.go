package fileuser

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
	return new(FileUser).Init(scheme, db)
}

type FileUser struct {
	crudbase.Crud[entity.FileUser, schemas.FileUserItem, schemas.FileUserDetail, schemas.FileUserSave]
}

func (u *FileUser) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		u.Crud.Register(api)...)
	return
}

func (u *FileUser) Init(scheme string, db *gorm.DB) *FileUser {
	u.Crud.Init("FileUser", db, func(ctx context.Context, db *gorm.DB) (crud.ICrud[entity.FileUser], context.CancelFunc) {
		return service.NewFileUserService(ctx, db)
	}, scheme)
	return u
}
