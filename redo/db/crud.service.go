package db

import (
	"my-app/backend/pkg/database/interfaces"
	"my-app/backend/pkg/database/vmodel"
)

type CrudService[TEntity interfaces.IEntity] struct {
	db *DB
}

func NewCrudService[TEntity interfaces.IEntity](database *DB) interfaces.ICrudService[TEntity] {
	return &CrudService[TEntity]{
		db: database,
	}
}

// Save implements interfaces.ICrudService
func (c *CrudService[TEntity]) Delete(id int64) (affected int64, err error) {
	result := c.db.Delete(new(TEntity), id)
	affected = result.RowsAffected
	err = result.Error
	return
}

// Save implements interfaces.ICrudService
func (c *CrudService[TEntity]) Save(entity TEntity) (affected int64, err error) {
	result := c.db.Save(&entity)
	affected = result.RowsAffected
	err = result.Error
	return
}

// SaveAll implements interfaces.ICrudService
func (c *CrudService[TEntity]) SaveAll(entities []TEntity) (affected int64, err error) {
	result := c.db.Save(&entities)
	affected = result.RowsAffected
	err = result.Error
	return
}

// FindOne implements interfaces.ICrudService
func (c *CrudService[TEntity]) FindOne(condition vmodel.QueryCondition) (entity TEntity, err error) {
	tx := c.db.Limit(1)
	condition(func(query any, args ...any) {
		tx = tx.Where(query, args...)
	})
	err = tx.Find(&entity).Error
	return
}

// All implements interfaces.ICrudService
func (c *CrudService[TEntity]) All() (entities []TEntity, err error) {
	err = c.db.Find(&entities).Error
	return
}

// GetById implements interfaces.ICrudService
func (c *CrudService[TEntity]) GetById(id int64) (entity TEntity, err error) {
	err = c.db.First(&entity, id).Error
	return
}

// Query implements interfaces.ICrudService
func (c *CrudService[TEntity]) Query(criteria *vmodel.Criteria, condition vmodel.QueryCondition, includes ...string) (entities []TEntity, err error) {
	criteria = vmodel.NewCriteria(criteria)

	tx := c.db.Limit(criteria.Size).Offset(criteria.Offset())

	for _, include := range includes {
		tx = tx.Preload(include)
	}

	if len(criteria.Fields) > 0 {
		tx = tx.Select(criteria.Fields)
	}

	for _, sort := range criteria.Sorts {
		switch sort.Order {
		case vmodel.OrdAscending:
			tx = tx.Order(sort.Column + " ASC")
		case vmodel.OrdDescending:
			tx = tx.Order(sort.Column + " DESC")
		}
	}

	condition(func(query any, args ...any) {
		tx = tx.Where(query, args...)
	})

	err = tx.Find(&entities).Error
	return
}
