package services

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/wailsapp/wails/v3/pkg/application"
	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/cmd/wails/views"
	"majinyao.cn/my-app/backend/internal/app"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/internal/service"
	"majinyao.cn/my-app/backend/pkg/db"
)

type UserService struct {
	ctx context.Context
}

func (s *UserService) ServiceStartup(ctx context.Context, options application.ServiceOptions) error {
	s.ctx = ctx
	return nil
}

func (s *UserService) ServiceShutdown() error {
	return nil
}

func (s *UserService) CreateReservedUser(form views.CreateReservedUserForm) (res views.ReservedUserInfo) {
	userService, cancel := service.NewUserService(s.ctx, app.DB)
	defer cancel()

	var entity entity.User
	err := copier.CopyWithOption(&entity, &form, db.DefaultCopierOption)
	if err != nil {
		return
	}

	entity.SysOp(true)
	affected, err := userService.Create(&entity)
	if err != nil {
		return
	}
	if affected < 1 {
		return
	}

	err = copier.CopyWithOption(&res, &entity, db.DefaultCopierOption)
	if err != nil {
		return
	}
	return res
}

func (s *UserService) GetReservedUserInfo() (res views.ReservedUserInfo) {
	userService, cancel := service.NewUserService(s.ctx, app.DB)
	defer cancel()

	_, err := userService.ScanOne(&res, func(tx *gorm.DB) (*gorm.DB, error) {
		return tx.Where("reserved = ?", true), nil
	})
	if err != nil {
		return
	}
	return
}
