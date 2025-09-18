package api

import (
	"majinyao.cn/my-app/backend/pkg/api/middlewares/authfwt"
	"majinyao.cn/my-app/backend/pkg/router"
)

type Options struct {
	Router        router.Options  `json:"router"`
	AuthFwt       authfwt.Options `json:"authFwt"`
	MemcacheLimit int             `json:"memcacheLimit"`
}
