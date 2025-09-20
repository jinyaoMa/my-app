package service

import (
	"context"
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

type IPermissionService interface {
	crud.ICrud[entity.Permission]
}

func NewPermissionService(ctx context.Context, db *gorm.DB) (IPermissionService, context.CancelFunc) {
	s, cancel := new(PermissionService).InitWithCancelUnderContext(ctx, db)
	return s, cancel
}

func UsePermissionService(db *gorm.DB) IPermissionService {
	return new(PermissionService).Init(db)
}

type PermissionService struct {
	crud.Crud[entity.Permission]
}

func (s *PermissionService) Init(db *gorm.DB) *PermissionService {
	s.Crud.Init(db)
	return s
}

func (s *PermissionService) InitWithCancelUnderContext(ctx context.Context, db *gorm.DB) (*PermissionService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithCancelUnderContext(ctx, db)
	return s, cancel
}

func (s *PermissionService) InitWithTimeoutUnderContext(ctx context.Context, db *gorm.DB, timeout time.Duration) (*PermissionService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithTimeoutUnderContext(ctx, db, timeout)
	return s, cancel
}
