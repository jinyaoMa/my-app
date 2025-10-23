package service

import (
	"context"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
	"majinyao.cn/my-app/backend/pkg/db/dbcontext"
)

type IRoleService interface {
	crud.ICrud[entity.Role]
}

func NewRoleService(ctx context.Context, options crud.Options, db *gorm.DB) (IRoleService, context.CancelFunc) {
	s, cancel := new(RoleService).InitUnderContextWithCancel(ctx, options, db)
	return s, cancel
}

func UseRoleService(ctx context.Context, options crud.Options, db *gorm.DB) IRoleService {
	return new(RoleService).Init(ctx, options, db)
}

type RoleService struct {
	crud.Crud[entity.Role]
	ctx context.Context
}

func (s *RoleService) InitUnderContextWithCancel(ctx context.Context, options crud.Options, db *gorm.DB) (*RoleService, context.CancelFunc) {
	db, cancel := dbcontext.SectionUnderContextWithCancel(ctx, db)
	return s.Init(ctx, options, db), cancel
}

func (s *RoleService) Init(ctx context.Context, options crud.Options, db *gorm.DB) *RoleService {
	s.Crud.Init(options, db)
	s.ctx = ctx
	return s
}
