package service

import (
	i "my-app/backend/internal/interfaces"
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/database/entity"
	"my-app/backend/pkg/database/interfaces"
)

type OptionService struct {
	interfaces.ICrudService[*entity.Option]
	db *database.Database
}

// GetByOptionName implements interfaces.IOptionService.
func (s *OptionService) GetByOptionName(name string) (value string, err error) {
	var opt *entity.Option
	opt, err = s.FindOne(func(where func(query any, args ...any)) {
		where(&entity.Option{
			Key: name,
		})
	})
	if err != nil {
		return
	}
	return opt.Value, nil
}

func NewOptionService(db *database.Database) i.IOptionService {
	return &OptionService{
		ICrudService: database.NewCrudService[*entity.Option](db),
		db:           db,
	}
}
