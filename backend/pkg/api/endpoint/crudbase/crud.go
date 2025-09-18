package crudbase

import (
	"context"
	"strings"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/api/schema"
	"majinyao.cn/my-app/backend/pkg/db"
	"majinyao.cn/my-app/backend/pkg/db/crud"
	"majinyao.cn/my-app/backend/pkg/utils"
)

type Crud[
	T db.EntityIdGetter,
	TItem schema.EntityIdGetter,
	TDetail schema.EntityIdGetter,
	TSave schema.EntityIdGetter,
] struct {
	EntityName          string
	SlugifiedName       string
	SpacedName          string
	LowerCaseSpacedName string
	Db                  *gorm.DB
	CopierOption        copier.Option
	GetCrudService      func(ctx context.Context, tx *gorm.DB) (crud.ICrudService[T], context.CancelFunc)
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
func (c *Crud[T, TItem, TDetail, TSave]) Init(name string, db *gorm.DB, copierOption copier.Option, getCrudService func(ctx context.Context, tx *gorm.DB) (crud.ICrudService[T], context.CancelFunc), schemes ...string) *Crud[T, TItem, TDetail, TSave] {
	c.EntityName = name
	c.SlugifiedName = utils.SlugifyLDCamelCaseVariableName(name)
	c.SpacedName = utils.SpaceLDCamelCaseVariableName(name)
	c.LowerCaseSpacedName = strings.ToLower(c.SpacedName)
	c.Db = db
	c.CopierOption = copierOption
	if getCrudService == nil {
		c.GetCrudService = func(ctx context.Context, tx *gorm.DB) (crud.ICrudService[T], context.CancelFunc) {
			return crud.NewtWithCancelUnderContext[T](ctx, tx)
		}
	} else {
		c.GetCrudService = getCrudService
	}
	c.Schemes = schemes
	return c
}
