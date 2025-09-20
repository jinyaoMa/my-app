package service

import (
	"context"
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

type IGroupRoleService interface {
	crud.ICrud[entity.GroupRole]
}

func NewGroupRoleService(ctx context.Context, db *gorm.DB) (IGroupRoleService, context.CancelFunc) {
	s, cancel := new(GroupRoleService).InitWithCancelUnderContext(ctx, db)
	return s, cancel
}

func UseGroupRoleService(db *gorm.DB) IGroupRoleService {
	return new(GroupRoleService).Init(db)
}

type GroupRoleService struct {
	crud.Crud[entity.GroupRole]
}

func (s *GroupRoleService) Init(db *gorm.DB) *GroupRoleService {
	s.Crud.Init(db)
	return s
}

func (s *GroupRoleService) InitWithCancelUnderContext(ctx context.Context, db *gorm.DB) (*GroupRoleService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithCancelUnderContext(ctx, db)
	return s, cancel
}

func (s *GroupRoleService) InitWithTimeoutUnderContext(ctx context.Context, db *gorm.DB, timeout time.Duration) (*GroupRoleService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithTimeoutUnderContext(ctx, db, timeout)
	return s, cancel
}
