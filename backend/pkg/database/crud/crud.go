package crud

import (
	"my-app/backend/pkg/database/engine"
	"my-app/backend/pkg/database/interfaces"
	"my-app/backend/pkg/database/options"

	"gorm.io/gorm"
)

type Crud[TEntity interfaces.IEntity] struct {
	*engine.Engine[interfaces.IEntity]
}

// GetById implements interfaces.ICrud
func (c *Crud[TEntity]) GetById(id int64) (entity TEntity) {
	panic("unimplemented")
}

// Query implements interfaces.ICrud
func (c *Crud[TEntity]) Query(criteria *options.OCriteria, conditions ...interfaces.QueryCondition) (entities []TEntity, err error) {
	criteria = options.NewOCriteria(criteria)

	err = c.Transaction(func(tx *gorm.DB) error {
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

		for _, condition := range conditions {
			query, args := condition()
			tx = tx.Where(query, args...)
		}

		return tx.Find(&entities).Error
	})

	return
}

func NewCrud[TEntity interfaces.IEntity](engine *engine.Engine[interfaces.IEntity], entity TEntity) interfaces.ICrud[TEntity] {
	return &Crud[TEntity]{
		Engine: engine,
	}
}
