package service

import (
	"context"
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

type IOptionService interface {
	crud.ICrudService[entity.Option]
	LoadOrCreateByKey(o *entity.Option) (err error)
}

func NewOptionService(ctx context.Context, tx *gorm.DB) (IOptionService, context.CancelFunc) {
	s, cancel := new(OptionService).InitWithCancelUnderContext(ctx, tx)
	return s, cancel
}

func UseOptionService(tx *gorm.DB) IOptionService {
	return new(OptionService).Init(tx)
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

func (s *OptionService) Init(tx *gorm.DB) *OptionService {
	s.Crud.Init(tx)
	return s
}

func (s *OptionService) InitWithCancelUnderContext(ctx context.Context, tx *gorm.DB) (*OptionService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithCancelUnderContext(ctx, tx)
	return s, cancel
}

func (s *OptionService) InitWithTimeoutUnderContext(ctx context.Context, tx *gorm.DB, timeout time.Duration) (*OptionService, context.CancelFunc) {
	_, cancel := s.Crud.InitWithTimeoutUnderContext(ctx, tx, timeout)
	return s, cancel
}
