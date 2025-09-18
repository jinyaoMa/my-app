package schemas

import "majinyao.cn/my-app/backend/pkg/api/schema"

type FileCategoryItem struct {
	schema.EntityId
	schema.EntityTime
	Code string `json:"code" doc:"Code"`
	Name string `json:"name" doc:"Name"`
}

type FileCategoryDetail struct {
	schema.EntityId
	schema.EntityTime
	Code string `json:"code" doc:"Code"`
	Name string `json:"name" doc:"Name"`

	FileExtensions []FileExtensionItem `json:"fileExtensions" doc:"File Extensions"`
}

type FileCategorySave struct {
	schema.EntityId
	Code string `json:"code" doc:"Code"`
	Name string `json:"name" doc:"Name"`
}
