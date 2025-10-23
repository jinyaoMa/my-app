package auth

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"majinyao.cn/my-app/backend/api/apir"
	"majinyao.cn/my-app/backend/internal/service"
	"majinyao.cn/my-app/backend/internal/view"
	"majinyao.cn/my-app/backend/pkg/api/middlewares/authfwt"
	"majinyao.cn/my-app/backend/pkg/api/schema"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
)

type RoleInput struct {
	roleId       datatype.Id
	RoleId       string `query:"roleId" required:"true" maxLength:"32" doc:"Role Id (Base36)"`
	RefreshToken string `query:"refresh" required:"true" maxLength:"32" doc:"Refresh Token"`
	VisitorId    string `header:"X-Visitor-ID" required:"true" maxLength:"32" doc:"Visitor Id"`
}

func (i *RoleInput) Resolve(ctx huma.Context) (errs []error) {
	id, err := datatype.ParseIdFromB36(i.RoleId)
	if err != nil {
		errs = append(errs, err)
		return
	}

	i.roleId = id
	return
}

func (a *Auth) RegisterRole(api huma.API) (op huma.Operation) {
	op = huma.Operation{
		Method:       http.MethodGet,
		Path:         "/auth/role",
		Summary:      "Switch Role",
		Description:  "Switch current user role to specified role.",
		OperationID:  "auth.get.role",
		Tags:         []string{"Auth"},
		Errors:       []int{http.StatusInternalServerError},
		MaxBodyBytes: 1,
	}

	op.Security = append(op.Security, map[string][]string{
		apir.AuthScheme: nil,
	})

	handler := func(ctx context.Context, input *RoleInput) (output *schema.Response[schema.AuthLogin], err error) {
		f := authfwt.GetFwtFromContext[view.UserData](ctx)
		claims := authfwt.GetClaimsFromContext[view.UserData](ctx)
		service, cancel := service.NewUserService(ctx, a.CrudOptions, a.Db, a.Store)
		defer cancel()

		accessToken, refreshToken, expiredAt, err := f.Refresh(claims, input.RefreshToken, func(data view.UserData) (newData view.UserData, err error) {
			newData, err = service.VerifyUserDataRoleId(data, input.VisitorId, input.roleId)
			return
		})
		if err != nil {
			return schema.Fail[schema.AuthLogin](http.StatusInternalServerError, err.Error()), nil
		}

		return schema.Succeed(schema.AuthLogin{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			ExpiredAt:    expiredAt,
		}, 1), nil
	}

	huma.Register(api, op, handler)
	return
}
