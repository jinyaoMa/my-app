package service

import (
	i "my-app/backend/internal/interfaces"
	"my-app/backend/pkg/database"
	"my-app/backend/pkg/database/entity"
	"my-app/backend/pkg/database/interfaces"
	"strconv"
)

type OptionService struct {
	interfaces.ICrudService[*entity.Option]
	db *database.Database
}

// GetUint16ByOptionName implements interfaces.IOptionService.
func (s *OptionService) GetUint16ByOptionName(name string) (value uint16, err error) {
	var v string
	v, err = s.GetByOptionName(name)
	if err != nil {
		return
	}

	var tmp uint64
	tmp, err = strconv.ParseUint(v, 10, 16)
	if err != nil {
		return
	}

	return uint16(tmp), nil
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
