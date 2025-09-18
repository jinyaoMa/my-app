package service

import (
	"context"
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

type IPermissionService interface {
	crud.ICrudService[entity.Permission]
}

func NewPermissionService(ctx context.Context, tx *gorm.DB) (IPermissionService, context.CancelFunc) {
	s, cancel := new(PermissionService).InitWithCancelUnderContext(ctx, tx)
	return s, cancel
}

func UsePermissionService(tx *gorm.DB) IPermissionService {
	return new(PermissionService).Init(tx)
}

type PermissionService struct {
	crud.Crud[entity.Permission]
}

func (s *PermissionService) Init(tx *gorm.DB) *PermissionService {
	s.Crud.Init(tx)
	return s
}

func (s *PermissionService) InitWithCancelUnderContext(ctx context.Context, tx *gorm.DB) (*PermissionService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithCancelUnderContext(ctx, tx)
	return s, cancel
}

func (s *PermissionService) InitWithTimeoutUnderContext(ctx context.Context, tx *gorm.DB, timeout time.Duration) (*PermissionService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithTimeoutUnderContext(ctx, tx, timeout)
	return s, cancel
}
