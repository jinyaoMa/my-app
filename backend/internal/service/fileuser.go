package service

import (
	"context"
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

type IFileUserService interface {
	crud.ICrud[entity.FileUser]
}

func NewFileUserService(ctx context.Context, db *gorm.DB) (IFileUserService, context.CancelFunc) {
	s, cancel := new(FileUserService).InitWithCancelUnderContext(ctx, db)
	return s, cancel
}

func UseFileUserService(db *gorm.DB) IFileUserService {
	return new(FileUserService).Init(db)
}

type FileUserService struct {
	crud.Crud[entity.FileUser]
}

func (s *FileUserService) Init(db *gorm.DB) *FileUserService {
	s.Crud.Init(db)
	return s
}

func (s *FileUserService) InitWithCancelUnderContext(ctx context.Context, db *gorm.DB) (*FileUserService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithCancelUnderContext(ctx, db)
	return s, cancel
}

func (s *FileUserService) InitWithTimeoutUnderContext(ctx context.Context, db *gorm.DB, timeout time.Duration) (*FileUserService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithTimeoutUnderContext(ctx, db, timeout)
	return s, cancel
}
