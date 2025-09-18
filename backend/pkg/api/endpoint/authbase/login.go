package authbase

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"majinyao.cn/my-app/backend/pkg/api/middlewares/authfwt"
	"majinyao.cn/my-app/backend/pkg/api/schema"
)

type LoginInput struct {
	Body      schema.AuthLoginRequest
	VisitorId string `header:"X-Visitor-ID" required:"true" doc:"Visitor Id"`
}

func (a *Auth[T]) RegisterLogin(api huma.API) (op huma.Operation) {
	op = huma.Operation{
		Method:      http.MethodPost,
		Path:        "/auth/login",
		Summary:     "Auth Login",
		Description: "Login by username and password, and return access token and refresh token. BTW, visitor id is required for distinguishing clients/devices.",
		OperationID: "auth.login",
		Tags:        []string{"Auth"},
	}

	handler := func(ctx context.Context, input *LoginInput) (output *schema.Response[schema.AuthLogin], err error) {
		f := authfwt.GetFwtFromContext[T](ctx)
		verifier, cancel := a.NewVerifier(ctx, a.Db)
		defer cancel()

		userdata, err := verifier.VerifyLogin(input)
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
