package service

import (
	"context"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
	"majinyao.cn/my-app/backend/pkg/db/dbcontext"
)

type IGroupService interface {
	crud.ICrud[entity.Group]
}

func NewGroupService(ctx context.Context, options crud.Options, db *gorm.DB) (IGroupService, context.CancelFunc) {
	s, cancel := new(GroupService).InitUnderContextWithCancel(ctx, options, db)
	return s, cancel
}

func UseGroupService(ctx context.Context, options crud.Options, db *gorm.DB) IGroupService {
	return new(GroupService).Init(ctx, options, db)
}

type GroupService struct {
	crud.Crud[entity.Group]
	ctx context.Context
}

func (s *GroupService) InitUnderContextWithCancel(ctx context.Context, options crud.Options, db *gorm.DB) (*GroupService, context.CancelFunc) {
	db, cancel := dbcontext.SectionUnderContextWithCancel(ctx, db)
	return s.Init(ctx, options, db), cancel
}

func (s *GroupService) Init(ctx context.Context, options crud.Options, db *gorm.DB) *GroupService {
	s.Crud.Init(options, db)
	s.ctx = ctx
	return s
}
