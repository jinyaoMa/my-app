package service

import (
	"context"
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

type IFileExtensionService interface {
	crud.ICrud[entity.FileExtension]
}

func NewFileExtensionService(ctx context.Context, db *gorm.DB) (IFileExtensionService, context.CancelFunc) {
	s, cancel := new(FileExtensionService).InitWithCancelUnderContext(ctx, db)
	return s, cancel
}

func UseFileExtensionService(db *gorm.DB) IFileExtensionService {
	return new(FileExtensionService).Init(db)
}

type FileExtensionService struct {
	crud.Crud[entity.FileExtension]
}

func (s *FileExtensionService) Init(db *gorm.DB) *FileExtensionService {
	s.Crud.Init(db)
	return s
}

func (s *FileExtensionService) InitWithCancelUnderContext(ctx context.Context, db *gorm.DB) (*FileExtensionService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithCancelUnderContext(ctx, db)
	return s, cancel
}

func (s *FileExtensionService) InitWithTimeoutUnderContext(ctx context.Context, db *gorm.DB, timeout time.Duration) (*FileExtensionService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithTimeoutUnderContext(ctx, db, timeout)
	return s, cancel
}
