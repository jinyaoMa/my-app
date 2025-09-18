package service

import (
	"context"
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

type IFileUserService interface {
	crud.ICrudService[entity.FileUser]
}

func NewFileUserService(ctx context.Context, tx *gorm.DB) (IFileUserService, context.CancelFunc) {
	s, cancel := new(FileUserService).InitWithCancelUnderContext(ctx, tx)
	return s, cancel
}

func UseFileUserService(tx *gorm.DB) IFileUserService {
	return new(FileUserService).Init(tx)
}

type FileUserService struct {
	crud.Crud[entity.FileUser]
}

func (s *FileUserService) Init(tx *gorm.DB) *FileUserService {
	s.Crud.Init(tx)
	return s
}

func (s *FileUserService) InitWithCancelUnderContext(ctx context.Context, tx *gorm.DB) (*FileUserService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithCancelUnderContext(ctx, tx)
	return s, cancel
}

func (s *FileUserService) InitWithTimeoutUnderContext(ctx context.Context, tx *gorm.DB, timeout time.Duration) (*FileUserService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithTimeoutUnderContext(ctx, tx, timeout)
	return s, cancel
}
