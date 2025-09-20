package service

import (
	"context"
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

type IRoleService interface {
	crud.ICrud[entity.Role]
}

func NewRoleService(ctx context.Context, db *gorm.DB) (IRoleService, context.CancelFunc) {
	s, cancel := new(RoleService).InitWithCancelUnderContext(ctx, db)
	return s, cancel
}

func UseRoleService(db *gorm.DB) IRoleService {
	return new(RoleService).Init(db)
}

type RoleService struct {
	crud.Crud[entity.Role]
}

func (s *RoleService) Init(db *gorm.DB) *RoleService {
	s.Crud.Init(db)
	return s
}

func (s *RoleService) InitWithCancelUnderContext(ctx context.Context, db *gorm.DB) (*RoleService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithCancelUnderContext(ctx, db)
	return s, cancel
}

func (s *RoleService) InitWithTimeoutUnderContext(ctx context.Context, db *gorm.DB, timeout time.Duration) (*RoleService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithTimeoutUnderContext(ctx, db, timeout)
	return s, cancel
}
