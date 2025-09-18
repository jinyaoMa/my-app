package service

import (
	"context"
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

type IRolePermissionService interface {
	crud.ICrudService[entity.RolePermission]
}

func NewRolePermissionService(ctx context.Context, tx *gorm.DB) (IRolePermissionService, context.CancelFunc) {
	s, cancel := new(RolePermissionService).InitWithCancelUnderContext(ctx, tx)
	return s, cancel
}

func UseRolePermissionService(tx *gorm.DB) IRolePermissionService {
	return new(RolePermissionService).Init(tx)
}

type RolePermissionService struct {
	crud.Crud[entity.RolePermission]
}

func (s *RolePermissionService) Init(tx *gorm.DB) *RolePermissionService {
	s.Crud.Init(tx)
	return s
}

func (s *RolePermissionService) InitWithCancelUnderContext(ctx context.Context, tx *gorm.DB) (*RolePermissionService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithCancelUnderContext(ctx, tx)
	return s, cancel
}

func (s *RolePermissionService) InitWithTimeoutUnderContext(ctx context.Context, tx *gorm.DB, timeout time.Duration) (*RolePermissionService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithTimeoutUnderContext(ctx, tx, timeout)
	return s, cancel
}
