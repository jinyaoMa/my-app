package service

import (
	"context"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
	"majinyao.cn/my-app/backend/pkg/db/dbcontext"
)

type IUserPasswordService interface {
	crud.ICrud[entity.UserPassword]
}

func NewUserPasswordService(ctx context.Context, options crud.Options, db *gorm.DB) (IUserPasswordService, context.CancelFunc) {
	s, cancel := new(UserPasswordService).InitUnderContextWithCancel(ctx, options, db)
	return s, cancel
}

func UseUserPasswordService(ctx context.Context, options crud.Options, db *gorm.DB) IUserPasswordService {
	return new(UserPasswordService).Init(ctx, options, db)
}

type UserPasswordService struct {
	crud.Crud[entity.UserPassword]
	ctx context.Context
}

func (s *UserPasswordService) InitUnderContextWithCancel(ctx context.Context, options crud.Options, db *gorm.DB) (*UserPasswordService, context.CancelFunc) {
	db, cancel := dbcontext.SectionUnderContextWithCancel(ctx, db)
	return s.Init(ctx, options, db), cancel
}

func (s *UserPasswordService) Init(ctx context.Context, options crud.Options, db *gorm.DB) *UserPasswordService {
	s.Crud.Init(options, db)
	s.ctx = ctx
	return s
}
