package service

import (
	"context"
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

type IUserRoleService interface {
	crud.ICrudService[entity.UserRole]
}

func NewUserRoleService(ctx context.Context, tx *gorm.DB) (IUserRoleService, context.CancelFunc) {
	s, cancel := new(UserRoleService).InitWithCancelUnderContext(ctx, tx)
	return s, cancel
}

func UseUserRoleService(tx *gorm.DB) IUserRoleService {
	return new(UserRoleService).Init(tx)
}

type UserRoleService struct {
	crud.Crud[entity.UserRole]
}

func (s *UserRoleService) Init(tx *gorm.DB) *UserRoleService {
	s.Crud.Init(tx)
	return s
}

func (s *UserRoleService) InitWithCancelUnderContext(ctx context.Context, tx *gorm.DB) (*UserRoleService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithCancelUnderContext(ctx, tx)
	return s, cancel
}

func (s *UserRoleService) InitWithTimeoutUnderContext(ctx context.Context, tx *gorm.DB, timeout time.Duration) (*UserRoleService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithTimeoutUnderContext(ctx, tx, timeout)
	return s, cancel
}
