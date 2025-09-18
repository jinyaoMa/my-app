package api

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
	"majinyao.cn/my-app/backend/api/schemas"
	"majinyao.cn/my-app/backend/pkg/api/middlewares/authfwt"
	"majinyao.cn/my-app/backend/pkg/api/middlewares/memcache"
	"majinyao.cn/my-app/backend/pkg/flag"
	"majinyao.cn/my-app/backend/pkg/router"
)

func New(ctx context.Context, tx *gorm.DB, options Options) (router.IRouter, error) {
	return router.New("/api", options.Router, func(humaapi huma.API) error {
		var operationIdEnumMap map[string]int
		authfwt, err := authfwt.New[schemas.UserData](humaapi, authScheme, options.AuthFwt, func(ctx huma.Context, scopes []string) error {
			var requiredEnums []int
			for _, scope := range scopes {
				if enum, ok := operationIdEnumMap[scope]; ok {
					requiredEnums = append(requiredEnums, enum)
				}
			}

			cache := memcache.Get(ctx)
			userdata := authfwt.GetUserData[schemas.UserData](ctx)
			authPerm, _ := cache.Get(fmt.Sprintf("auth_perm_%d", userdata.UserId))
			perm, _ := authPerm.(flag.IFlag)
			if slices.ContainsFunc(requiredEnums, perm.IsOn) {
				return nil
			}
			return errors.New("no permission")
		})
		if err != nil {
			return err
		}

		humaapi.UseMiddleware(memcache.New(options.MemcacheLimit), authfwt)

		operationIdEnumMap, err = setup(ctx, tx, humaapi)
		if err != nil {
			return err
		}
		return nil
	})
}
