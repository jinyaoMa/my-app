package crud

import (
	"errors"
	"slices"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
	"majinyao.cn/my-app/backend/pkg/db/model"
)

type ICrud[T model.IdGetter] interface {
	SaveCopy(copyFromView any, copyToView any, choice Choice) (newEntity T, affected int64, err error)
	Create(entity *T, choices ...Choice) (affected int64, err error)
	CreateCopy(copyFromView any, copyToView any, choices ...Choice) (entity T, affected int64, err error)
	BatchCreate(entities *[]T, choices ...Choice) (affected int64, err error)
	Update(entity T, choices ...Choice) (newEntity T, affected int64, err error)
	UpdateCopy(copyFromView any, copyToView any, choices ...Choice) (newEntity T, affected int64, err error)
	Delete(id datatype.Id) (entity T, affected int64, err error)
	BatchDelete(ids []datatype.Id) (affected int64, err error)
	Recover(id datatype.Id, fieldName string) (entity T, affected int64, err error)
	BatchRecover(ids []datatype.Id, fieldName string) (affected int64, err error)
	GetById(id datatype.Id, includes Includes) (entity T, notFound bool, err error)
	ScanById(entity any, id datatype.Id, includes Includes) (notFound bool, err error)
	GetCopyById(copyToView any, id datatype.Id, includes Includes) (entity T, notFound bool, err error)
	All(includes Includes) (entities []T, total int64, err error)
	ScanAll(entities any, includes Includes) (total int64, err error)
	CopyAll(copyToList any, includes Includes) (entities []T, total int64, err error)
	Query(criteria Criteria, conds ...func(tx *gorm.DB) (*gorm.DB, error)) (entities []T, total int64, err error)
	QueryScan(entities any, criteria Criteria, conds ...func(tx *gorm.DB) (*gorm.DB, error)) (total int64, err error)
	QueryCopy(copyToList any, criteria Criteria, conds ...func(tx *gorm.DB) (*gorm.DB, error)) (entities []T, total int64, err error)
	FindOne(cond func(tx *gorm.DB) (*gorm.DB, error), includes Includes) (entity T, notFound bool, err error)
	ScanOne(entity any, cond func(tx *gorm.DB) (*gorm.DB, error), includes Includes) (notFound bool, err error)
	CopyOne(copyToView any, cond func(tx *gorm.DB) (*gorm.DB, error), includes Includes) (entity T, notFound bool, err error)
	Search(keyword string, fields []string, limit int, exact bool, includes Includes, conds ...func(tx *gorm.DB) (*gorm.DB, error)) (entities []T, total int64, err error)
	SearchScan(entities any, keyword string, fields []string, limit int, exact bool, includes Includes, conds ...func(tx *gorm.DB) (*gorm.DB, error)) (total int64, err error)
	SearchCopy(copyToList any, keyword string, fields []string, limit int, exact bool, includes Includes, conds ...func(tx *gorm.DB) (*gorm.DB, error)) (entities []T, total int64, err error)
	Count(cond func(tx *gorm.DB) (*gorm.DB, error)) (total int64, err error)
}

func New[T model.IdGetter](options Options, db *gorm.DB) ICrud[T] {
	return new(Crud[T]).Init(options, db)
}

type Crud[T model.IdGetter] struct {
	Options Options
	Db      *gorm.DB
}

