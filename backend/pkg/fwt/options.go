package fwt

import (
	"time"

	"majinyao.cn/my-app/backend/pkg/codegen"
	"majinyao.cn/my-app/backend/pkg/crypto/hasher"
)

type Options struct {
	Hasher        hasher.Options  `json:"hasher"`
	Codegen       codegen.Options `json:"codegen"`
	Issuer        string          `json:"issuer"`
	Subject       string          `json:"subject"`
	Epoch         time.Time       `json:"epoch"`
	ExpiredAge    int             `json:"expiredAge"`    // in seconds
	RefreshAge    int             `json:"refreshAge"`    // in seconds
	RefreshLength int             `json:"refreshLength"` // refresh token length
}
