package engine

import "my-app/backend/pkg/database/interfaces"

func (e *Engine[TEntity]) Query(criteria interfaces.Criteria) (entities []TEntity, err error) {
	err = e.Find(entities)
	return
}
