package crud

import (
	"my-app/backend/internal/entity"
	"my-app/backend/internal/interfaces"
	"my-app/backend/pkg/db"
	"strconv"
)

type CRUDOption struct {
	*db.CRUD[*entity.Option]
}

// GetUint16ByOptionName implements interfaces.ICRUDOption.
func (s *CRUDOption) GetUint16ByOptionName(name string) (value uint16, opt *entity.Option, err error) {
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

// GetByOptionName implements interfaces.ICRUDOption.
func (s *CRUDOption) GetByOptionName(name string) (opt *entity.Option, err error) {
	return s.FindOne(func(where func(query any, args ...any)) {
		where(&entity.Option{
			Key: name,
		})
	})
}

func NewCRUDOption(dbs *db.DB) interfaces.ICRUDOption {
	return &CRUDOption{
		CRUD: db.NewCRUD[*entity.Option](dbs),
	}
}
