package authfwt

import "majinyao.cn/my-app/backend/pkg/fwt"

type Options struct {
	Fwt        fwt.Options `json:"fwt"`
	CacheLimit int         `json:"cacheLimit"`
}
