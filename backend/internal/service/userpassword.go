package service

import (
	"context"
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

type IUserPasswordService interface {
	crud.ICrud[entity.UserPassword]
}

func NewUserPasswordService(ctx context.Context, db *gorm.DB) (IUserPasswordService, context.CancelFunc) {
	s, cancel := new(UserPasswordService).InitWithCancelUnderContext(ctx, db)
	return s, cancel
}

func UseUserPasswordService(db *gorm.DB) IUserPasswordService {
	return new(UserPasswordService).Init(db)
}

type UserPasswordService struct {
	crud.Crud[entity.UserPassword]
}

func (s *UserPasswordService) Init(db *gorm.DB) *UserPasswordService {
	s.Crud.Init(db)
	return s
}

func (s *UserPasswordService) InitWithCancelUnderContext(ctx context.Context, db *gorm.DB) (*UserPasswordService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithCancelUnderContext(ctx, db)
	return s, cancel
}

func (s *UserPasswordService) InitWithTimeoutUnderContext(ctx context.Context, db *gorm.DB, timeout time.Duration) (*UserPasswordService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithTimeoutUnderContext(ctx, db, timeout)
	return s, cancel
}
