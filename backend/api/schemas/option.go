package schemas

import (
	"majinyao.cn/my-app/backend/pkg/api/schema"
)

type OptionItem struct {
	schema.ModelId
	schema.ModelTime
	schema.ModelReserved
	schema.ModelHmac
	Key   string `json:"key" doc:"Key"`
	Value string `json:"value" doc:"Value"`
}

type OptionSave struct {
	schema.ModelId
	Key   string `json:"key" doc:"Key"`
	Value string `json:"value" doc:"Value"`
}

type OptionDetail struct {
	OptionSave
	schema.ModelTime
	schema.ModelReserved
	schema.ModelHmac
}

type OptionSearchItem struct {
	schema.ModelId
	Key   string `json:"key" doc:"Key"`
	Value string `json:"value" doc:"Value"`
}
