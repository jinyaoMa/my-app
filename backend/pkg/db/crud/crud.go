package crud

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/db"
)

type ICrudService[T db.EntityIdGetter] interface {
	Create(entity *T, choices ...Choice) (affected int64, err error)
	BatchCreate(entities *[]T, choices ...Choice) (affected int64, err error)
	Update(entity *T, choices ...Choice) (affected int64, err error)
	Delete(ids ...int64) (affected int64, err error)
	GetById(id int64, includes ...string) (entity T, notFound bool, err error)
	ScanById(entity any, id int64, includes ...string) (notFound bool, err error)
	All(includes ...string) (entities []T, total int64, err error)
	ScanAll(entities any, includes ...string) (total int64, err error)
	Query(criteria Criteria) (entities []T, total int64, err error)
	ScanQuery(entities any, criteria Criteria) (total int64, err error)
	QueryWithCondition(criteria Criteria, condition func(tx *gorm.DB) (*gorm.DB, error)) (entities []T, total int64, err error)
	ScanQueryWithCondition(entities any, criteria Criteria, condition func(tx *gorm.DB) (*gorm.DB, error)) (total int64, err error)
}

func New[T db.EntityIdGetter](tx *gorm.DB) ICrudService[T] {
	return new(Crud[T]).Init(tx)
}

func NewtWithCancelUnderContext[T db.EntityIdGetter](ctx context.Context, tx *gorm.DB) (ICrudService[T], context.CancelFunc) {
	return new(Crud[T]).InitWithCancelUnderContext(ctx, tx)
}

func NewWithTimeoutUnderContext[T db.EntityIdGetter](ctx context.Context, tx *gorm.DB, timeout time.Duration) (ICrudService[T], context.CancelFunc) {
	return new(Crud[T]).InitWithTimeoutUnderContext(ctx, tx, timeout)
}

type Crud[T db.EntityIdGetter] struct {
	Db *gorm.DB
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

func (c *Crud[T]) Delete(ids ...int64) (affected int64, err error) {
	result := c.Db.Delete(&[]T{}, ids)
	affected = result.RowsAffected
	err = result.Error
	return
}

func (c *Crud[T]) GetById(id int64, includes ...string) (entity T, notFound bool, err error) {
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

func (c *Crud[T]) ScanById(entity any, id int64, includes ...string) (notFound bool, err error) {
	tx := c.Db.Model(new(T))
	for _, include := range includes {
		tx = tx.Preload(include)
	}

	res := tx.Where([]int64{id}).
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

func (c *Crud[T]) Query(criteria Criteria) (entities []T, total int64, err error) {
	return c.QueryWithCondition(criteria, nil)
}

func (c *Crud[T]) ScanQuery(entities any, criteria Criteria) (total int64, err error) {
	return c.ScanQueryWithCondition(entities, criteria, nil)
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

func (c *Crud[T]) ScanQueryWithCondition(entities any, criteria Criteria, condition func(tx *gorm.DB) (*gorm.DB, error)) (total int64, err error) {
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

func (c *Crud[T]) Init(tx *gorm.DB) *Crud[T] {
	c.Db = tx
	return c
}

func (c *Crud[T]) InitWithCancelUnderContext(ctx context.Context, tx *gorm.DB) (*Crud[T], context.CancelFunc) {
	tx, cancel := db.SectionUnderContextWithCancel(ctx, tx)
	return c.Init(tx), cancel
}

func (c *Crud[T]) InitWithTimeoutUnderContext(ctx context.Context, tx *gorm.DB, timeout time.Duration) (*Crud[T], context.CancelFunc) {
	tx, cancel := db.SectionUnderContextWithTimeout(ctx, tx, timeout)
	return c.Init(tx), cancel
}
