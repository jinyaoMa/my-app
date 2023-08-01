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
func (s *OptionService) GetUint16ByOptionName(name string) (value uint16, opt *entity.Option, err error) {
	opt, err = s.GetByOptionName(name)
	if err != nil {
		return
	}

	var tmp uint64
	tmp, err = strconv.ParseUint(opt.Value, 10, 16)
	if err != nil {
		return
	}

	return uint16(tmp), opt, nil
}

// GetByOptionName implements interfaces.IOptionService.
func (s *OptionService) GetByOptionName(name string) (opt *entity.Option, err error) {
	return s.FindOne(func(where func(query any, args ...any)) {
		where(&entity.Option{
			Key: name,
		})
	})
}

func NewOptionService(db *database.Database) i.IOptionService {
	return &OptionService{
		ICrudService: database.NewCrudService[*entity.Option](db),
		db:           db,
	}
}
