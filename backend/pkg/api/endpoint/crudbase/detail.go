package crudbase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"majinyao.cn/my-app/backend/pkg/api/schema"
	"majinyao.cn/my-app/backend/pkg/db/crud"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
)

type DetailInput struct {
	id   datatype.Id
	Id   string        `path:"id" required:"true" doc:"Model Id (Base36)"`
	Body crud.Includes `doc:"Detail Includes"`
}

func (i *DetailInput) Resolve(ctx huma.Context) (errs []error) {
	id, err := datatype.ParseIdFromB36(i.Id)
	if err != nil {
		errs = append(errs, err)
		return
	}

	i.id = id
	return
}

func (c *Crud[T, TItem, TSave, TDetail, TSearchItem]) RegisterDetail(api huma.API) (op huma.Operation) {
	op = huma.Operation{
		Method:       http.MethodPost,
		Path:         fmt.Sprintf("/%s/detail/{id}", c.SlugifiedName),
		Summary:      fmt.Sprintf("Get %s Detail", c.SpacedName),
		Description:  fmt.Sprintf("Get %s detail by id, and users can include associated entities by passing includes.", c.LowerCaseSpacedName),
		OperationID:  fmt.Sprintf("%s.post.detail", c.SlugifiedName),
		Tags:         []string{c.EntityName},
		Errors:       []int{http.StatusInternalServerError},
		MaxBodyBytes: 2 * 1024,
	}

	for _, scheme := range c.Schemes {
		op.Security = append(op.Security, map[string][]string{
			scheme: {op.OperationID},
		})
	}

	handler := func(ctx context.Context, input *DetailInput) (output *schema.Response[TDetail], err error) {
		service, cancel := c.GetCrudService(ctx)
		defer cancel()

		var detail TDetail
		_, notFound, err := service.GetCopyById(&detail, input.id, input.Body)
		if err != nil {
			return nil, err
		}
		if notFound {
			return schema.Fail[TDetail](http.StatusNotFound, "not found"), nil
		}
		return schema.Succeed(detail, 1), nil
	}

	huma.Register(api, op, handler)
	return
}
