package schemas

import (
	"majinyao.cn/my-app/backend/pkg/api/schema"
)

type OptionItem struct {
	schema.EntityId
	schema.EntityTime
	Key   string `json:"key" doc:"Key"`
	Value string `json:"value" doc:"Value"`
}

type OptionDetail struct {
	schema.EntityId
	schema.EntityTime
	Key   string `json:"key" doc:"Key"`
	Value string `json:"value" doc:"Value"`
}

type OptionSave struct {
	schema.EntityId
	Key   string `json:"key" doc:"Key"`
	Value string `json:"value" doc:"Value"`
}
