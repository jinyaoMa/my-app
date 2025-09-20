package service

import (
	"context"
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

type IFileCategoryService interface {
	crud.ICrud[entity.FileCategory]
}

func NewFileCategoryService(ctx context.Context, db *gorm.DB) (IFileCategoryService, context.CancelFunc) {
	s, cancel := new(FileCategoryService).InitWithCancelUnderContext(ctx, db)
	return s, cancel
}

func UseFileCategoryService(db *gorm.DB) IFileCategoryService {
	return new(FileCategoryService).Init(db)
}

type FileCategoryService struct {
	crud.Crud[entity.FileCategory]
}

func (s *FileCategoryService) Init(db *gorm.DB) *FileCategoryService {
	s.Crud.Init(db)
	return s
}

func (s *FileCategoryService) InitWithCancelUnderContext(ctx context.Context, db *gorm.DB) (*FileCategoryService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithCancelUnderContext(ctx, db)
	return s, cancel
}

func (s *FileCategoryService) InitWithTimeoutUnderContext(ctx context.Context, db *gorm.DB, timeout time.Duration) (*FileCategoryService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithTimeoutUnderContext(ctx, db, timeout)
	return s, cancel
}
