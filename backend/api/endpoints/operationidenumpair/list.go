package operationidenumpair

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jinzhu/copier"
	"majinyao.cn/my-app/backend/api/schemas"
	"majinyao.cn/my-app/backend/internal/service"
	"majinyao.cn/my-app/backend/pkg/api/schema"
	"majinyao.cn/my-app/backend/pkg/db"
)

type ListInput struct {
	OperationID string `query:"operationId" doc:"Filter by Operation ID"`
}

func (p *OperationIdEnumPair) RegisterList(api huma.API) (op huma.Operation) {
	op = huma.Operation{
		Method:      http.MethodGet,
		Path:        "/operation-id-enum-pair/list",
		Summary:     "Get Operation Id Enum Pair List",
		Description: "Get operation id enum pair list, filtered by operation id.",
		OperationID: "operation-id-enum-pair.list",
		Tags:        []string{"OperationIdEnumPair"},
	}

	if p.Scheme != "" {
		op.Security = append(op.Security, map[string][]string{
			p.Scheme: {op.OperationID},
		})
	}

	handler := func(ctx context.Context, input *ListInput) (output *schema.Response[[]schemas.OperationIdEnumPairItem], err error) {
		service, cancel := service.NewOperationIdEnumPairService(ctx, p.Db)
		defer cancel()

		entities, total, err := service.List(input.OperationID)
		if err != nil {
			return nil, err
		}

		var list []schemas.OperationIdEnumPairItem
		err = copier.CopyWithOption(&list, &entities, db.DefaultCopierOption)
		if err != nil {
			return nil, err
		}

		return schema.Succeed(list, total), nil
	}

	huma.Register(api, op, handler)
	return
}
