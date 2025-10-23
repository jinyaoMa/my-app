package service

import (
	"context"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
	"majinyao.cn/my-app/backend/pkg/db/dbcontext"
)

type IFileCategoryService interface {
	crud.ICrud[entity.FileCategory]
}

func NewFileCategoryService(ctx context.Context, options crud.Options, db *gorm.DB) (IFileCategoryService, context.CancelFunc) {
	s, cancel := new(FileCategoryService).InitUnderContextWithCancel(ctx, options, db)
	return s, cancel
}

func UseFileCategoryService(ctx context.Context, options crud.Options, db *gorm.DB) IFileCategoryService {
	return new(FileCategoryService).Init(ctx, options, db)
}

type FileCategoryService struct {
	crud.Crud[entity.FileCategory]
	ctx context.Context
}

func (s *FileCategoryService) InitUnderContextWithCancel(ctx context.Context, options crud.Options, db *gorm.DB) (*FileCategoryService, context.CancelFunc) {
	db, cancel := dbcontext.SectionUnderContextWithCancel(ctx, db)
	return s.Init(ctx, options, db), cancel
}

func (s *FileCategoryService) Init(ctx context.Context, options crud.Options, db *gorm.DB) *FileCategoryService {
	s.Crud.Init(options, db)
	s.ctx = ctx
	return s
}
