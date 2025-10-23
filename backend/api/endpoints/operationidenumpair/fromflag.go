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

type FromFlagInput struct {
	FlagBase64 string `query:"flag" required:"true" doc:"Flag (Base64)"`
}

func (p *OperationIdEnumPair) RegisterFromFlag(api huma.API) (op huma.Operation) {
	op = huma.Operation{
		Method:       http.MethodPost,
		Path:         "/operation-id-enum-pair/from-flag",
		Summary:      "Get Pair List From Flag",
		Description:  "Get operation id enum pair list from flag (base64).",
		OperationID:  "operation-id-enum-pair.post.from-flag",
		Tags:         []string{"OperationIdEnumPair"},
		Errors:       []int{http.StatusInternalServerError},
		MaxBodyBytes: 1,
	}

	op.Security = append(op.Security, map[string][]string{
		apir.AuthScheme: {op.OperationID},
	})

	handler := func(ctx context.Context, input *FromFlagInput) (output *schema.Response[[]schemas.OperationIdEnumPairItem], err error) {
		service, cancel := service.NewOperationIdEnumPairService(ctx, p.Db)
		defer cancel()

		var list []schemas.OperationIdEnumPairItem
		_, total, err := service.ListCopyFromFlagBase64(&list, input.FlagBase64)
		if err != nil {
			return nil, err
		}
		return schema.Succeed(list, total), nil
	}

	huma.Register(api, op, handler)
	return
}
