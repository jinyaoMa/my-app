package service

import (
	"context"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
	"majinyao.cn/my-app/backend/pkg/db/dbcontext"
)

type IGroupRoleService interface {
	crud.ICrud[entity.GroupRole]
}

func NewGroupRoleService(ctx context.Context, options crud.Options, db *gorm.DB) (IGroupRoleService, context.CancelFunc) {
	s, cancel := new(GroupRoleService).InitUnderContextWithCancel(ctx, options, db)
	return s, cancel
}

func UseGroupRoleService(ctx context.Context, options crud.Options, db *gorm.DB) IGroupRoleService {
	return new(GroupRoleService).Init(ctx, options, db)
}

type GroupRoleService struct {
	crud.Crud[entity.GroupRole]
	ctx context.Context
}

func (s *GroupRoleService) InitUnderContextWithCancel(ctx context.Context, options crud.Options, db *gorm.DB) (*GroupRoleService, context.CancelFunc) {
	db, cancel := dbcontext.SectionUnderContextWithCancel(ctx, db)
	return s.Init(ctx, options, db), cancel
}

func (s *GroupRoleService) Init(ctx context.Context, options crud.Options, db *gorm.DB) *GroupRoleService {
	s.Crud.Init(options, db)
	s.ctx = ctx
	return s
}
