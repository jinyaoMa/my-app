package service

import (
	"context"
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

type IGroupRoleService interface {
	crud.ICrudService[entity.GroupRole]
}

func NewGroupRoleService(ctx context.Context, tx *gorm.DB) (IGroupRoleService, context.CancelFunc) {
	s, cancel := new(GroupRoleService).InitWithCancelUnderContext(ctx, tx)
	return s, cancel
}

func UseGroupRoleService(tx *gorm.DB) IGroupRoleService {
	return new(GroupRoleService).Init(tx)
}

type GroupRoleService struct {
	crud.Crud[entity.GroupRole]
}

func (s *GroupRoleService) Init(tx *gorm.DB) *GroupRoleService {
	s.Crud.Init(tx)
	return s
}

func (s *GroupRoleService) InitWithCancelUnderContext(ctx context.Context, tx *gorm.DB) (*GroupRoleService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithCancelUnderContext(ctx, tx)
	return s, cancel
}

func (s *GroupRoleService) InitWithTimeoutUnderContext(ctx context.Context, tx *gorm.DB, timeout time.Duration) (*GroupRoleService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithTimeoutUnderContext(ctx, tx, timeout)
	return s, cancel
}
