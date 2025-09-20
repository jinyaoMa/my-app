package service

import (
	"context"
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

type IGroupService interface {
	crud.ICrud[entity.Group]
}

func NewGroupService(ctx context.Context, db *gorm.DB) (IGroupService, context.CancelFunc) {
	s, cancel := new(GroupService).InitWithCancelUnderContext(ctx, db)
	return s, cancel
}

func UseGroupService(db *gorm.DB) IGroupService {
	return new(GroupService).Init(db)
}

type GroupService struct {
	crud.Crud[entity.Group]
}

func (s *GroupService) Init(db *gorm.DB) *GroupService {
	s.Crud.Init(db)
	return s
}

func (s *GroupService) InitWithCancelUnderContext(ctx context.Context, db *gorm.DB) (*GroupService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithCancelUnderContext(ctx, db)
	return s, cancel
}

func (s *GroupService) InitWithTimeoutUnderContext(ctx context.Context, db *gorm.DB, timeout time.Duration) (*GroupService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithTimeoutUnderContext(ctx, db, timeout)
	return s, cancel
}
