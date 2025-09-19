package service

import (
	"context"
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/dbcontext"
)

type IOperationIdEnumPairService interface {
	LoadOrCreate(ms *entity.OperationIdEnumPairs) (err error)
	List(operationId string) (entities []entity.OperationIdEnumPairs, total int64, err error)
}

func NewOperationIdEnumPairService(ctx context.Context, tx *gorm.DB) (IOperationIdEnumPairService, context.CancelFunc) {
	s, cancel := new(OperationIdEnumPairService).InitWithCancelUnderContext(ctx, tx)
	return s, cancel
}

func UseOperationIdEnumPairService(tx *gorm.DB) IOperationIdEnumPairService {
	return new(OperationIdEnumPairService).Init(tx)
}

type OperationIdEnumPairService struct {
	Db *gorm.DB
}

func (s *OperationIdEnumPairService) LoadOrCreate(ms *entity.OperationIdEnumPairs) (err error) {
	if ms == nil {
		return nil
	}

	return s.Db.Transaction(func(tx *gorm.DB) error {
		for i := range *ms {
			res := tx.FirstOrCreate(&(*ms)[i], entity.OperationIdEnumPair{
				OperationId: (*ms)[i].OperationId,
			})
			if res.Error != nil {
				return res.Error
			}
		}
		return nil
	})
}

func (s *OperationIdEnumPairService) List(operationId string) (entities []entity.OperationIdEnumPairs, total int64, err error) {
	tx := s.Db.Model(new(entity.OperationIdEnumPair))
	res := tx.Where("operation_id LIKE ?", "%"+operationId+"%").Find(&entities)
	total = res.RowsAffected
	err = res.Error
	return
}

func (s *OperationIdEnumPairService) Init(tx *gorm.DB) *OperationIdEnumPairService {
	s.Db = tx
	return s
}

func (s *OperationIdEnumPairService) InitWithCancelUnderContext(ctx context.Context, tx *gorm.DB) (*OperationIdEnumPairService, context.CancelFunc) {
	tx, cancel := dbcontext.SectionUnderContextWithCancel(ctx, tx)
	return s.Init(tx), cancel
}

func (s *OperationIdEnumPairService) InitWithTimeoutUnderContext(ctx context.Context, tx *gorm.DB, timeout time.Duration) (*OperationIdEnumPairService, context.CancelFunc) {
	tx, cancel := dbcontext.SectionUnderContextWithTimeout(ctx, tx, timeout)
	return s.Init(tx), cancel
}
