package service

import (
	"my-app/backend/internal/crud"
	"my-app/backend/internal/entity"
	"my-app/backend/internal/interfaces"
	"my-app/backend/pkg/db"

	"github.com/devfeel/mapper"
)

type Option struct {
	crud   interfaces.ICRUDOption
	mapper mapper.IMapper
}

// GetOptions implements interfaces.IServiceOption.
func (o *Option) GetOptions() (opts []*entity.Option) {
	opts, _ = o.crud.All()
	return
}

// SaveOption implements interfaces.IServiceOption.
func (o *Option) SaveOption(opt *entity.Option) (ok bool) {
	_, err := o.crud.Save(opt)
	return err == nil
}

func NewOption(dbs *db.DB, m mapper.IMapper) *Option {
	return &Option{
		crud:   crud.NewOption(dbs),
		mapper: m,
	}
}

func NewIOption(dbs *db.DB, m mapper.IMapper) interfaces.IServiceOption {
	return NewOption(dbs, m)
}