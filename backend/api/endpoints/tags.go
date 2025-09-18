package endpoints

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"majinyao.cn/my-app/backend/pkg/api/schema"
	"majinyao.cn/my-app/backend/pkg/utils"
)

func (e *endpoints) RegisterTags(api huma.API) (op huma.Operation) {
	op = huma.Operation{
		Method:      http.MethodGet,
		Path:        "/endpoints/tags",
		Summary:     "Get Tags of Endpoints",
		Description: "Get all tags of registered endpoints.",
		OperationID: "endpoints.tags",
		Tags:        []string{"Endpoints"},
	}

	if e.Scheme != "" {
		op.Security = append(op.Security, map[string][]string{
			e.Scheme: {op.OperationID},
		})
	}

	e.Ops = append(e.Ops, op)
	handler := func(ctx context.Context, input *struct{}) (output *schema.Response[[]string], err error) {
		var tags []string
		for _, op := range e.Ops {
			tags = append(tags, op.Tags...)
		}
		tags = utils.SliceUnique(tags, func(tag string) string {
			return tag
		})
		return schema.Succeed(tags, int64(len(tags))), nil
	}

	huma.Register(api, op, handler)
	return
}
