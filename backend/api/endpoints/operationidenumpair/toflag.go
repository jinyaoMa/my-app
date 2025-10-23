package operationidenumpair

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"majinyao.cn/my-app/backend/api/apir"
	"majinyao.cn/my-app/backend/internal/service"
	"majinyao.cn/my-app/backend/pkg/api/schema"
)

type ToFlagInput struct {
	Body []int `doc:"Operation Id Enums"`
}

func (p *OperationIdEnumPair) RegisterToFlag(api huma.API) (op huma.Operation) {
	op = huma.Operation{
		Method:      http.MethodPost,
		Path:        "/operation-id-enum-pair/to-flag",
		Summary:     "Convert Enums to Flag",
		Description: "Convert enums of operation ids/roles to base64 flag.",
		OperationID: "operation-id-enum-pair.post.to-flag",
		Tags:        []string{"OperationIdEnumPair"},
		Errors:      []int{http.StatusInternalServerError},
		// MaxBodyBytes: 4 * 1024,
	}

	op.Security = append(op.Security, map[string][]string{
		apir.AuthScheme: {op.OperationID},
	})

	handler := func(ctx context.Context, input *ToFlagInput) (output *schema.Response[string], err error) {
		service, cancel := service.NewOperationIdEnumPairService(ctx, p.Db)
		defer cancel()

		f, err := service.ConvertEnumsToFlag(input.Body)
		if err != nil {
			return nil, err
		}
		return schema.Succeed(f.ToBase64(), 1), nil
	}

	huma.Register(api, op, handler)
	return
}
