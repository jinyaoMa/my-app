package service

import (
	"context"
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

type IFileService interface {
	crud.ICrud[entity.File]
}

func NewFileService(ctx context.Context, db *gorm.DB) (IFileService, context.CancelFunc) {
	s, cancel := new(FileService).InitWithCancelUnderContext(ctx, db)
	return s, cancel
}

func UseFileService(db *gorm.DB) IFileService {
	return new(FileService).Init(db)
}

type FileService struct {
	crud.Crud[entity.File]
}

func (s *FileService) Init(db *gorm.DB) *FileService {
	s.Crud.Init(db)
	return s
}

func (s *FileService) InitWithCancelUnderContext(ctx context.Context, db *gorm.DB) (*FileService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithCancelUnderContext(ctx, db)
	return s, cancel
}

func (s *FileService) InitWithTimeoutUnderContext(ctx context.Context, db *gorm.DB, timeout time.Duration) (*FileService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithTimeoutUnderContext(ctx, db, timeout)
	return s, cancel
}
