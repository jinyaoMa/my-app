package db

import "gorm.io/gorm/clause"

type CRUD[TEntity IEntity] struct {
	db *DB
}

// All implements ICRUD.
func (crud *CRUD[TEntity]) All() (entities []TEntity, err error) {
	err = crud.db.Find(&entities).Error
	return
}

// BuildQuery implements ICRUD.
func (crud *CRUD[TEntity]) BuildQuery(criteria *QueryCriteria, condition QueryCondition, includes ...string) (db *DB, err error) {
	criteria, err = NewQueryCriteria(criteria)
	if err != nil {
		return nil, err
	}

	tx := crud.db.Limit(criteria.Size).Offset(criteria.Offset())

	for _, include := range includes {
		tx = tx.Preload(include)
	}

	if len(criteria.Fields) > 0 {
		tx = tx.Select(criteria.Fields)
	}

	for _, sort := range criteria.Sorts {
		tx = tx.Order(clause.OrderByColumn{
			Column: clause.Column{Name: sort.Column},
			Desc:   sort.Desc,
		})
	}

	for _, filter := range criteria.Filters {
		tx = tx.Where(filter.Condition, filter.Params...)
	}

	condition(func(query any, args ...any) {
		tx = tx.Where(query, args...)
	})

	return &DB{
		options: crud.db.options,
		DB:      tx,
	}, nil
}

// Delete implements ICRUD.
func (crud *CRUD[TEntity]) Delete(id int64) (affected int64, err error) {
	result := crud.db.Delete(new(TEntity), id)
	affected = result.RowsAffected
	err = result.Error
	return
}

// FindOne implements ICRUD.
func (crud *CRUD[TEntity]) FindOne(condition QueryCondition) (entity TEntity, err error) {
	tx := crud.db.Limit(1)
	condition(func(query any, args ...any) {
		tx = tx.Where(query, args...)
	})
	err = tx.First(&entity).Error
	return
}

// GetById implements ICRUD.
func (crud *CRUD[TEntity]) GetById(id int64) (entity TEntity, err error) {
	err = crud.db.First(&entity, id).Error
	return
}

// Query implements ICRUD.
func (crud *CRUD[TEntity]) Query(criteria *QueryCriteria, condition QueryCondition, includes ...string) (entities []TEntity, err error) {
	tx, err := crud.BuildQuery(criteria, condition, includes...)
	if err != nil {
		return nil, err
	}

	err = tx.Find(&entities).Error
	return
}

// Save implements ICRUD.
func (crud *CRUD[TEntity]) Save(entity TEntity) (affected int64, err error) {
	result := crud.db.Save(&entity)
	affected = result.RowsAffected
	err = result.Error
	return
}

// SaveAll implements ICRUD.
func (crud *CRUD[TEntity]) SaveAll(entities []TEntity) (affected int64, err error) {
	result := crud.db.Save(&entities)
	affected = result.RowsAffected
	err = result.Error
	return
}

func NewCRUD[TEntity IEntity](db *DB) (crud *CRUD[TEntity], iCrud ICRUD[TEntity]) {
	crud = &CRUD[TEntity]{
		db: db,
	}
	return crud, crud
}
