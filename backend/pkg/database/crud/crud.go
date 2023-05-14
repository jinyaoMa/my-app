package crud

import (
	"my-app/backend/pkg/database/engine"
	"my-app/backend/pkg/database/interfaces"
	"my-app/backend/pkg/database/options"

	"xorm.io/xorm"
)

type Crud[TEntity interfaces.IEntity] struct {
	engine *engine.Engine[interfaces.IEntity]
}

// GetById implements interfaces.ICrud
func (c *Crud[TEntity]) GetById(id int64) (entity TEntity) {
	panic("unimplemented")
}

// Query implements interfaces.ICrud
func (c *Crud[TEntity]) Query(criteria *options.OCriteria, conditions ...interfaces.QueryCondition) (entities []TEntity, err error) {
	err = c.engine.HandleSession(func(session *xorm.Session) error {
		criteria = options.NewOCriteria(criteria)

		session = session.Limit(criteria.Size, criteria.Offset())

		if len(criteria.Fields) > 0 {
			session = session.Cols(criteria.Fields...)
		}

		for _, sort := range criteria.Sorts {
			switch sort.Order {
			case options.OrdAscending:
				session = session.Asc(sort.Column)
			case options.OrdDescending:
				session = session.Desc(sort.Column)
			}
		}

		for _, condition := range conditions {
			query, args := condition()
			session = session.Where(query, args...)
		}

		return session.Find(&entities)
	})
	return
}

func NewCrud[TEntity interfaces.IEntity](engine *engine.Engine[interfaces.IEntity], entity TEntity) interfaces.ICrud[TEntity] {
	return &Crud[TEntity]{
		engine: engine,
	}
}
