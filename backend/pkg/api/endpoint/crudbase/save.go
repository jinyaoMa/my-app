package crudbase

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/danielgtaylor/huma/v2"
	"majinyao.cn/my-app/backend/pkg/api/schema"
	"majinyao.cn/my-app/backend/pkg/db/crud"
	"majinyao.cn/my-app/backend/pkg/utils"
)

type SaveInput[TSave any] struct {
	Body        TSave
	Supplements []string `query:"supplements" required:"false" doc:"Supplement/Force Update Fields"`
}

func (c *Crud[T, TItem, TSave, TDetail, TSearchItem]) RegisterSave(api huma.API) (op huma.Operation) {
	op = huma.Operation{
		Method:       http.MethodPut,
		Path:         fmt.Sprintf("/%s/save", c.SlugifiedName),
		Summary:      fmt.Sprintf("Save %s", c.SpacedName),
		Description:  fmt.Sprintf("Save %s, if it is a new entity which is transient or id=0, it will be created, otherwise it will be updated, and only the exposed fields will be updated.", c.LowerCaseSpacedName),
		OperationID:  fmt.Sprintf("%s.put.save", c.SlugifiedName),
		Tags:         []string{c.EntityName},
		Errors:       []int{http.StatusInternalServerError},
		MaxBodyBytes: 4 * 1024,
	}

	for _, scheme := range c.Schemes {
		op.Security = append(op.Security, map[string][]string{
			scheme: {op.OperationID},
		})
	}

	var tsave TSave
	defaultSelects := utils.SliceMap(
		reflect.VisibleFields(reflect.TypeOf(tsave)),
		func(e reflect.StructField) string {
			return e.Name
		},
	)

	handler := func(ctx context.Context, input *SaveInput[TSave]) (output *schema.Response[TSave], err error) {
		service, cancel := c.GetCrudService(ctx)
		defer cancel()

		var save TSave
		_, affected, err := service.SaveCopy(&input.Body, &save, crud.Choice{
			Selects: append(defaultSelects, input.Supplements...),
		})
		if err != nil {
			return nil, err
		}
		if affected < 1 {
			return schema.Fail[TSave](http.StatusNotModified, "zero affected"), nil
		}
		return schema.Succeed(save, affected), nil
	}

	huma.Register(api, op, handler)
	return
}
