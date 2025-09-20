package crud

import (
	"context"
	"errors"
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
	"majinyao.cn/my-app/backend/pkg/db/dbcontext"
	"majinyao.cn/my-app/backend/pkg/db/model"
)

type ICrud[T model.IdGetter] interface {
	SetCopierOption(copierOption copier.Option)
	SaveCopy(copyFrom any, choice Choice) (entity T, affected int64, err error)
	Create(entity *T, choices ...Choice) (affected int64, err error)
	CreateCopy(copyFrom any, choices ...Choice) (entity T, affected int64, err error)
	BatchCreate(entities *[]T, choices ...Choice) (affected int64, err error)
	Update(entity *T, choices ...Choice) (affected int64, err error)
	UpdateCopy(copyFrom any, choices ...Choice) (entity T, affected int64, err error)
	Delete(ids ...datatype.Id) (affected int64, err error)
	GetById(id datatype.Id, includes ...string) (entity T, notFound bool, err error)
	ScanById(entity any, id datatype.Id, includes ...string) (notFound bool, err error)
	GetCopyById(copyToEntity any, id datatype.Id, includes ...string) (entity T, notFound bool, err error)
	All(includes ...string) (entities []T, total int64, err error)
	ScanAll(entities any, includes ...string) (total int64, err error)
	CopyAll(copyToEntities any, includes ...string) (entities []T, total int64, err error)
	Query(criteria Criteria) (entities []T, total int64, err error)
	QueryScan(entities any, criteria Criteria) (total int64, err error)
	QueryCopy(copyToEntities any, criteria Criteria) (entities []T, total int64, err error)
	QueryWithCondition(criteria Criteria, condition func(tx *gorm.DB) (*gorm.DB, error)) (entities []T, total int64, err error)
	QueryScanWithCondition(entities any, criteria Criteria, condition func(tx *gorm.DB) (*gorm.DB, error)) (total int64, err error)
	QueryCopyWithCondition(copyToEntities any, criteria Criteria, condition func(tx *gorm.DB) (*gorm.DB, error)) (entities []T, total int64, err error)
	FindOne(condition func(tx *gorm.DB) (*gorm.DB, error), includes ...string) (entity T, notFound bool, err error)
	ScanOne(entity any, condition func(tx *gorm.DB) (*gorm.DB, error), includes ...string) (notFound bool, err error)
}

func New[T model.IdGetter](db *gorm.DB) ICrud[T] {
	return new(Crud[T]).Init(db)
}

func NewtWithCancelUnderContext[T model.IdGetter](ctx context.Context, db *gorm.DB) (ICrud[T], context.CancelFunc) {
	return new(Crud[T]).InitWithCancelUnderContext(ctx, db)
}

func NewWithTimeoutUnderContext[T model.IdGetter](ctx context.Context, db *gorm.DB, timeout time.Duration) (ICrud[T], context.CancelFunc) {
	return new(Crud[T]).InitWithTimeoutUnderContext(ctx, db, timeout)
}

type Crud[T model.IdGetter] struct {
	Db           *gorm.DB
	CopierOption copier.Option
}

func (c *Crud[T]) SetCopierOption(copierOption copier.Option) {
	c.CopierOption = copierOption
}

func (c *Crud[T]) SaveCopy(copyFrom any, choice Choice) (entity T, affected int64, err error) {
	err = copier.CopyWithOption(&entity, copyFrom, c.CopierOption)
	if err != nil {
		return
	}

	if entity.IsTransient() {
		affected, err = c.Create(&entity)
	} else {
		affected, err = c.Update(&entity, choice)
	}
	return
}

func (c *Crud[T]) Create(entity *T, choices ...Choice) (affected int64, err error) {
	tx := c.Db
	for _, choice := range choices {
		tx = choice.Apply(tx)
	}

	result := tx.Create(entity)
	affected = result.RowsAffected
	err = result.Error
	return
}

func (c *Crud[T]) CreateCopy(copyFrom any, choices ...Choice) (entity T, affected int64, err error) {
	err = copier.CopyWithOption(&entity, copyFrom, c.CopierOption)
	if err != nil {
		return
	}
	affected, err = c.Create(&entity, choices...)
	return
}

func (c *Crud[T]) BatchCreate(entities *[]T, choices ...Choice) (affected int64, err error) {
	tx := c.Db
	for _, choice := range choices {
		tx = choice.Apply(tx)
	}

	result := tx.CreateInBatches(entities, len(*entities))
	affected = result.RowsAffected
	err = result.Error
	return
}

func (c *Crud[T]) Update(entity *T, choices ...Choice) (affected int64, err error) {
	tx := c.Db.Model(entity)
	for _, choice := range choices {
		tx = choice.Apply(tx)
	}

	result := tx.Updates(entity)
	affected = result.RowsAffected
	err = result.Error
	return
}

func (c *Crud[T]) UpdateCopy(copyFrom any, choices ...Choice) (entity T, affected int64, err error) {
	err = copier.CopyWithOption(&entity, copyFrom, c.CopierOption)
	if err != nil {
		return
	}
	affected, err = c.Update(&entity, choices...)
	return
}

func (c *Crud[T]) Delete(ids ...datatype.Id) (affected int64, err error) {
	result := c.Db.Delete(&[]T{}, ids)
	affected = result.RowsAffected
	err = result.Error
	return
}

func (c *Crud[T]) GetById(id datatype.Id, includes ...string) (entity T, notFound bool, err error) {
	tx := c.Db
	for _, include := range includes {
		tx = tx.Preload(include)
	}

	res := tx.Take(&entity, id)
	err = res.Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		notFound = true
		err = nil
		return
	}
	return
}

