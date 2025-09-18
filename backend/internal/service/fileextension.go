package service

import (
	"context"
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

type IFileExtensionService interface {
	crud.ICrudService[entity.FileExtension]
}

func NewFileExtensionService(ctx context.Context, tx *gorm.DB) (IFileExtensionService, context.CancelFunc) {
	s, cancel := new(FileExtensionService).InitWithCancelUnderContext(ctx, tx)
	return s, cancel
}

func UseFileExtensionService(tx *gorm.DB) IFileExtensionService {
	return new(FileExtensionService).Init(tx)
}

type FileExtensionService struct {
	crud.Crud[entity.FileExtension]
}

func (s *FileExtensionService) Init(tx *gorm.DB) *FileExtensionService {
	s.Crud.Init(tx)
	return s
}

func (s *FileExtensionService) InitWithCancelUnderContext(ctx context.Context, tx *gorm.DB) (*FileExtensionService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithCancelUnderContext(ctx, tx)
	return s, cancel
}

func (s *FileExtensionService) InitWithTimeoutUnderContext(ctx context.Context, tx *gorm.DB, timeout time.Duration) (*FileExtensionService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithTimeoutUnderContext(ctx, tx, timeout)
	return s, cancel
}
