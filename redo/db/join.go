package db

import (
	"my-app/backend/pkg/database/entity"

	"gorm.io/gorm"
)

func join(db *gorm.DB, joins ...OptionJoin) error {
	joins = append(joins, []OptionJoin{
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
