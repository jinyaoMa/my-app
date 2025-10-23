package user

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"majinyao.cn/my-app/backend/api/apir"
	"majinyao.cn/my-app/backend/internal/service"
	"majinyao.cn/my-app/backend/internal/view"
	"majinyao.cn/my-app/backend/pkg/api/schema"
)

func (u *User) RegisterCurrent(api huma.API) (op huma.Operation) {
	op = huma.Operation{
		Method:       http.MethodGet,
		Path:         "/user/current",
		Summary:      "Get Current User",
		Description:  "Get current user.",
		OperationID:  "user.get.current",
		Tags:         []string{"User"},
		Errors:       []int{http.StatusInternalServerError},
		MaxBodyBytes: 1,
	}

	op.Security = append(op.Security, map[string][]string{
		apir.AuthScheme: {},
	})

	handler := func(ctx context.Context, input *struct{}) (output *schema.Response[view.CurrentUser], err error) {
		service, cancel := service.NewUserService(ctx, u.CrudOptions, u.Db, u.Store)
		defer cancel()

		currentUser, err := service.GetCurrentUser()
		if err != nil {
			return nil, err
		}
		return schema.Succeed(currentUser, 1), nil
	}

	huma.Register(api, op, handler)
	return
}
