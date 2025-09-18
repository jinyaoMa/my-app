package service

import (
	"context"
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

type IFileService interface {
	crud.ICrudService[entity.File]
}

func NewFileService(ctx context.Context, tx *gorm.DB) (IFileService, context.CancelFunc) {
	s, cancel := new(FileService).InitWithCancelUnderContext(ctx, tx)
	return s, cancel
}

func UseFileService(tx *gorm.DB) IFileService {
	return new(FileService).Init(tx)
}

type FileService struct {
	crud.Crud[entity.File]
}

func (s *FileService) Init(tx *gorm.DB) *FileService {
	s.Crud.Init(tx)
	return s
}

func (s *FileService) InitWithCancelUnderContext(ctx context.Context, tx *gorm.DB) (*FileService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithCancelUnderContext(ctx, tx)
	return s, cancel
}

func (s *FileService) InitWithTimeoutUnderContext(ctx context.Context, tx *gorm.DB, timeout time.Duration) (*FileService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithTimeoutUnderContext(ctx, tx, timeout)
	return s, cancel
}
