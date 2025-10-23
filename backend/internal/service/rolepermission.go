package service

import (
	"context"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
	"majinyao.cn/my-app/backend/pkg/db/dbcontext"
)

type IRolePermissionService interface {
	crud.ICrud[entity.RolePermission]
}

func NewRolePermissionService(ctx context.Context, options crud.Options, db *gorm.DB) (IRolePermissionService, context.CancelFunc) {
	s, cancel := new(RolePermissionService).InitUnderContextWithCancel(ctx, options, db)
	return s, cancel
}

func UseRolePermissionService(ctx context.Context, options crud.Options, db *gorm.DB) IRolePermissionService {
	return new(RolePermissionService).Init(ctx, options, db)
}

type RolePermissionService struct {
	crud.Crud[entity.RolePermission]
	ctx context.Context
}

func (s *RolePermissionService) InitUnderContextWithCancel(ctx context.Context, options crud.Options, db *gorm.DB) (*RolePermissionService, context.CancelFunc) {
	db, cancel := dbcontext.SectionUnderContextWithCancel(ctx, db)
	return s.Init(ctx, options, db), cancel
}

func (s *RolePermissionService) Init(ctx context.Context, options crud.Options, db *gorm.DB) *RolePermissionService {
	s.Crud.Init(options, db)
	s.ctx = ctx
	return s
}
