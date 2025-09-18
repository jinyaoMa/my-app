package service

import (
	"context"
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

type IGroupService interface {
	crud.ICrudService[entity.Group]
}

func NewGroupService(ctx context.Context, tx *gorm.DB) (IGroupService, context.CancelFunc) {
	s, cancel := new(GroupService).InitWithCancelUnderContext(ctx, tx)
	return s, cancel
}

func UseGroupService(tx *gorm.DB) IGroupService {
	return new(GroupService).Init(tx)
}

type GroupService struct {
	crud.Crud[entity.Group]
}

func (s *GroupService) Init(tx *gorm.DB) *GroupService {
	s.Crud.Init(tx)
	return s
}

func (s *GroupService) InitWithCancelUnderContext(ctx context.Context, tx *gorm.DB) (*GroupService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithCancelUnderContext(ctx, tx)
	return s, cancel
}

func (s *GroupService) InitWithTimeoutUnderContext(ctx context.Context, tx *gorm.DB, timeout time.Duration) (*GroupService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithTimeoutUnderContext(ctx, tx, timeout)
	return s, cancel
}
