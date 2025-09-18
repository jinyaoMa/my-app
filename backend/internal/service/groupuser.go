package service

import (
	"context"
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

type IGroupUserService interface {
	crud.ICrudService[entity.GroupUser]
}

func NewGroupUserService(ctx context.Context, tx *gorm.DB) (IGroupUserService, context.CancelFunc) {
	s, cancel := new(GroupUserService).InitWithCancelUnderContext(ctx, tx)
	return s, cancel
}

func UseGroupUserService(tx *gorm.DB) IGroupUserService {
	return new(GroupUserService).Init(tx)
}

type GroupUserService struct {
	crud.Crud[entity.GroupUser]
}

func (s *GroupUserService) Init(tx *gorm.DB) *GroupUserService {
	s.Crud.Init(tx)
	return s
}

func (s *GroupUserService) InitWithCancelUnderContext(ctx context.Context, tx *gorm.DB) (*GroupUserService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithCancelUnderContext(ctx, tx)
	return s, cancel
}

func (s *GroupUserService) InitWithTimeoutUnderContext(ctx context.Context, tx *gorm.DB, timeout time.Duration) (*GroupUserService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithTimeoutUnderContext(ctx, tx, timeout)
	return s, cancel
}
