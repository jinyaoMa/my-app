package service

import (
	"context"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
	"majinyao.cn/my-app/backend/pkg/db/dbcontext"
)

type IFileUserService interface {
	crud.ICrud[entity.FileUser]
}

func NewFileUserService(ctx context.Context, options crud.Options, db *gorm.DB) (IFileUserService, context.CancelFunc) {
	s, cancel := new(FileUserService).InitUnderContextWithCancel(ctx, options, db)
	return s, cancel
}

func UseFileUserService(ctx context.Context, options crud.Options, db *gorm.DB) IFileUserService {
	return new(FileUserService).Init(ctx, options, db)
}

type FileUserService struct {
	crud.Crud[entity.FileUser]
	ctx context.Context
}

func (s *FileUserService) InitUnderContextWithCancel(ctx context.Context, options crud.Options, db *gorm.DB) (*FileUserService, context.CancelFunc) {
	db, cancel := dbcontext.SectionUnderContextWithCancel(ctx, db)
	return s.Init(ctx, options, db), cancel
}

func (s *FileUserService) Init(ctx context.Context, options crud.Options, db *gorm.DB) *FileUserService {
	s.Crud.Init(options, db)
	s.ctx = ctx
	return s
}
