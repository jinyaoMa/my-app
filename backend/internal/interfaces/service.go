package interfaces

import "my-app/backend/internal/entity"

type IServiceOption interface {
	GetOptions() (opts []*entity.Option)
	SaveOption(opt *entity.Option) (ok bool)
}
