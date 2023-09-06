package db

import (
	"my-app/backend/pkg/db/param"
)

type CRUD[TEntity IEntity] struct {
	DB *DB
}

func (c *CRUD[TEntity]) mergeEntity(entity TEntity) TEntity {
	entity.SetIdGenerator(c.DB.config.IdGenerator)
	entity.SetCodeGenerator(c.DB.config.CodeGenerator)
	entity.SetDataCipher(c.DB.config.DataCipher)
	return entity
}

func (c *CRUD[TEntity]) mergeEntities(entities []TEntity) []TEntity {
	for _, entity := range entities {
		c.mergeEntity(entity)
	}
	return entities
}

// Save implements ICRUD
func (c *CRUD[TEntity]) Delete(id int64) (affected int64, err error) {
	result := c.DB.Delete(new(TEntity), id)
	affected = result.RowsAffected
	err = result.Error
	return
}

// Save implements ICRUD
func (c *CRUD[TEntity]) Save(entity TEntity) (affected int64, err error) {
	c.mergeEntity(entity)
	result := c.DB.Save(&entity)
	affected = result.RowsAffected
	err = result.Error
	return
}

// SaveAll implements ICRUD
func (c *CRUD[TEntity]) SaveAll(entities []TEntity) (affected int64, err error) {
	c.mergeEntities(entities)
	result := c.DB.Save(&entities)
	affected = result.RowsAffected
	err = result.Error
	return
}

// FindOne implements ICRUD
func (c *CRUD[TEntity]) FindOne(condition param.QueryCondition) (entity TEntity, err error) {
	tx := c.DB.Limit(1)
	condition(func(query any, args ...any) {
		tx = tx.Where(query, args...)
	})
	err = tx.First(&entity).Error
	if err == nil {
		c.mergeEntity(entity)
	}
	return
}

// All implements ICRUD
func (c *CRUD[TEntity]) All() (entities []TEntity, err error) {
	err = c.DB.Find(&entities).Error
	if err == nil {
		c.mergeEntities(entities)
	}
	return
}

// GetById implements ICRUD
func (c *CRUD[TEntity]) GetById(id int64) (entity TEntity, err error) {
	err = c.DB.First(&entity, id).Error
	if err == nil {
		c.mergeEntity(entity)
	}
	return
}

// Query implements ICRUD
func (c *CRUD[TEntity]) Query(criteria *param.Criteria, condition param.QueryCondition, includes ...string) (entities []TEntity, err error) {
	err = c.BuildQuery(criteria, condition, includes...).Find(&entities).Error
	if err == nil {
		c.mergeEntities(entities)
	}
	return
}

// BuildQuery implements ICRUD
func (c *CRUD[TEntity]) BuildQuery(criteria *param.Criteria, condition param.QueryCondition, includes ...string) (db *DB) {
	criteria = param.NewCriteria(criteria)

	tx := c.DB.Limit(criteria.Size).Offset(criteria.Offset())

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

	for _, filter := range criteria.Filters {
		tx = tx.Where(filter.Condition, filter.Params...)
	}

	condition(func(query any, args ...any) {
		tx = tx.Where(query, args...)
	})

	return &DB{
		config: c.DB.config,
		DB:     tx,
	}
}

func NewCRUD[TEntity IEntity](db *DB) *CRUD[TEntity] {
	return &CRUD[TEntity]{
		DB: db,
	}
}

func NewICRUD[TEntity IEntity](db *DB) ICRUD[TEntity] {
	return NewCRUD[TEntity](db)
}
