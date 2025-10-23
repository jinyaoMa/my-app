package service

import (
	"context"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
	"majinyao.cn/my-app/backend/pkg/db/dbcontext"
)

type IPermissionService interface {
	crud.ICrud[entity.Permission]
}

func NewPermissionService(ctx context.Context, options crud.Options, db *gorm.DB) (IPermissionService, context.CancelFunc) {
	s, cancel := new(PermissionService).InitUnderContextWithCancel(ctx, options, db)
	return s, cancel
}

func UsePermissionService(ctx context.Context, options crud.Options, db *gorm.DB) IPermissionService {
	return new(PermissionService).Init(ctx, options, db)
}

type PermissionService struct {
	crud.Crud[entity.Permission]
	ctx context.Context
}

func (s *PermissionService) InitUnderContextWithCancel(ctx context.Context, options crud.Options, db *gorm.DB) (*PermissionService, context.CancelFunc) {
	db, cancel := dbcontext.SectionUnderContextWithCancel(ctx, db)
	return s.Init(ctx, options, db), cancel
}

func (s *PermissionService) Init(ctx context.Context, options crud.Options, db *gorm.DB) *PermissionService {
	s.Crud.Init(options, db)
	s.ctx = ctx
	return s
}
