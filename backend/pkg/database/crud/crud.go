package crud

import (
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/database/interfaces"
	"my-app/backend/pkg/database/options"

	"gorm.io/gorm"
)

type Crud[TEntity interfaces.IEntity] struct {
	db *database.Database
}

// All implements interfaces.ICrud
func (c *Crud[TEntity]) All() (entities []TEntity) {
	c.db.Find(&entities)
	return
}

// GetById implements interfaces.ICrud
func (c *Crud[TEntity]) GetById(id int64) (entity TEntity) {
	c.db.Limit(1).Find(&entity, id)
	return
}

// Query implements interfaces.ICrud
func (c *Crud[TEntity]) Query(criteria *options.OCriteria, condition interfaces.QueryCondition) (entities []TEntity, err error) {
	criteria = options.NewOCriteria(criteria)

	err = c.db.Transaction(func(tx *gorm.DB) error {
		tx = tx.Limit(criteria.Size).Offset(criteria.Offset())

		if len(criteria.Fields) > 0 {
			tx = tx.Select(criteria.Fields)
		}

		for _, sort := range criteria.Sorts {
			switch sort.Order {
			case options.OrdAscending:
				tx = tx.Order(sort.Column + " asc")
			case options.OrdDescending:
				tx = tx.Order(sort.Column + " desc")
			}
		}

		condition(func(query any, args ...any) {
			tx = tx.Where(query, args...)
		})

		return tx.Find(&entities).Error
	})

	return
}

func NewCrud[TEntity interfaces.IEntity](database *database.Database, entity TEntity) interfaces.ICrud[TEntity] {
	return &Crud[TEntity]{
		db: database,
	}
}
