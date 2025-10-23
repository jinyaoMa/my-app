package crudbase

import (
	"context"
	"strings"

	"github.com/danielgtaylor/huma/v2"
	"majinyao.cn/my-app/backend/pkg/db/crud"
	"majinyao.cn/my-app/backend/pkg/db/model"
	"majinyao.cn/my-app/backend/pkg/utils"
)

type Crud[T model.IdGetter, TItem any, TSave any, TDetail any, TSearchItem any] struct {
	EntityName          string
	SlugifiedName       string
	SpacedName          string
	LowerCaseSpacedName string
	GetCrudService      func(ctx context.Context) (crud.ICrud[T], context.CancelFunc)
	Schemes             []string
}

func (c *Crud[T, TItem, TSave, TDetail, TSearchItem]) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		c.RegisterDelete(api),
		c.RegisterDetail(api),
		c.RegisterQuery(api),
		c.RegisterSave(api),
		c.RegisterSearch(api),
		c.RegisterRecycle(api, "deleted_at"),
		c.RegisterRecover(api, "deleted_at"),
		c.RegisterBatchDelete(api),
		c.RegisterBatchRecover(api, "deleted_at"),
	)
	return
}

// name could be entity/module name in camelcase, w/o special characters, should use rule for naming variable
func (c *Crud[T, TItem, TSave, TDetail, TSearchItem]) Init(name string, getCrudService func(ctx context.Context) (crud.ICrud[T], context.CancelFunc), schemes ...string) *Crud[T, TItem, TSave, TDetail, TSearchItem] {
	c.EntityName = name
	c.SlugifiedName = utils.SlugifyLDCamelCaseVariableName(name)
	c.SpacedName = utils.SpaceLDCamelCaseVariableName(name)
	c.LowerCaseSpacedName = strings.ToLower(c.SpacedName)
	c.GetCrudService = getCrudService
	c.Schemes = schemes
	return c
}
