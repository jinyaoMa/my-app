package crudbase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"majinyao.cn/my-app/backend/pkg/api/schema"
	"majinyao.cn/my-app/backend/pkg/db"
)

type DeleteInput struct {
	id int64
	Id string `query:"id"`
}

func (i *DeleteInput) Resolve(ctx huma.Context) (errs []error) {
	id, err := db.ConvertStringToId(i.Id)
	if err != nil {
		errs = append(errs, err)
		return
	}

	i.id = id
	return
}

func (c *Crud[T, TItem, TDetail, TSave]) RegisterDelete(api huma.API) (op huma.Operation) {
	op = huma.Operation{
		Method:      http.MethodDelete,
		Path:        fmt.Sprintf("/%s", c.SlugifiedName),
		Summary:     fmt.Sprintf("Delete %s", c.SpacedName),
		Description: fmt.Sprintf("Delete %s by id.", c.LowerCaseSpacedName),
		OperationID: fmt.Sprintf("%s.delete", c.SlugifiedName),
		Tags:        []string{c.EntityName},
	}

	for _, scheme := range c.Schemes {
		op.Security = append(op.Security, map[string][]string{
			scheme: {op.OperationID},
		})
	}

	handler := func(ctx context.Context, input *DeleteInput) (output *schema.Response[string], err error) {
		service, cancel := c.GetCrudService(ctx, c.Db)
		defer cancel()

		affected, err := service.Delete(input.id)
		if err != nil {
			return nil, err
		}
		if affected < 1 {
			return schema.Fail[string](http.StatusNotModified, "zero affected"), nil
		}

		return schema.Succeed(input.Id, affected), nil
	}

	huma.Register(api, op, handler)
	return
}
