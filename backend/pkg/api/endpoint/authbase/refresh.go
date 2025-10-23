package authbase

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"majinyao.cn/my-app/backend/pkg/api/middlewares/authfwt"
	"majinyao.cn/my-app/backend/pkg/api/schema"
)

type RefreshInput struct {
	RefreshToken string `path:"refresh" required:"true" maxLength:"32" doc:"Refresh Token"`
	VisitorId    string `header:"X-Visitor-ID" required:"true" maxLength:"32" doc:"Visitor Id"`
}

func (a *Auth[T]) RegisterRefresh(api huma.API) (op huma.Operation) {
	op = huma.Operation{
		Method:       http.MethodGet,
		Path:         "/auth/refresh/{refresh}",
		Summary:      "Refresh",
		Description:  "Refresh by refresh token, and return access token and refresh token.",
		OperationID:  "auth.get.refresh",
		Tags:         []string{"Auth"},
		Errors:       []int{http.StatusInternalServerError},
		MaxBodyBytes: 1,
	}

	if a.Scheme != "" {
		op.Security = append(op.Security, map[string][]string{
			a.Scheme: nil,
		})
	}

	handler := func(ctx context.Context, input *RefreshInput) (output *schema.Response[schema.AuthLogin], err error) {
		f := authfwt.GetFwtFromContext[T](ctx)
		claims := authfwt.GetClaimsFromContext[T](ctx)
		verifier, cancel := a.NewVerifier(ctx)
		defer cancel()

		accessToken, refreshToken, expiredAt, err := f.Refresh(claims, input.RefreshToken, func(data T) (newData T, err error) {
			newData, err = verifier.VerifyUserData(data, input.VisitorId)
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
