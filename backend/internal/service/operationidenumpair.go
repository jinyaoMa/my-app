package service

import (
	"context"
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
	"majinyao.cn/my-app/backend/pkg/db/dbcontext"
)

type IOperationIdEnumPairService interface {
	LoadOrCreate(ms *entity.OperationIdEnumPairs) (err error)
	List(operationId string) (entities []entity.OperationIdEnumPairs, total int64, err error)
	ListCopy(copyToEntities any, operationId string) (entities []entity.OperationIdEnumPairs, total int64, err error)
}

func NewOperationIdEnumPairService(ctx context.Context, db *gorm.DB) (IOperationIdEnumPairService, context.CancelFunc) {
	s, cancel := new(OperationIdEnumPairService).InitWithCancelUnderContext(ctx, db)
	return s, cancel
}

func UseOperationIdEnumPairService(db *gorm.DB) IOperationIdEnumPairService {
	return new(OperationIdEnumPairService).Init(db)
}

type OperationIdEnumPairService struct {
	Db           *gorm.DB
	CopierOption copier.Option
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

func (s *OperationIdEnumPairService) ListCopy(copyToEntities any, operationId string) (entities []entity.OperationIdEnumPairs, total int64, err error) {
	entities, total, err = s.List(operationId)
	if err != nil {
		return
	}
	err = copier.CopyWithOption(copyToEntities, &entities, s.CopierOption)
	if err != nil {
		return
	}
	return
}

func (s *OperationIdEnumPairService) Init(db *gorm.DB) *OperationIdEnumPairService {
	s.Db = db
	s.CopierOption = crud.DefaultCopierOption
	return s
}

func (s *OperationIdEnumPairService) InitWithCancelUnderContext(ctx context.Context, db *gorm.DB) (*OperationIdEnumPairService, context.CancelFunc) {
	db, cancel := dbcontext.SectionUnderContextWithCancel(ctx, db)
	return s.Init(db), cancel
}

func (s *OperationIdEnumPairService) InitWithTimeoutUnderContext(ctx context.Context, db *gorm.DB, timeout time.Duration) (*OperationIdEnumPairService, context.CancelFunc) {
	db, cancel := dbcontext.SectionUnderContextWithTimeout(ctx, db, timeout)
	return s.Init(db), cancel
}
