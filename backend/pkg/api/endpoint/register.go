package endpoint

import (
	"github.com/danielgtaylor/huma/v2"
)

type Register interface {
	Register(api huma.API) (ops []huma.Operation)
}
