package crudbase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jinzhu/copier"
	"majinyao.cn/my-app/backend/pkg/api/schema"
	"majinyao.cn/my-app/backend/pkg/db"
)

type DetailInput struct {
	id       int64
	Id       string   `query:"id"`
	Includes []string `query:"includes" doc:"Included Associations"`
}

func (i *DetailInput) Resolve(ctx huma.Context) (errs []error) {
	id, err := db.ConvertStringToId(i.Id)
	if err != nil {
		errs = append(errs, err)
		return
	}

	i.id = id
	return
}

func (c *Crud[T, TItem, TDetail, TSave]) RegisterDetail(api huma.API) (op huma.Operation) {
	op = huma.Operation{
		Method:      http.MethodGet,
		Path:        fmt.Sprintf("/%s", c.SlugifiedName),
		Summary:     fmt.Sprintf("Get %s Detail", c.SpacedName),
		Description: fmt.Sprintf("Get %s detail by id, and users can include associated entities by passing includes.", c.LowerCaseSpacedName),
		OperationID: fmt.Sprintf("%s.detail", c.SlugifiedName),
		Tags:        []string{c.EntityName},
	}

	for _, scheme := range c.Schemes {
		op.Security = append(op.Security, map[string][]string{
			scheme: {op.OperationID},
		})
	}

	handler := func(ctx context.Context, input *DetailInput) (output *schema.Response[TDetail], err error) {
		service, cancel := c.GetCrudService(ctx, c.Db)
		defer cancel()

		entity, notFound, err := service.GetById(input.id, input.Includes...)
		if err != nil {
			return nil, err
		}
		if notFound {
			return schema.Fail[TDetail](http.StatusNotFound, "not found"), nil
		}

		var detail TDetail
		err = copier.CopyWithOption(&detail, &entity, c.CopierOption)
		if err != nil {
			return nil, err
		}

		return schema.Succeed(detail, 1), nil
	}

	huma.Register(api, op, handler)
	return
}
