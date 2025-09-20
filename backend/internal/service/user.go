package service

import (
	"context"
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
)

type IUserService interface {
	crud.ICrud[entity.User]
	GetByAccountPassword(account string, password string, includes ...string) (user entity.User, notFound bool, err error)
}

func NewUserService(ctx context.Context, db *gorm.DB) (IUserService, context.CancelFunc) {
	s, cancel := new(UserService).InitWithCancelUnderContext(ctx, db)
	return s, cancel
}

func UseUserService(db *gorm.DB) IUserService {
	return new(UserService).Init(db)
}

type UserService struct {
	crud.Crud[entity.User]
}

func (s *UserService) GetByAccountPassword(account string, password string, includes ...string) (user entity.User, notFound bool, err error) {
	return s.FindOne(func(tx *gorm.DB) (*gorm.DB, error) {
		tx = tx.Where(entity.User{
			Account:  datatype.Encrypted(account),
			Password: datatype.Password(password),
		})
		return tx, nil
	}, includes...)
}

func (s *UserService) Init(db *gorm.DB) *UserService {
	s.Crud.Init(db)
	return s
}

func (s *UserService) InitWithCancelUnderContext(ctx context.Context, db *gorm.DB) (*UserService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithCancelUnderContext(ctx, db)
	return s, cancel
}

func (s *UserService) InitWithTimeoutUnderContext(ctx context.Context, db *gorm.DB, timeout time.Duration) (*UserService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithTimeoutUnderContext(ctx, db, timeout)
	return s, cancel
}
