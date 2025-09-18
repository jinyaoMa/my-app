package service

import (
	"context"
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

type IRoleService interface {
	crud.ICrudService[entity.Role]
}

func NewRoleService(ctx context.Context, tx *gorm.DB) (IRoleService, context.CancelFunc) {
	s, cancel := new(RoleService).InitWithCancelUnderContext(ctx, tx)
	return s, cancel
}

func UseRoleService(tx *gorm.DB) IRoleService {
	return new(RoleService).Init(tx)
}

type RoleService struct {
	crud.Crud[entity.Role]
}

func (s *RoleService) Init(tx *gorm.DB) *RoleService {
	s.Crud.Init(tx)
	return s
}

func (s *RoleService) InitWithCancelUnderContext(ctx context.Context, tx *gorm.DB) (*RoleService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithCancelUnderContext(ctx, tx)
	return s, cancel
}

func (s *RoleService) InitWithTimeoutUnderContext(ctx context.Context, tx *gorm.DB, timeout time.Duration) (*RoleService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithTimeoutUnderContext(ctx, tx, timeout)
	return s, cancel
}
