package service

import (
	"context"
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

type IUserRoleService interface {
	crud.ICrud[entity.UserRole]
}

func NewUserRoleService(ctx context.Context, db *gorm.DB) (IUserRoleService, context.CancelFunc) {
	s, cancel := new(UserRoleService).InitWithCancelUnderContext(ctx, db)
	return s, cancel
}

func UseUserRoleService(db *gorm.DB) IUserRoleService {
	return new(UserRoleService).Init(db)
}

type UserRoleService struct {
	crud.Crud[entity.UserRole]
}

func (s *UserRoleService) Init(db *gorm.DB) *UserRoleService {
	s.Crud.Init(db)
	return s
}

func (s *UserRoleService) InitWithCancelUnderContext(ctx context.Context, db *gorm.DB) (*UserRoleService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithCancelUnderContext(ctx, db)
	return s, cancel
}

func (s *UserRoleService) InitWithTimeoutUnderContext(ctx context.Context, db *gorm.DB, timeout time.Duration) (*UserRoleService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithTimeoutUnderContext(ctx, db, timeout)
	return s, cancel
}
