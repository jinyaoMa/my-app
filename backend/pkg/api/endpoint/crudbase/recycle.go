package crudbase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/pkg/api/schema"
	"majinyao.cn/my-app/backend/pkg/db/crud"
)

type RecycleInput struct {
	Body crud.Criteria
}

func (c *Crud[T, TItem, TSave, TDetail, TSearchItem]) RegisterRecycle(api huma.API, fieldName string, conds ...func(tx *gorm.DB) (*gorm.DB, error)) (op huma.Operation) {
	op = huma.Operation{
		Method:       http.MethodPost,
		Path:         fmt.Sprintf("/%s/recycle", c.SlugifiedName),
		Summary:      fmt.Sprintf("Get %s Recycle List", c.SpacedName),
		Description:  fmt.Sprintf("Get deleted %s and present results in list, and users can filter, sort and page data by passing criteria.", c.LowerCaseSpacedName),
		OperationID:  fmt.Sprintf("%s.post.recycle", c.SlugifiedName),
		Tags:         []string{c.EntityName},
		Errors:       []int{http.StatusInternalServerError},
		MaxBodyBytes: 4 * 1024,
	}

	for _, scheme := range c.Schemes {
		op.Security = append(op.Security, map[string][]string{
			scheme: {op.OperationID},
		})
	}

	handler := func(ctx context.Context, input *RecycleInput) (output *schema.Response[[]TItem], err error) {
		service, cancel := c.GetCrudService(ctx)
		defer cancel()

		var list []TItem
		_, total, err := service.QueryCopy(&list, input.Body, append(conds, func(tx *gorm.DB) (*gorm.DB, error) {
			if fieldName != "" {
				tx = tx.Where(fieldName + " IS NOT NULL")
			}
			return tx.Unscoped().Where("deleted_at IS NOT NULL"), nil
		})...)
		if err != nil {
			return nil, err
		}
		return schema.Succeed(list, total), nil
	}

	huma.Register(api, op, handler)
	return
}
