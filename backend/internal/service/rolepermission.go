package service

import (
	"context"
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

type IRolePermissionService interface {
	crud.ICrud[entity.RolePermission]
}

func NewRolePermissionService(ctx context.Context, db *gorm.DB) (IRolePermissionService, context.CancelFunc) {
	s, cancel := new(RolePermissionService).InitWithCancelUnderContext(ctx, db)
	return s, cancel
}

func UseRolePermissionService(db *gorm.DB) IRolePermissionService {
	return new(RolePermissionService).Init(db)
}

type RolePermissionService struct {
	crud.Crud[entity.RolePermission]
}

func (s *RolePermissionService) Init(db *gorm.DB) *RolePermissionService {
	s.Crud.Init(db)
	return s
}

func (s *RolePermissionService) InitWithCancelUnderContext(ctx context.Context, db *gorm.DB) (*RolePermissionService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithCancelUnderContext(ctx, db)
	return s, cancel
}

func (s *RolePermissionService) InitWithTimeoutUnderContext(ctx context.Context, db *gorm.DB, timeout time.Duration) (*RolePermissionService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithTimeoutUnderContext(ctx, db, timeout)
	return s, cancel
}
