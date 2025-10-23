package service

import (
	"context"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
	"majinyao.cn/my-app/backend/pkg/db/dbcontext"
)

type IOptionService interface {
	crud.ICrud[entity.Option]
	LoadOrCreateByKey(o *entity.Option) (err error)
}

func NewOptionService(ctx context.Context, options crud.Options, db *gorm.DB) (IOptionService, context.CancelFunc) {
	s, cancel := new(OptionService).InitUnderContextWithCancel(ctx, options, db)
	return s, cancel
}

func UseOptionService(ctx context.Context, options crud.Options, db *gorm.DB) IOptionService {
	return new(OptionService).Init(ctx, options, db)
}

type OptionService struct {
	crud.Crud[entity.Option]
	ctx context.Context
}

func (s *OptionService) LoadOrCreateByKey(o *entity.Option) (err error) {
	res := s.Db.FirstOrCreate(o, entity.Option{
		Key: o.Key,
	})
	err = res.Error
	return
}

func (s *OptionService) InitUnderContextWithCancel(ctx context.Context, options crud.Options, db *gorm.DB) (*OptionService, context.CancelFunc) {
	db, cancel := dbcontext.SectionUnderContextWithCancel(ctx, db)
	return s.Init(ctx, options, db), cancel
}

func (s *OptionService) Init(ctx context.Context, options crud.Options, db *gorm.DB) *OptionService {
	s.Crud.Init(options, db)
	s.ctx = ctx
	return s
}
