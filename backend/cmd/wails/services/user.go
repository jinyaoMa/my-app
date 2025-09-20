package services

import (
	"context"
	"errors"

	"github.com/jinzhu/copier"
	"github.com/wailsapp/wails/v3/pkg/application"
	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/cmd/wails/views"
	"majinyao.cn/my-app/backend/internal/app"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/internal/service"
	"majinyao.cn/my-app/backend/pkg/db/crud"
	"majinyao.cn/my-app/backend/pkg/db/dbcontext"
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
	db, cancel := dbcontext.SectionUnderContextWithCancel(s.ctx, app.DB)
	defer cancel()

	db.Transaction(func(tx *gorm.DB) error {
		roleService := service.UseRoleService(tx)
		userService := service.UseUserService(tx)
		userRoleService := service.UseUserRoleService(tx)

		reservedRole, notFound, err := roleService.FindOne(func(tx *gorm.DB) (*gorm.DB, error) {
			return tx.Where("reserved = ?", true), nil
		})
		if err != nil {
			return err
		}
		if notFound {
			return errors.New("reserved role not found")
		}

		var user entity.User
		err = copier.CopyWithOption(&user, &form, crud.DefaultCopierOption)
		if err != nil {
			return err
		}

		user.SysOp(true)
		affected, err := userService.Create(&user)
		if err != nil {
			return err
		}
		if affected < 1 {
			return errors.New("create reserved user failed")
		}

		affected, err = userRoleService.Create(&entity.UserRole{
			UserId: user.Id,
			RoleId: reservedRole.Id,
		})
		if err != nil {
			return err
		}
		if affected < 1 {
			return errors.New("create reserved user role failed")
		}

		err = copier.CopyWithOption(&res, &user, crud.DefaultCopierOption)
		if err != nil {
			return err
		}
		return nil
	})
	return
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
