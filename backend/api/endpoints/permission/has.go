package permission

import (
	"context"
	"net/http"
	"slices"
	"strings"

	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/api/apir"
	"majinyao.cn/my-app/backend/internal/service"
	"majinyao.cn/my-app/backend/pkg/api/schema"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
)

type HasInput struct {
	exceptId *datatype.Id
	ExceptId string `query:"exceptId" required:"false" doc:"Except Permission Id"`
	Field    string `query:"field" required:"true" doc:"Check Field"`
	Value    string `query:"value" required:"true" doc:"Check Value"`
}

func (i *HasInput) Resolve(ctx huma.Context) (errs []error) {
  if i.ExceptId == "" {
		return
	}

	id, err := datatype.ParseIdFromB36(i.ExceptId)
	if err != nil {
		errs = append(errs, err)
		return
	}

	i.exceptId = &id
	return
}

func (p *Permission) RegisterHas(api huma.API) (op huma.Operation) {
	op = huma.Operation{
		Method:       http.MethodGet,
		Path:         "/permission/has",
		Summary:      "Permission Has Field Value",
		Description:  "Check Permission if the given field with the given value exists.",
		OperationID:  "permission.get.has",
		Tags:         []string{"Permission"},
		Errors:       []int{http.StatusInternalServerError},
		MaxBodyBytes: 1,
	}

	op.Security = append(op.Security, map[string][]string{
		apir.AuthScheme: {op.OperationID},
	})

	availFields := []string{"code"}

	handler := func(ctx context.Context, input *HasInput) (output *schema.Response[bool], err error) {
		service, cancel := service.NewPermissionService(ctx, p.CrudOptions, p.Db)
		defer cancel()

		if !slices.Contains(availFields, strings.ToLower(input.Field)) {
			return schema.Fail[bool](http.StatusBadRequest, "field not available"), nil
		}

		total, err := service.Count(func(tx *gorm.DB) (*gorm.DB, error) {
      tx = tx.Unscoped()
      if input.exceptId != nil {
				tx = tx.Where("id <> ?", *input.exceptId)
			}
			return tx.Where(input.Field+" = ?", input.Value), nil
		})
		if err != nil {
			return nil, err
		}
		return schema.Succeed(total > 0, total), nil
	}

	huma.Register(api, op, handler)
	return
}
