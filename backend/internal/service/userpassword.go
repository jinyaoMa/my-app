package service

import (
	"context"
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

type IUserPasswordService interface {
	crud.ICrudService[entity.UserPassword]
}

func NewUserPasswordService(ctx context.Context, tx *gorm.DB) (IUserPasswordService, context.CancelFunc) {
	s, cancel := new(UserPasswordService).InitWithCancelUnderContext(ctx, tx)
	return s, cancel
}

func UseUserPasswordService(tx *gorm.DB) IUserPasswordService {
	return new(UserPasswordService).Init(tx)
}

type UserPasswordService struct {
	crud.Crud[entity.UserPassword]
}

func (s *UserPasswordService) Init(tx *gorm.DB) *UserPasswordService {
	s.Crud.Init(tx)
	return s
}

func (s *UserPasswordService) InitWithCancelUnderContext(ctx context.Context, tx *gorm.DB) (*UserPasswordService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithCancelUnderContext(ctx, tx)
	return s, cancel
}

func (s *UserPasswordService) InitWithTimeoutUnderContext(ctx context.Context, tx *gorm.DB, timeout time.Duration) (*UserPasswordService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithTimeoutUnderContext(ctx, tx, timeout)
	return s, cancel
}
