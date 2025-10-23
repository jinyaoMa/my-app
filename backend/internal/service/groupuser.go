package service

import (
	"context"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
	"majinyao.cn/my-app/backend/pkg/db/dbcontext"
)

type IGroupUserService interface {
	crud.ICrud[entity.GroupUser]
}

func NewGroupUserService(ctx context.Context, options crud.Options, db *gorm.DB) (IGroupUserService, context.CancelFunc) {
	s, cancel := new(GroupUserService).InitUnderContextWithCancel(ctx, options, db)
	return s, cancel
}

func UseGroupUserService(ctx context.Context, options crud.Options, db *gorm.DB) IGroupUserService {
	return new(GroupUserService).Init(ctx, options, db)
}

type GroupUserService struct {
	crud.Crud[entity.GroupUser]
	ctx context.Context
}

func (s *GroupUserService) InitUnderContextWithCancel(ctx context.Context, options crud.Options, db *gorm.DB) (*GroupUserService, context.CancelFunc) {
	db, cancel := dbcontext.SectionUnderContextWithCancel(ctx, db)
	return s.Init(ctx, options, db), cancel
}

func (s *GroupUserService) Init(ctx context.Context, options crud.Options, db *gorm.DB) *GroupUserService {
	s.Crud.Init(options, db)
	s.ctx = ctx
	return s
}