func (c *Crud[T]) SaveCopy(copyFromView any, copyToView any, choice Choice) (newEntity T, affected int64, err error) {
	err = copier.CopyWithOption(&newEntity, copyFromView, GlobalCopierOption)
	if err != nil {
		return
	}

	if newEntity.IsTransient() {
		affected, err = c.Create(&newEntity)
	} else {
		newEntity, affected, err = c.Update(newEntity, choice)
	}
	if affected == 1 {
		err = copier.CopyWithOption(copyToView, &newEntity, GlobalCopierOption)
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

func (c *Crud[T]) CreateCopy(copyFromView any, copyToView any, choices ...Choice) (entity T, affected int64, err error) {
	err = copier.CopyWithOption(&entity, copyFromView, GlobalCopierOption)
	if err != nil {
		return
	}
	affected, err = c.Create(&entity, choices...)
	if affected == 1 {
		err = copier.CopyWithOption(copyToView, &entity, GlobalCopierOption)
	}
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

func (c *Crud[T]) Update(entity T, choices ...Choice) (newEntity T, affected int64, err error) {
	err = c.Db.Transaction(func(tx *gorm.DB) error {
		result := tx.Take(&newEntity, entity.GetId())
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return result.Error
		}

		tx = tx.Model(&newEntity)
		for _, choice := range choices {
			tx = choice.Apply(tx)
		}

		result = tx.Updates(entity)
		affected = result.RowsAffected
		return result.Error
	})
	return
}

func (c *Crud[T]) UpdateCopy(copyFromView any, copyToView any, choices ...Choice) (newEntity T, affected int64, err error) {
	err = copier.CopyWithOption(&newEntity, copyFromView, GlobalCopierOption)
	if err != nil {
		return
	}
	newEntity, affected, err = c.Update(newEntity, choices...)
	if affected == 1 {
		err = copier.CopyWithOption(copyToView, &newEntity, GlobalCopierOption)
	}
	return
}

func (c *Crud[T]) Delete(id datatype.Id) (entity T, affected int64, err error) {
	err = c.Db.Transaction(func(tx *gorm.DB) error {
		result := tx.Take(&entity, id)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return result.Error
		}

		result = tx.Delete(&entity)
		affected = result.RowsAffected
		return result.Error
	})
	return
}

func (c *Crud[T]) BatchDelete(ids []datatype.Id) (affected int64, err error) {
	if len(ids) == 0 {
		return
	}
	result := c.Db.Delete(&[]T{}, ids)
	affected = result.RowsAffected
	err = result.Error
	return
}

func (c *Crud[T]) Recover(id datatype.Id, fieldName string) (entity T, affected int64, err error) {
	if !regField.MatchString(fieldName) {
		return
	}

	err = c.Db.Transaction(func(tx *gorm.DB) error {
		result := tx.Unscoped().Where(fieldName+" IS NOT NULL").Take(&entity, id)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil
			}
			return result.Error
		}

		result = tx.Unscoped().Model(&entity).Update(fieldName, nil)
		affected = result.RowsAffected
		return result.Error
	})
	return
}

func (c *Crud[T]) BatchRecover(ids []datatype.Id, fieldName string) (affected int64, err error) {
	if len(ids) == 0 || !regField.MatchString(fieldName) {
		return
	}

	var t T
	result := c.Db.Unscoped().Model(&t).
		Where(fieldName+" IS NOT NULL").
		Where(t.GetIdFieldName()+" IN ?", ids).
		Update(fieldName, nil)
	affected = result.RowsAffected
	err = result.Error
	return
}

func (c *Crud[T]) GetById(id datatype.Id, includes Includes) (entity T, notFound bool, err error) {
	tx := includes.Apply(c.Db,
		c.Options.CriteriaMaxSize,
		c.Options.CriteriaIncludesMaxLength,
		c.Options.CriteriaSortsMaxLength)

	result := tx.Take(&entity, id)
	err = result.Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		notFound = true
		err = nil
		return
	}
	return
}

func (c *Crud[T]) ScanById(entity any, id datatype.Id, includes Includes) (notFound bool, err error) {
	tx := includes.Apply(c.Db.Model(new(T)),
		c.Options.CriteriaMaxSize,
		c.Options.CriteriaIncludesMaxLength,
		c.Options.CriteriaSortsMaxLength)

	result := tx.Where([]datatype.Id{id}).
		Limit(1).
		Scan(entity)

	err = result.Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		notFound = true
		err = nil
		return
	}
	return
}

func (c *Crud[T]) GetCopyById(copyToView any, id datatype.Id, includes Includes) (entity T, notFound bool, err error) {
	entity, notFound, err = c.GetById(id, includes)
	if err == nil && !notFound {
		err = copier.CopyWithOption(copyToView, &entity, GlobalCopierOption)
		if err != nil {
			return
		}
	}
	return
}

