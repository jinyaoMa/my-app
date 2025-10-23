package service

import (
	"context"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/db/crud"
	"majinyao.cn/my-app/backend/pkg/db/dbcontext"
	"majinyao.cn/my-app/backend/pkg/flag"
)

type IOperationIdEnumPairService interface {
	LoadOrCreate(ms *entity.OperationIdEnumPairs) (err error)
	List(operationId string) (entities entity.OperationIdEnumPairs, total int64, err error)
	ListCopy(copyToEntities any, operationId string) (entities entity.OperationIdEnumPairs, total int64, err error)
	ListFromFlagBase64(flagBase64 string) (entities entity.OperationIdEnumPairs, total int64, err error)
	ListCopyFromFlagBase64(copyToEntities any, flagBase64 string) (entities entity.OperationIdEnumPairs, total int64, err error)
	ConvertEnumsToFlag(enums []int) (f flag.IFlag, err error)
}

func NewOperationIdEnumPairService(ctx context.Context, db *gorm.DB) (IOperationIdEnumPairService, context.CancelFunc) {
	s, cancel := new(OperationIdEnumPairService).InitUnderContextWithCancel(ctx, db)
	return s, cancel
}

func UseOperationIdEnumPairService(ctx context.Context, db *gorm.DB) IOperationIdEnumPairService {
	return new(OperationIdEnumPairService).Init(ctx, db)
}

type OperationIdEnumPairService struct {
	ctx context.Context
	Db  *gorm.DB
}

func (s *OperationIdEnumPairService) ListFromFlagBase64(flagBase64 string) (entities entity.OperationIdEnumPairs, total int64, err error) {
	f, err := flag.FromBase64(flagBase64)
	if err != nil {
		return
	}
	ps := f.ToPositions()
  if len(ps) == 0 {
		return
	}

	tx := s.Db.Model(new(entity.OperationIdEnumPair))
	result := tx.Find(&entities, ps)
	total = result.RowsAffected
	err = result.Error
	return
}

func (s *OperationIdEnumPairService) ListCopyFromFlagBase64(copyToEntities any, flagBase64 string) (entities entity.OperationIdEnumPairs, total int64, err error) {
	entities, total, err = s.ListFromFlagBase64(flagBase64)
	if err != nil {
		return
	}
	err = copier.CopyWithOption(copyToEntities, &entities, crud.GlobalCopierOption)
	if err != nil {
		return
	}
	return
}

func (s *OperationIdEnumPairService) ConvertEnumsToFlag(enums []int) (f flag.IFlag, err error) {
	lastOperationIdEnumPair := entity.OperationIdEnumPair{}
	if res := s.Db.Order("enum DESC").First(&lastOperationIdEnumPair); res.Error != nil {
		return nil, res.Error
	}

	f = flag.Make(lastOperationIdEnumPair.Enum+1, false)
	for _, enum := range enums {
		if enum > lastOperationIdEnumPair.Enum {
			continue
		}
		f = f.TurnOn(enum, true)
	}
	return
}

func (s *OperationIdEnumPairService) LoadOrCreate(ms *entity.OperationIdEnumPairs) (err error) {
	if ms == nil {
		return nil
	}

	return s.Db.Transaction(func(tx *gorm.DB) error {
		for i := range *ms {
			if err := tx.FirstOrCreate(&(*ms)[i], entity.OperationIdEnumPair{
				OperationId: (*ms)[i].OperationId,
			}).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *OperationIdEnumPairService) List(operationId string) (entities entity.OperationIdEnumPairs, total int64, err error) {
	tx := s.Db.Model(new(entity.OperationIdEnumPair))
	result := tx.Where("operation_id LIKE ?", "%"+operationId+"%").Find(&entities)
	total = result.RowsAffected
	err = result.Error
	return
}

func (s *OperationIdEnumPairService) ListCopy(copyToEntities any, operationId string) (entities entity.OperationIdEnumPairs, total int64, err error) {
	entities, total, err = s.List(operationId)
	if err != nil {
		return
	}
	err = copier.CopyWithOption(copyToEntities, &entities, crud.GlobalCopierOption)
	if err != nil {
		return
	}
	return
}

func (s *OperationIdEnumPairService) InitUnderContextWithCancel(ctx context.Context, db *gorm.DB) (*OperationIdEnumPairService, context.CancelFunc) {
	db, cancel := dbcontext.SectionUnderContextWithCancel(ctx, db)
	return s.Init(ctx, db), cancel
}

func (s *OperationIdEnumPairService) Init(ctx context.Context, db *gorm.DB) *OperationIdEnumPairService {
	s.ctx = ctx
	s.Db = db
	return s
}
