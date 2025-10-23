package service

import (
	"context"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
	"majinyao.cn/my-app/backend/pkg/db/dbcontext"
)

type IUserRoleService interface {
	crud.ICrud[entity.UserRole]
}

func NewUserRoleService(ctx context.Context, options crud.Options, db *gorm.DB) (IUserRoleService, context.CancelFunc) {
	s, cancel := new(UserRoleService).InitUnderContextWithCancel(ctx, options, db)
	return s, cancel
}

func UseUserRoleService(ctx context.Context, options crud.Options, db *gorm.DB) IUserRoleService {
	return new(UserRoleService).Init(ctx, options, db)
}

type UserRoleService struct {
	crud.Crud[entity.UserRole]
	ctx context.Context
}

func (s *UserRoleService) InitUnderContextWithCancel(ctx context.Context, options crud.Options, db *gorm.DB) (*UserRoleService, context.CancelFunc) {
	db, cancel := dbcontext.SectionUnderContextWithCancel(ctx, db)
	return s.Init(ctx, options, db), cancel
}

func (s *UserRoleService) Init(ctx context.Context, options crud.Options, db *gorm.DB) *UserRoleService {
	s.Crud.Init(options, db)
	s.ctx = ctx
	return s
}
