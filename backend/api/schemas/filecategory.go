package schemas

import "majinyao.cn/my-app/backend/pkg/api/schema"

type FileCategoryItem struct {
	schema.ModelId
	schema.ModelTime
	Code string `json:"code" doc:"Code"`
	Name string `json:"name" doc:"Name"`
}

type FileCategoryDetail struct {
	schema.ModelId
	schema.ModelTime
	Code string `json:"code" doc:"Code"`
	Name string `json:"name" doc:"Name"`

	FileExtensions []FileExtensionItem `json:"fileExtensions" doc:"File Extensions"`
}

type FileCategorySave struct {
	schema.ModelId
	Code string `json:"code" doc:"Code"`
	Name string `json:"name" doc:"Name"`
}
