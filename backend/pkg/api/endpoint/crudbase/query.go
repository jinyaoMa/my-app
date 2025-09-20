package crudbase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"majinyao.cn/my-app/backend/pkg/api/schema"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

type QueryInput struct {
	Body crud.Criteria
}

func (c *Crud[T, TItem, TDetail, TSave]) RegisterQuery(api huma.API) (op huma.Operation) {
	op = huma.Operation{
		Method:      http.MethodPost,
		Path:        fmt.Sprintf("/%s/query", c.SlugifiedName),
		Summary:     fmt.Sprintf("Query %s List", c.SpacedName),
		Description: fmt.Sprintf("Query %s and present results in list, and users can filter, sort and page data by passing criteria.", c.LowerCaseSpacedName),
		OperationID: fmt.Sprintf("%s.query", c.SlugifiedName),
		Tags:        []string{c.EntityName},
	}

	for _, scheme := range c.Schemes {
		op.Security = append(op.Security, map[string][]string{
			scheme: {op.OperationID},
		})
	}

	handler := func(ctx context.Context, input *QueryInput) (output *schema.Response[[]TItem], err error) {
		service, cancel := c.GetCrudService(ctx, c.Db)
		defer cancel()

		var list []TItem
		_, total, err := service.QueryCopy(&list, input.Body)
		if err != nil {
			return nil, err
		}
		return schema.Succeed(list, total), nil
	}

	huma.Register(api, op, handler)
	return
}
