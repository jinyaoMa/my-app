package service

import (
	"context"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
	"majinyao.cn/my-app/backend/pkg/db/dbcontext"
)

type IFileExtensionService interface {
	crud.ICrud[entity.FileExtension]
}

func NewFileExtensionService(ctx context.Context, options crud.Options, db *gorm.DB) (IFileExtensionService, context.CancelFunc) {
	s, cancel := new(FileExtensionService).InitUnderContextWithCancel(ctx, options, db)
	return s, cancel
}

func UseFileExtensionService(ctx context.Context, options crud.Options, db *gorm.DB) IFileExtensionService {
	return new(FileExtensionService).Init(ctx, options, db)
}

type FileExtensionService struct {
	crud.Crud[entity.FileExtension]
	ctx context.Context
}

func (s *FileExtensionService) InitUnderContextWithCancel(ctx context.Context, options crud.Options, db *gorm.DB) (*FileExtensionService, context.CancelFunc) {
	db, cancel := dbcontext.SectionUnderContextWithCancel(ctx, db)
	return s.Init(ctx, options, db), cancel
}

func (s *FileExtensionService) Init(ctx context.Context, options crud.Options, db *gorm.DB) *FileExtensionService {
	s.Crud.Init(options, db)
	s.ctx = ctx
	return s
}