func (c *Crud[T]) ScanById(entity any, id datatype.Id, includes ...string) (notFound bool, err error) {
	tx := c.Db.Model(new(T))
	for _, include := range includes {
		tx = tx.Preload(include)
	}

	res := tx.Where([]datatype.Id{id}).
		Limit(1).
		Scan(entity)

	err = res.Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		notFound = true
		err = nil
		return
	}
	return
}

func (c *Crud[T]) GetCopyById(copyToEntity any, id datatype.Id, includes ...string) (entity T, notFound bool, err error) {
	entity, notFound, err = c.GetById(id, includes...)
	if err == nil && !notFound {
		err = copier.CopyWithOption(copyToEntity, &entity, c.CopierOption)
		if err != nil {
			return
		}
	}
	return
}

func (c *Crud[T]) All(includes ...string) (entities []T, total int64, err error) {
	tx := c.Db.Model(new(T))
	for _, include := range includes {
		tx = tx.Preload(include)
	}

	res := tx.Find(&entities)
	total = res.RowsAffected
	err = res.Error
	return
}

func (c *Crud[T]) ScanAll(entities any, includes ...string) (total int64, err error) {
	tx := c.Db.Model(new(T))
	for _, include := range includes {
		tx = tx.Preload(include)
	}

	res := tx.Scan(entities)
	total = res.RowsAffected
	err = res.Error
	return
}

func (c *Crud[T]) CopyAll(copyToEntities any, includes ...string) (entities []T, total int64, err error) {
	entities, total, err = c.All(includes...)
	if err != nil {
		return
	}
	err = copier.CopyWithOption(copyToEntities, &entities, c.CopierOption)
	if err != nil {
		return
	}
	return
}

func (c *Crud[T]) Query(criteria Criteria) (entities []T, total int64, err error) {
	return c.QueryWithCondition(criteria, nil)
}

func (c *Crud[T]) QueryScan(entities any, criteria Criteria) (total int64, err error) {
	return c.QueryScanWithCondition(entities, criteria, nil)
}

func (c *Crud[T]) QueryCopy(copyToEntities any, criteria Criteria) (entities []T, total int64, err error) {
	entities, total, err = c.Query(criteria)
	if err != nil {
		return
	}
	err = copier.CopyWithOption(copyToEntities, &entities, c.CopierOption)
	if err != nil {
		return
	}
	return
}

func (c *Crud[T]) QueryWithCondition(criteria Criteria, condition func(tx *gorm.DB) (*gorm.DB, error)) (entities []T, total int64, err error) {
	tx := c.Db.Model(new(T))

	if condition != nil {
		tx, err = condition(tx)
		if err != nil {
			return
		}
	}

	err = tx.Count(&total).Error
	if err != nil {
		return
	}

	tx = criteria.Apply(tx)

	err = tx.Find(&entities).Error
	return
}

func (c *Crud[T]) QueryScanWithCondition(entities any, criteria Criteria, condition func(tx *gorm.DB) (*gorm.DB, error)) (total int64, err error) {
	tx := c.Db.Model(new(T))

	if condition != nil {
		tx, err = condition(tx)
		if err != nil {
			return
		}
	}

	err = tx.Count(&total).Error
	if err != nil {
		return
	}

	tx = criteria.Apply(tx)

	err = tx.Scan(entities).Error
	return
}

func (c *Crud[T]) QueryCopyWithCondition(copyToEntities any, criteria Criteria, condition func(tx *gorm.DB) (*gorm.DB, error)) (entities []T, total int64, err error) {
	entities, total, err = c.QueryWithCondition(criteria, condition)
	if err != nil {
		return
	}
	err = copier.CopyWithOption(copyToEntities, &entities, c.CopierOption)
	if err != nil {
		return
	}
	return
}

func (c *Crud[T]) FindOne(condition func(tx *gorm.DB) (*gorm.DB, error), includes ...string) (entity T, notFound bool, err error) {
	tx := c.Db
	for _, include := range includes {
		tx = tx.Preload(include)
	}

	if condition != nil {
		tx, err = condition(tx)
		if err != nil {
			return
		}
	}

	res := tx.Take(&entity)
	err = res.Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		notFound = true
		err = nil
		return
	}
	return
}

func (c *Crud[T]) ScanOne(entity any, condition func(tx *gorm.DB) (*gorm.DB, error), includes ...string) (notFound bool, err error) {
	tx := c.Db.Model(new(T))
	for _, include := range includes {
		tx = tx.Preload(include)
	}

	if condition != nil {
		tx, err = condition(tx)
		if err != nil {
			return
		}
	}

	res := tx.Limit(1).Scan(entity)
	err = res.Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		notFound = true
		err = nil
		return
	}
	return
}

func (c *Crud[T]) Init(db *gorm.DB) *Crud[T] {
	c.Db = db
	c.CopierOption = DefaultCopierOption
	return c
}

func (c *Crud[T]) InitWithCancelUnderContext(ctx context.Context, db *gorm.DB) (*Crud[T], context.CancelFunc) {
	db, cancel := dbcontext.SectionUnderContextWithCancel(ctx, db)
	return c.Init(db), cancel
}

func (c *Crud[T]) InitWithTimeoutUnderContext(ctx context.Context, db *gorm.DB, timeout time.Duration) (*Crud[T], context.CancelFunc) {
	db, cancel := dbcontext.SectionUnderContextWithTimeout(ctx, db, timeout)
	return c.Init(db), cancel
}
