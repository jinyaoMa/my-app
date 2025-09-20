package api

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/api/schemas"
	"majinyao.cn/my-app/backend/internal/service"
	"majinyao.cn/my-app/backend/pkg/api/middlewares/authfwt"
	"majinyao.cn/my-app/backend/pkg/api/middlewares/memcache"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
	"majinyao.cn/my-app/backend/pkg/flag"
	"majinyao.cn/my-app/backend/pkg/router"
)

func New(ctx context.Context, db *gorm.DB, options Options) (router.IRouter, error) {
	return router.New("/api", options.Router, func(humaapi huma.API) error {
		var operationIdEnumMap map[string]int
		authfwt, err := authfwt.New[schemas.UserData](humaapi, authScheme, options.AuthFwt, func(ctx huma.Context, scopes []string) error {
			var requiredEnums []int
			for _, scope := range scopes {
				if enum, ok := operationIdEnumMap[scope]; ok {
					requiredEnums = append(requiredEnums, enum)
				}
			}

			cache := memcache.GetFromHumaContext(ctx)
			claims := authfwt.GetClaimsFromHumaContext[schemas.UserData](ctx)

			var perm flag.IFlag
			authPerm, err := cache.Get(fmt.Sprintf("auth_perm_%d", claims.Data.UserId))
			if err == nil {
				perm, _ = authPerm.(flag.IFlag)
			} else {
				userService, cancel := service.NewUserService(ctx.Context(), db)
				defer cancel()

				user, notFound, err := userService.GetById(datatype.Id(claims.Data.UserId), "Roles.Permissions")
				if err != nil {
					return err
				}
				if notFound {
					return errors.New("user not found")
				}

				for _, role := range user.Roles {
					for _, p := range role.Permissions {
						perm = p.GetFlag().Or(perm)
					}
				}

				_, _, err = cache.Set(fmt.Sprintf("auth_perm_%d", user.Id.Int64()), perm)
				if err != nil {
					return err
				}
			}

			if perm != nil && slices.ContainsFunc(requiredEnums, perm.IsOn) {
				return nil
			}
			return errors.New("no permission")
		})
		if err != nil {
			return err
		}

		humaapi.UseMiddleware(memcache.New(options.MemcacheLimit), authfwt)

		operationIdEnumMap, err = setup(ctx, db, humaapi)
		if err != nil {
			return err
		}
		return nil
	})
}
