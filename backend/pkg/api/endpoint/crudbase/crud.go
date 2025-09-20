package crudbase

import (
	"context"
	"strings"

	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/api/schema"
	"majinyao.cn/my-app/backend/pkg/db/crud"
	"majinyao.cn/my-app/backend/pkg/db/model"
	"majinyao.cn/my-app/backend/pkg/utils"
)

type Crud[
	T model.IdGetter,
	TItem schema.ModelIdGetter,
	TDetail schema.ModelIdGetter,
	TSave schema.ModelIdGetter,
] struct {
	EntityName          string
	SlugifiedName       string
	SpacedName          string
	LowerCaseSpacedName string
	Db                  *gorm.DB
	GetCrudService      func(ctx context.Context, db *gorm.DB) (crud.ICrud[T], context.CancelFunc)
	Schemes             []string
}

func (c *Crud[T, TItem, TDetail, TSave]) Register(api huma.API) (ops []huma.Operation) {
	ops = append(ops,
		c.RegisterDelete(api),
		c.RegisterDetail(api),
		c.RegisterQuery(api),
		c.RegisterSave(api),
	)
	return
}

// name could be entity/module name in camelcase, w/o special characters, should use rule for naming variable
func (c *Crud[T, TItem, TDetail, TSave]) Init(name string, db *gorm.DB, getCrudService func(ctx context.Context, db *gorm.DB) (crud.ICrud[T], context.CancelFunc), schemes ...string) *Crud[T, TItem, TDetail, TSave] {
	c.EntityName = name
	c.SlugifiedName = utils.SlugifyLDCamelCaseVariableName(name)
	c.SpacedName = utils.SpaceLDCamelCaseVariableName(name)
	c.LowerCaseSpacedName = strings.ToLower(c.SpacedName)
	c.Db = db
	if getCrudService == nil {
		c.GetCrudService = func(ctx context.Context, db *gorm.DB) (crud.ICrud[T], context.CancelFunc) {
			return crud.NewtWithCancelUnderContext[T](ctx, db)
		}
	} else {
		c.GetCrudService = getCrudService
	}
	c.Schemes = schemes
	return c
}
