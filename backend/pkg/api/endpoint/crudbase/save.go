package crudbase

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jinzhu/copier"
	"majinyao.cn/my-app/backend/pkg/api/schema"
	"majinyao.cn/my-app/backend/pkg/db/crud"
	"majinyao.cn/my-app/backend/pkg/utils"
)

type SaveInput[TSave any] struct {
	Body TSave
}

func (c *Crud[T, TItem, TDetail, TSave]) RegisterSave(api huma.API) (op huma.Operation) {
	op = huma.Operation{
		Method:      http.MethodPost,
		Path:        fmt.Sprintf("/%s", c.SlugifiedName),
		Summary:     fmt.Sprintf("Save %s", c.SpacedName),
		Description: fmt.Sprintf("Save %s, if it is a new entity which is transient or id=0, it will be created, otherwise it will be updated, and only the exposed fields will be updated.", c.LowerCaseSpacedName),
		OperationID: fmt.Sprintf("%s.save", c.SlugifiedName),
		Tags:        []string{c.EntityName},
	}

	for _, scheme := range c.Schemes {
		op.Security = append(op.Security, map[string][]string{
			scheme: {op.OperationID},
		})
	}

	var tsave TSave
	choice := crud.Choice{
		Selects: utils.SliceMap(
			reflect.VisibleFields(reflect.TypeOf(tsave)),
			func(e reflect.StructField) string {
				return e.Name
			},
		),
	}

	handler := func(ctx context.Context, input *SaveInput[TSave]) (output *schema.Response[string], err error) {
		service, cancel := c.GetCrudService(ctx, c.Db)
		defer cancel()

		var entity T
		err = copier.CopyWithOption(&entity, &input.Body, c.CopierOption)
		if err != nil {
			return nil, err
		}

		var affected int64
		if input.Body.IsTransient() {
			affected, err = service.Create(&entity)
		} else {
			affected, err = service.Update(&entity, choice)
		}
		if err != nil {
			return nil, err
		}
		if affected < 1 {
			return schema.Fail[string](http.StatusNotModified, "zero affected"), nil
		}

		return schema.Succeed(entity.GetId().HexString(), affected), nil
	}

	huma.Register(api, op, handler)
	return
}
