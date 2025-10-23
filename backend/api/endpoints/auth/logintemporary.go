package auth

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"majinyao.cn/my-app/backend/internal/service"
	"majinyao.cn/my-app/backend/internal/view"
	"majinyao.cn/my-app/backend/pkg/api/endpoint/authbase"
	"majinyao.cn/my-app/backend/pkg/api/middlewares/authfwt"
	"majinyao.cn/my-app/backend/pkg/api/schema"
)

func (a *Auth) RegisterLoginTemporary(api huma.API) (op huma.Operation) {
	op = huma.Operation{
		Method:       http.MethodPost,
		Path:         "/auth/login/temporary",
		Summary:      "Login Temporary",
		Description:  "Login by username and temporary password, and return access token and refresh token. BTW, visitor id is required for distinguishing clients/devices.",
		OperationID:  "auth.post.login.temporary",
		Tags:         []string{"Auth"},
		Errors:       []int{http.StatusInternalServerError},
		MaxBodyBytes: 512,
	}

	handler := func(ctx context.Context, input *authbase.LoginInput) (output *schema.Response[schema.AuthLogin], err error) {
		f := authfwt.GetFwtFromContext[view.UserData](ctx)
		service, cancel := service.NewUserService(ctx, a.CrudOptions, a.Db, a.Store)
		defer cancel()

		userdata, err := service.VerifyLoginWithTemporaryPassword(input)
		if err != nil {
			return schema.Fail[schema.AuthLogin](http.StatusUnauthorized, err.Error()), nil
		}

		accessToken, refreshToken, expiredAt, err := f.Generate(userdata)
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
