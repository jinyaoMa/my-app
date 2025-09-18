package service

import (
	"context"
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/api/schemas"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/api/endpoint/authbase"
	"majinyao.cn/my-app/backend/pkg/db"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

type IUserService interface {
	authbase.Verifier[schemas.UserData]
	crud.ICrudService[entity.User]
}

func NewUserService(ctx context.Context, tx *gorm.DB) (IUserService, context.CancelFunc) {
	s, cancel := new(UserService).InitWithCancelUnderContext(ctx, tx)
	return s, cancel
}

func UseUserService(tx *gorm.DB) IUserService {
	return new(UserService).Init(tx)
}

type UserService struct {
	crud.Crud[entity.User]
}

func (s *UserService) VerifyUserData(userdata schemas.UserData) (newUserdata schemas.UserData, err error) {
	return userdata, nil
}

func (s *UserService) VerifyLogin(input *authbase.LoginInput) (userdata schemas.UserData, err error) {
	var userId int64
	return schemas.UserData{
		Identity: db.ConvertIdToString(userId) + "_" + input.VisitorId,
		UserId:   userId,
	}, nil
}

func (s *UserService) Init(tx *gorm.DB) *UserService {
	s.Crud.Init(tx)
	return s
}

func (s *UserService) InitWithCancelUnderContext(ctx context.Context, tx *gorm.DB) (*UserService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithCancelUnderContext(ctx, tx)
	return s, cancel
}

func (s *UserService) InitWithTimeoutUnderContext(ctx context.Context, tx *gorm.DB, timeout time.Duration) (*UserService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithTimeoutUnderContext(ctx, tx, timeout)
	return s, cancel
}
