package db

import (
	"my-app/pkg/db/interfaces"
	"my-app/pkg/db/param"
)

type CRUD[TEntity interfaces.IEntity] struct {
	db *DB
}

// Save implements interfaces.ICRUD
func (c *CRUD[TEntity]) Delete(id int64) (affected int64, err error) {
	result := c.db.Delete(new(TEntity), id)
	affected = result.RowsAffected
	err = result.Error
	return
}

// Save implements interfaces.ICRUD
func (c *CRUD[TEntity]) Save(entity TEntity) (affected int64, err error) {
	result := c.db.Save(&entity)
	affected = result.RowsAffected
	err = result.Error
	return
}

// SaveAll implements interfaces.ICRUD
func (c *CRUD[TEntity]) SaveAll(entities []TEntity) (affected int64, err error) {
	result := c.db.Save(&entities)
	affected = result.RowsAffected
	err = result.Error
	return
}

// FindOne implements interfaces.ICRUD
func (c *CRUD[TEntity]) FindOne(condition param.QueryCondition) (entity TEntity, err error) {
	tx := c.db.Limit(1)
	condition(func(query any, args ...any) {
		tx = tx.Where(query, args...)
	})
	err = tx.Find(&entity).Error
	return
}

// All implements interfaces.ICRUD
func (c *CRUD[TEntity]) All() (entities []TEntity, err error) {
	err = c.db.Find(&entities).Error
	return
}

// GetById implements interfaces.ICRUD
func (c *CRUD[TEntity]) GetById(id int64) (entity TEntity, err error) {
	err = c.db.First(&entity, id).Error
	return
}

// Query implements interfaces.ICRUD
func (c *CRUD[TEntity]) Query(criteria *param.Criteria, condition param.QueryCondition, includes ...string) (entities []TEntity, err error) {
	criteria = param.NewCriteria(criteria)

	tx := c.db.Limit(criteria.Size).Offset(criteria.Offset())

	for _, include := range includes {
		tx = tx.Preload(include)
	}

	if len(criteria.Fields) > 0 {
		tx = tx.Select(criteria.Fields)
	}

	for _, sort := range criteria.Sorts {
		switch sort.Order {
		case param.OrdAscending:
			tx = tx.Order(sort.Column + " ASC")
		case param.OrdDescending:
			tx = tx.Order(sort.Column + " DESC")
		}
	}

	condition(func(query any, args ...any) {
		tx = tx.Where(query, args...)
	})

	err = tx.Find(&entities).Error
	return
}

func NewCRUD[TEntity interfaces.IEntity](database *DB) *CRUD[TEntity] {
	return &CRUD[TEntity]{
		db: database,
	}
}

func NewICRUD[TEntity interfaces.IEntity](database *DB) interfaces.ICRUD[TEntity] {
	return NewCRUD[TEntity](database)
}
