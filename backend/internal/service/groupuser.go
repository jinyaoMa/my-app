package service

import (
	"context"
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

type IGroupUserService interface {
	crud.ICrud[entity.GroupUser]
}

func NewGroupUserService(ctx context.Context, db *gorm.DB) (IGroupUserService, context.CancelFunc) {
	s, cancel := new(GroupUserService).InitWithCancelUnderContext(ctx, db)
	return s, cancel
}

func UseGroupUserService(db *gorm.DB) IGroupUserService {
	return new(GroupUserService).Init(db)
}

type GroupUserService struct {
	crud.Crud[entity.GroupUser]
}

func (s *GroupUserService) Init(db *gorm.DB) *GroupUserService {
	s.Crud.Init(db)
	return s
}

func (s *GroupUserService) InitWithCancelUnderContext(ctx context.Context, db *gorm.DB) (*GroupUserService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithCancelUnderContext(ctx, db)
	return s, cancel
}

func (s *GroupUserService) InitWithTimeoutUnderContext(ctx context.Context, db *gorm.DB, timeout time.Duration) (*GroupUserService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithTimeoutUnderContext(ctx, db, timeout)
	return s, cancel
}