func (c *Crud[T]) All(includes Includes) (entities []T, total int64, err error) {
	tx := includes.Apply(c.Db.Model(new(T)),
		c.Options.CriteriaMaxSize,
		c.Options.CriteriaIncludesMaxLength,
		c.Options.CriteriaSortsMaxLength)

	result := tx.Find(&entities)
	total = result.RowsAffected
	err = result.Error
	return
}

func (c *Crud[T]) ScanAll(entities any, includes Includes) (total int64, err error) {
	tx := includes.Apply(c.Db.Model(new(T)),
		c.Options.CriteriaMaxSize,
		c.Options.CriteriaIncludesMaxLength,
		c.Options.CriteriaSortsMaxLength)

	result := tx.Scan(entities)
	total = result.RowsAffected
	err = result.Error
	return
}

func (c *Crud[T]) CopyAll(copyToList any, includes Includes) (entities []T, total int64, err error) {
	entities, total, err = c.All(includes)
	if err != nil {
		return
	}
	err = copier.CopyWithOption(copyToList, &entities, GlobalCopierOption)
	if err != nil {
		return
	}
	return
}

func (c *Crud[T]) Query(criteria Criteria, conds ...func(tx *gorm.DB) (*gorm.DB, error)) (entities []T, total int64, err error) {
	tx := c.Db.Model(new(T))

	for _, cond := range conds {
		if cond != nil {
			tx, err = cond(tx)
			if err != nil {
				return
			}
		}
	}

	tx, total, err = criteria.Apply(tx, c.Options)
	if err != nil || total == 0 {
		return
	}

	err = tx.Find(&entities).Error
	return
}

func (c *Crud[T]) QueryScan(entities any, criteria Criteria, conds ...func(tx *gorm.DB) (*gorm.DB, error)) (total int64, err error) {
	tx := c.Db.Model(new(T))

	for _, cond := range conds {
		if cond != nil {
			tx, err = cond(tx)
			if err != nil {
				return
			}
		}
	}

	tx, total, err = criteria.Apply(tx, c.Options)
	if err != nil || total == 0 {
		return
	}

	err = tx.Scan(entities).Error
	return
}

func (c *Crud[T]) QueryCopy(copyToList any, criteria Criteria, conds ...func(tx *gorm.DB) (*gorm.DB, error)) (entities []T, total int64, err error) {
	entities, total, err = c.Query(criteria, conds...)
	if err != nil {
		return
	}
	err = copier.CopyWithOption(copyToList, &entities, GlobalCopierOption)
	if err != nil {
		return
	}
	return
}

func (c *Crud[T]) FindOne(cond func(tx *gorm.DB) (*gorm.DB, error), includes Includes) (entity T, notFound bool, err error) {
	tx := includes.Apply(c.Db,
		c.Options.CriteriaMaxSize,
		c.Options.CriteriaIncludesMaxLength,
		c.Options.CriteriaSortsMaxLength)

	if cond != nil {
		tx, err = cond(tx)
		if err != nil {
			return
		}
	}

	result := tx.Take(&entity)
	err = result.Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		notFound = true
		err = nil
		return
	}
	return
}

func (c *Crud[T]) ScanOne(entity any, cond func(tx *gorm.DB) (*gorm.DB, error), includes Includes) (notFound bool, err error) {
	tx := includes.Apply(c.Db.Model(new(T)),
		c.Options.CriteriaMaxSize,
		c.Options.CriteriaIncludesMaxLength,
		c.Options.CriteriaSortsMaxLength)

	if cond != nil {
		tx, err = cond(tx)
		if err != nil {
			return
		}
	}

	result := tx.Limit(1).Scan(entity)
	err = result.Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		notFound = true
		err = nil
		return
	}
	return
}

