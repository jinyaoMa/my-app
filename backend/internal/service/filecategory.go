package service

import (
	"context"
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

type IFileCategoryService interface {
	crud.ICrudService[entity.FileCategory]
}

func NewFileCategoryService(ctx context.Context, tx *gorm.DB) (IFileCategoryService, context.CancelFunc) {
	s, cancel := new(FileCategoryService).InitWithCancelUnderContext(ctx, tx)
	return s, cancel
}

func UseFileCategoryService(tx *gorm.DB) IFileCategoryService {
	return new(FileCategoryService).Init(tx)
}

type FileCategoryService struct {
	crud.Crud[entity.FileCategory]
}

func (s *FileCategoryService) Init(tx *gorm.DB) *FileCategoryService {
	s.Crud.Init(tx)
	return s
}

func (s *FileCategoryService) InitWithCancelUnderContext(ctx context.Context, tx *gorm.DB) (*FileCategoryService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithCancelUnderContext(ctx, tx)
	return s, cancel
}

func (s *FileCategoryService) InitWithTimeoutUnderContext(ctx context.Context, tx *gorm.DB, timeout time.Duration) (*FileCategoryService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithTimeoutUnderContext(ctx, tx, timeout)
	return s, cancel
}
