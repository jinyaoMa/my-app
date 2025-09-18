package endpoints

import (
	"context"
	"net/http"
	"slices"
	"strings"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jinzhu/copier"
	"majinyao.cn/my-app/backend/api/schemas"
	"majinyao.cn/my-app/backend/pkg/api/schema"
	"majinyao.cn/my-app/backend/pkg/utils"
)

type ListInput struct {
	Tag         string `query:"tag" doc:"Filter by Tag"`
	OperationID string `query:"operationId" doc:"Filter by Operation ID"`
}

func (e *endpoints) RegisterList(api huma.API) (op huma.Operation) {
	op = huma.Operation{
		Method:      http.MethodGet,
		Path:        "/endpoints/list",
		Summary:     "List Endpoints",
		Description: "List all registered endpoints, filtered by tag, operation id, etc.",
		OperationID: "endpoints.list",
		Tags:        []string{"Endpoints"},
	}

	if e.Scheme != "" {
		op.Security = append(op.Security, map[string][]string{
			e.Scheme: {op.OperationID},
		})
	}

	e.Ops = append(e.Ops, op)
	handler := func(ctx context.Context, input *ListInput) (output *schema.Response[[]schemas.EndpointsItem], err error) {
		var list []schemas.EndpointsItem
		if err := copier.Copy(&list, &e.Ops); err != nil {
			return nil, err
		}
		if input.Tag != "" {
			list = utils.SliceFilter(list, func(op schemas.EndpointsItem) bool {
				return slices.ContainsFunc(op.Tags, func(tag string) bool {
					return strings.Contains(tag, input.Tag)
				})
			})
		}
		if input.OperationID != "" {
			list = utils.SliceFilter(list, func(op schemas.EndpointsItem) bool {
				return strings.Contains(op.OperationID, input.OperationID)
			})
		}
		return schema.Succeed(list, int64(len(list))), nil
	}

	huma.Register(api, op, handler)
	return
}
