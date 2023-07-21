package crud

import (
	"my-app/backend/pkg/database"
	iCrud "my-app/backend/pkg/database/crud/interfaces"
	iEntity "my-app/backend/pkg/database/entity/interfaces"
	"my-app/backend/pkg/database/options"
)

type Crud[TEntity iEntity.IEntity] struct {
	db *database.Database
}

func NewCrud[TEntity iEntity.IEntity](database *database.Database) iCrud.ICrud[TEntity] {
	return &Crud[TEntity]{
		db: database,
	}
}

// Save implements interfaces.ICrud
func (c *Crud[TEntity]) Delete(id int64) (affected int64, err error) {
	result := c.db.Delete(new(TEntity), id)
	affected = result.RowsAffected
	err = result.Error
	return
}

// Save implements interfaces.ICrud
func (c *Crud[TEntity]) Save(entity TEntity) (affected int64, err error) {
	result := c.db.Save(&entity)
	affected = result.RowsAffected
	err = result.Error
	return
}

// FindOne implements interfaces.ICrud
func (c *Crud[TEntity]) FindOne(condition iCrud.QueryCondition) (entity TEntity, err error) {
	tx := c.db.Limit(1)
	condition(func(query any, args ...any) {
		tx = tx.Where(query, args...)
	})
	err = tx.Find(&entity).Error
	return
}

// All implements interfaces.ICrud
func (c *Crud[TEntity]) All() (entities []TEntity, err error) {
	err = c.db.Find(&entities).Error
	return
}

// GetById implements interfaces.ICrud
func (c *Crud[TEntity]) GetById(id int64) (entity TEntity, err error) {
	err = c.db.First(&entity, id).Error
	return
}

// Query implements interfaces.ICrud
func (c *Crud[TEntity]) Query(criteria *options.OCriteria, condition iCrud.QueryCondition, includes ...string) (entities []TEntity, err error) {
	criteria = options.NewOCriteria(criteria)

	tx := c.db.Limit(criteria.Size).Offset(criteria.Offset())

	for _, include := range includes {
		tx = tx.Preload(include)
	}

	if len(criteria.Fields) > 0 {
		tx = tx.Select(criteria.Fields)
	}

	for _, sort := range criteria.Sorts {
		switch sort.Order {
		case options.OrdAscending:
			tx = tx.Order(sort.Column + " ASC")
		case options.OrdDescending:
			tx = tx.Order(sort.Column + " DESC")
		}
	}

	condition(func(query any, args ...any) {
		tx = tx.Where(query, args...)
	})

	err = tx.Find(&entities).Error
	return
}