func (c *Crud[T]) CopyOne(copyToView any, cond func(tx *gorm.DB) (*gorm.DB, error), includes Includes) (entity T, notFound bool, err error) {
	entity, notFound, err = c.FindOne(cond, includes)
	if err == nil && !notFound {
		err = copier.CopyWithOption(copyToView, &entity, GlobalCopierOption)
		if err != nil {
			return
		}
	}
	return
}

func (c *Crud[T]) Search(keyword string, fields []string, limit int, exact bool, includes Includes, conds ...func(tx *gorm.DB) (*gorm.DB, error)) (entities []T, total int64, err error) {
	tx := c.Db.Model(new(T)).Limit(limit)
	for _, cond := range conds {
		if cond != nil {
			tx, err = cond(tx)
			if err != nil {
				return
			}
		}
	}

	tx = includes.Apply(tx,
		c.Options.CriteriaMaxSize,
		c.Options.CriteriaIncludesMaxLength,
		c.Options.CriteriaSortsMaxLength)

	if keyword != "" {
		if exact {
			for i, field := range slices.DeleteFunc(fields, func(field string) bool {
				return !regField.MatchString(field)
			}) {
				if i == 0 {
					tx = tx.Where(field+" = ?", keyword)
				} else {
					tx = tx.Or(field+" = ?", keyword)
				}
			}
		} else {
			keyword = "%" + FixLikeKeyword(keyword) + "%"
			for i, field := range slices.DeleteFunc(fields, func(field string) bool {
				return !regField.MatchString(field)
			}) {
				if i == 0 {
					tx = tx.Where(field+" LIKE ?", keyword)
				} else {
					tx = tx.Or(field+" LIKE ?", keyword)
				}
			}
		}
	}

	result := tx.Find(&entities)
	total = result.RowsAffected
	err = result.Error
	return
}

func (c *Crud[T]) SearchScan(entities any, keyword string, fields []string, limit int, exact bool, includes Includes, conds ...func(tx *gorm.DB) (*gorm.DB, error)) (total int64, err error) {
	tx := c.Db.Model(new(T)).Limit(limit)
	for _, cond := range conds {
		if cond != nil {
			tx, err = cond(tx)
			if err != nil {
				return
			}
		}
	}

	tx = includes.Apply(tx,
		c.Options.CriteriaMaxSize,
		c.Options.CriteriaIncludesMaxLength,
		c.Options.CriteriaSortsMaxLength)

	if keyword != "" {
		if exact {
			for i, field := range slices.DeleteFunc(fields, func(field string) bool {
				return !regField.MatchString(field)
			}) {
				if i == 0 {
					tx = tx.Where(field+" = ?", keyword)
				} else {
					tx = tx.Or(field+" = ?", keyword)
				}
			}
		} else {
			keyword = "%" + FixLikeKeyword(keyword) + "%"
			for i, field := range slices.DeleteFunc(fields, func(field string) bool {
				return !regField.MatchString(field)
			}) {
				if i == 0 {
					tx = tx.Where(field+" LIKE ?", keyword)
				} else {
					tx = tx.Or(field+" LIKE ?", keyword)
				}
			}
		}
	}

	result := tx.Scan(entities)
	total = result.RowsAffected
	err = result.Error
	return
}

func (c *Crud[T]) SearchCopy(copyToList any, keyword string, fields []string, limit int, exact bool, includes Includes, conds ...func(tx *gorm.DB) (*gorm.DB, error)) (entities []T, total int64, err error) {
	entities, total, err = c.Search(keyword, fields, limit, exact, includes, conds...)
	if err != nil {
		return
	}
	err = copier.CopyWithOption(copyToList, &entities, GlobalCopierOption)
	if err != nil {
		return
	}
	return
}

func (c *Crud[T]) Count(cond func(tx *gorm.DB) (*gorm.DB, error)) (total int64, err error) {
	tx := c.Db.Model(new(T))
	if cond != nil {
		tx, err = cond(tx)
		if err != nil {
			return
		}
	}
	result := tx.Count(&total)
	err = result.Error
	return
}

func (c *Crud[T]) Init(options Options, db *gorm.DB) *Crud[T] {
	c.Options = options
	c.Db = db
	return c
}
