package crudbase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/api/schema"
)

type SearchInput struct {
	Body schema.CrudSearchRequest
}

func (c *Crud[T, TItem, TSave, TDetail, TSearchItem]) RegisterSearch(api huma.API, conds ...func(tx *gorm.DB) (*gorm.DB, error)) (op huma.Operation) {
	op = huma.Operation{
		Method:       http.MethodPost,
		Path:         fmt.Sprintf("/%s/search", c.SlugifiedName),
		Summary:      fmt.Sprintf("Search %s", c.SpacedName),
		Description:  fmt.Sprintf("Search %s by given keyword and fields, and limit the result count.", c.LowerCaseSpacedName),
		OperationID:  fmt.Sprintf("%s.post.search", c.SlugifiedName),
		Tags:         []string{c.EntityName},
		Errors:       []int{http.StatusInternalServerError},
		MaxBodyBytes: 2 * 1024,
	}

	for _, scheme := range c.Schemes {
		op.Security = append(op.Security, map[string][]string{
			scheme: {op.OperationID},
		})
	}

	handler := func(ctx context.Context, input *SearchInput) (output *schema.Response[[]TSearchItem], err error) {
		service, cancel := c.GetCrudService(ctx)
		defer cancel()

		var list []TSearchItem
		_, total, err := service.SearchCopy(&list,
			input.Body.Keyword,
			input.Body.Fields,
			input.Body.Limit,
			input.Body.Exact,
			input.Body.Includes,
			conds...)
		if err != nil {
			return nil, err
		}
		return schema.Succeed(list, total), nil
	}

	huma.Register(api, op, handler)
	return
}
