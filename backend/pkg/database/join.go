package database

import (
	"my-app/backend/pkg/database/entity"
	"my-app/backend/pkg/database/options"

	"gorm.io/gorm"
)

func join(db *gorm.DB, joins ...options.ODatabaseJoin) error {
	joins = append(joins, []options.ODatabaseJoin{
		{
			Model:     new(entity.User),
			Field:     "AccessableFiles",
			JoinTable: new(entity.UserFile),
		},
		{
			Model:     new(entity.File),
			Field:     "AccessableUsers",
			JoinTable: new(entity.UserFile),
		},
	}...)
	for _, params := range joins {
		if err := db.SetupJoinTable(params.Model, params.Field, params.JoinTable); err != nil {
			return err
		}
	}
	return nil
}
