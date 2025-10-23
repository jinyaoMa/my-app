package service

import (
	"context"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
	"majinyao.cn/my-app/backend/pkg/db/dbcontext"
)

type IFileGroupService interface {
	crud.ICrud[entity.FileGroup]
}

func NewFileGroupService(ctx context.Context, options crud.Options, db *gorm.DB) (IFileGroupService, context.CancelFunc) {
	s, cancel := new(FileGroupService).InitUnderContextWithCancel(ctx, options, db)
	return s, cancel
}

func UseFileGroupService(ctx context.Context, options crud.Options, db *gorm.DB) IFileGroupService {
	return new(FileGroupService).Init(ctx, options, db)
}

type FileGroupService struct {
	crud.Crud[entity.FileGroup]
	ctx context.Context
}

func (s *FileGroupService) InitUnderContextWithCancel(ctx context.Context, options crud.Options, db *gorm.DB) (*FileGroupService, context.CancelFunc) {
	db, cancel := dbcontext.SectionUnderContextWithCancel(ctx, db)
	return s.Init(ctx, options, db), cancel
}

func (s *FileGroupService) Init(ctx context.Context, options crud.Options, db *gorm.DB) *FileGroupService {
	s.Crud.Init(options, db)
	s.ctx = ctx
	return s
}
