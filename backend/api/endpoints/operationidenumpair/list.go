package operationidenumpair

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"majinyao.cn/my-app/backend/api/apir"
	"majinyao.cn/my-app/backend/api/schemas"
	"majinyao.cn/my-app/backend/internal/service"
	"majinyao.cn/my-app/backend/pkg/api/schema"
)

type ListInput struct {
	OperationID string `query:"operationId" doc:"Filter by Operation ID"`
}

func (p *OperationIdEnumPair) RegisterList(api huma.API) (op huma.Operation) {
	op = huma.Operation{
		Method:       http.MethodGet,
		Path:         "/operation-id-enum-pair/list",
		Summary:      "Get Pair List",
		Description:  "Get operation id enum pair list, filtered by operation id.",
		OperationID:  "operation-id-enum-pair.get.list",
		Tags:         []string{"OperationIdEnumPair"},
		Errors:       []int{http.StatusInternalServerError},
		MaxBodyBytes: 1,
	}

	op.Security = append(op.Security, map[string][]string{
		apir.AuthScheme: {op.OperationID},
	})

	handler := func(ctx context.Context, input *ListInput) (output *schema.Response[[]schemas.OperationIdEnumPairItem], err error) {
		service, cancel := service.NewOperationIdEnumPairService(ctx, p.Db)
		defer cancel()

		var list []schemas.OperationIdEnumPairItem
		_, total, err := service.ListCopy(&list, input.OperationID)
		if err != nil {
			return nil, err
		}
		return schema.Succeed(list, total), nil
	}

	huma.Register(api, op, handler)
	return
}
