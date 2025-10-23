package services

import (
	"context"

	"github.com/wailsapp/wails/v3/pkg/application"
	"majinyao.cn/my-app/backend/cmd/wails/views"
	"majinyao.cn/my-app/backend/internal/app"
	"majinyao.cn/my-app/backend/internal/service"
	"majinyao.cn/my-app/backend/internal/view"
)

type UserService struct {
	ctx         context.Context
	userService service.IUserService
}

func (s *UserService) ServiceStartup(ctx context.Context, options application.ServiceOptions) error {
	s.ctx = ctx
	s.userService, _ = service.NewUserService(s.ctx, app.CFG.Get().Crud, app.DB, app.STORE)
	return nil
}

func (s *UserService) ServiceShutdown() error {
	return nil
}

func (s *UserService) CreateReservedUser(form views.CreateReservedUserForm) (info view.ReservedUserInfo) {
	info, _ = s.userService.CreateReservedUser(view.CreateReservedUserForm{
		Name:     form.Name,
		Account:  form.Account,
		Password: form.Password,
	})
	return
}

func (s *UserService) GetReservedUserInfo() (info view.ReservedUserInfo) {
	info, _ = s.userService.GetReservedUserInfo()
	return
}
