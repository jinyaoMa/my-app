package service

import (
	"context"
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

type IFileGroupService interface {
	crud.ICrudService[entity.FileGroup]
}

func NewFileGroupService(ctx context.Context, tx *gorm.DB) (IFileGroupService, context.CancelFunc) {
	s, cancel := new(FileGroupService).InitWithCancelUnderContext(ctx, tx)
	return s, cancel
}

func UseFileGroupService(tx *gorm.DB) IFileGroupService {
	return new(FileGroupService).Init(tx)
}

type FileGroupService struct {
	crud.Crud[entity.FileGroup]
}

func (s *FileGroupService) Init(tx *gorm.DB) *FileGroupService {
	s.Crud.Init(tx)
	return s
}

func (s *FileGroupService) InitWithCancelUnderContext(ctx context.Context, tx *gorm.DB) (*FileGroupService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithCancelUnderContext(ctx, tx)
	return s, cancel
}

func (s *FileGroupService) InitWithTimeoutUnderContext(ctx context.Context, tx *gorm.DB, timeout time.Duration) (*FileGroupService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithTimeoutUnderContext(ctx, tx, timeout)
	return s, cancel
}
