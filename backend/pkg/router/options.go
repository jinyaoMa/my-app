package router

import (
	"time"

	"majinyao.cn/my-app/backend/pkg/cflog"
)

type Options struct {
	DocsPath    string `json:"docsPath"`
	DocsTitle   string `json:"docsTitle"`
	DocsVersion string `json:"docsVersion"`

	Cflog            cflog.Options `json:"cflog"`
	StaticsDirectory string        `json:"staticsDirectory"`

	EnableTimeout bool          `json:"enableTimeout"`
	Timeout       time.Duration `json:"timeout"` // timeout in seconds

	EnableCors     bool     `json:"enableCors"`
	AllowedOrigins []string `json:"allowedOrigins"`
	AllowedMethods []string `json:"allowedMethods"`
	AllowedHeaders []string `json:"allowedHeaders"`

	EnableHttpRate bool `json:"enableHttpRate"`
	LimitByIp      bool `json:"limitByIp"`
	RateLimit      int  `json:"rateLimit"` // limit per minute
}
