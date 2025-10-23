package crudbase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"majinyao.cn/my-app/backend/pkg/api/schema"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
)

type DeleteInput struct {
	id datatype.Id
	Id string `path:"id" required:"true" doc:"Model Id (Base36)"`
}

func (i *DeleteInput) Resolve(ctx huma.Context) (errs []error) {
	id, err := datatype.ParseIdFromB36(i.Id)
	if err != nil {
		errs = append(errs, err)
		return
	}

	i.id = id
	return
}

func (c *Crud[T, TItem, TSave, TDetail, TSearchItem]) RegisterDelete(api huma.API) (op huma.Operation) {
	op = huma.Operation{
		Method:       http.MethodDelete,
		Path:         fmt.Sprintf("/%s/delete/{id}", c.SlugifiedName),
		Summary:      fmt.Sprintf("Delete %s", c.SpacedName),
		Description:  fmt.Sprintf("Delete %s by id.", c.LowerCaseSpacedName),
		OperationID:  fmt.Sprintf("%s.delete.delete", c.SlugifiedName),
		Tags:         []string{c.EntityName},
		Errors:       []int{http.StatusInternalServerError},
		MaxBodyBytes: 1,
	}

	for _, scheme := range c.Schemes {
		op.Security = append(op.Security, map[string][]string{
			scheme: {op.OperationID},
		})
	}

	handler := func(ctx context.Context, input *DeleteInput) (output *schema.Response[string], err error) {
		service, cancel := c.GetCrudService(ctx)
		defer cancel()

		_, affected, err := service.Delete(input.id)
		if err != nil {
			return nil, err
		}
		if affected < 1 {
			return schema.Fail[string](http.StatusNotModified, "zero affected"), nil
		}
		return schema.Succeed(input.Id, affected), nil
	}

	huma.Register(api, op, handler)
	return
}

type BatchDeleteInput struct {
	ids  []datatype.Id
	Body []string `doc:"Model Ids (Base36)"`
}

func (i *BatchDeleteInput) Resolve(ctx huma.Context) (errs []error) {
	var rest []string
	for _, idB36 := range i.Body {
		id, err := datatype.ParseIdFromB36(idB36)
		if err != nil {
			continue
		}

		i.ids = append(i.ids, id)
		rest = append(rest, idB36)
	}
	i.Body = rest
	return
}

func (c *Crud[T, TItem, TSave, TDetail, TSearchItem]) RegisterBatchDelete(api huma.API) (op huma.Operation) {
	op = huma.Operation{
		Method:      http.MethodDelete,
		Path:        fmt.Sprintf("/%s/batch/delete", c.SlugifiedName),
		Summary:     fmt.Sprintf("Batch Delete %s", c.SpacedName),
		Description: fmt.Sprintf("Batch delete %s by ids.", c.LowerCaseSpacedName),
		OperationID: fmt.Sprintf("%s.delete.batch.delete", c.SlugifiedName),
		Tags:        []string{c.EntityName},
		Errors:      []int{http.StatusInternalServerError},
		// MaxBodyBytes: 4 * 1024,
	}

	for _, scheme := range c.Schemes {
		op.Security = append(op.Security, map[string][]string{
			scheme: {op.OperationID},
		})
	}

	handler := func(ctx context.Context, input *BatchDeleteInput) (output *schema.Response[[]string], err error) {
		service, cancel := c.GetCrudService(ctx)
		defer cancel()

		affected, err := service.BatchDelete(input.ids)
		if err != nil {
			return nil, err
		}
		if affected < 1 {
			return schema.Fail[[]string](http.StatusNotModified, "zero affected"), nil
		}
		return schema.Succeed(input.Body, affected), nil
	}

	huma.Register(api, op, handler)
	return
}
