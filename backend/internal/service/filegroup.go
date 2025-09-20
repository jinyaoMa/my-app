package service

import (
	"context"
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

type IFileGroupService interface {
	crud.ICrud[entity.FileGroup]
}

func NewFileGroupService(ctx context.Context, db *gorm.DB) (IFileGroupService, context.CancelFunc) {
	s, cancel := new(FileGroupService).InitWithCancelUnderContext(ctx, db)
	return s, cancel
}

func UseFileGroupService(db *gorm.DB) IFileGroupService {
	return new(FileGroupService).Init(db)
}

type FileGroupService struct {
	crud.Crud[entity.FileGroup]
}

func (s *FileGroupService) Init(db *gorm.DB) *FileGroupService {
	s.Crud.Init(db)
	return s
}

func (s *FileGroupService) InitWithCancelUnderContext(ctx context.Context, db *gorm.DB) (*FileGroupService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithCancelUnderContext(ctx, db)
	return s, cancel
}

func (s *FileGroupService) InitWithTimeoutUnderContext(ctx context.Context, db *gorm.DB, timeout time.Duration) (*FileGroupService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithTimeoutUnderContext(ctx, db, timeout)
	return s, cancel
}
