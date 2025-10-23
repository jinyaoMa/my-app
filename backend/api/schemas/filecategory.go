package schemas

import "majinyao.cn/my-app/backend/pkg/api/schema"

type FileCategoryItem struct {
	schema.ModelId
	schema.ModelTime
	schema.ModelReserved
	Code string `json:"code" doc:"Code"`
	Name string `json:"name" doc:"Name"`
}

type FileCategorySave struct {
	schema.ModelId
	Code string `json:"code" doc:"Code"`
	Name string `json:"name" doc:"Name"`
}

type FileCategoryDetail struct {
	FileCategorySave
	schema.ModelTime
	schema.ModelReserved

	FileExtensions []FileExtensionItem `json:"fileExtensions" doc:"File Extensions"`
}

type FileCategorySearchItem struct {
	schema.ModelId
	Code string `json:"code" doc:"Code"`
	Name string `json:"name" doc:"Name"`
}
