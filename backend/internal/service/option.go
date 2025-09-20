package service

import (
	"context"
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

type IOptionService interface {
	crud.ICrud[entity.Option]
	LoadOrCreateByKey(o *entity.Option) (err error)
}

func NewOptionService(ctx context.Context, db *gorm.DB) (IOptionService, context.CancelFunc) {
	s, cancel := new(OptionService).InitWithCancelUnderContext(ctx, db)
	return s, cancel
}

func UseOptionService(db *gorm.DB) IOptionService {
	return new(OptionService).Init(db)
}

type OptionService struct {
	crud.Crud[entity.Option]
}

func (s *OptionService) LoadOrCreateByKey(o *entity.Option) (err error) {
	res := s.Db.FirstOrCreate(o, entity.Option{
		Key: o.Key,
	})
	err = res.Error
	return
}

func (s *OptionService) Init(db *gorm.DB) *OptionService {
	s.Crud.Init(db)
	return s
}

func (s *OptionService) InitWithCancelUnderContext(ctx context.Context, db *gorm.DB) (*OptionService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithCancelUnderContext(ctx, db)
	return s, cancel
}

func (s *OptionService) InitWithTimeoutUnderContext(ctx context.Context, db *gorm.DB, timeout time.Duration) (*OptionService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithTimeoutUnderContext(ctx, db, timeout)
	return s, cancel
}
