package schemas

import "majinyao.cn/my-app/backend/pkg/api/schema"

type FileExtensionItem struct {
	schema.ModelId
	schema.ModelTime
	schema.ModelReserved
	Ext  string `json:"ext" doc:"Ext"`
	Name string `json:"name" doc:"Name"`
	Mime string `json:"mime" doc:"Mime"`

	FileCategoryId *string           `json:"fileCategoryId" doc:"File Category Id"`
	FileCategory   *FileCategoryItem `json:"fileCategory" doc:"File Category"`
}

type FileExtensionSave struct {
	schema.ModelId
	Ext  string `json:"ext" doc:"Ext"`
	Name string `json:"name" doc:"Name"`
	Mime string `json:"mime" doc:"Mime"`

	FileCategoryId *string `json:"fileCategoryId" doc:"File Category Id"`
}

type FileExtensionDetail struct {
	FileExtensionSave
	schema.ModelTime
	schema.ModelReserved

	FileCategory *FileCategoryItem `json:"fileCategory" doc:"File Category"`

	Files []FileItem `json:"files" doc:"Files"`
}

type FileExtensionSearchItem struct {
	schema.ModelId
	Ext  string `json:"ext" doc:"Ext"`
	Name string `json:"name" doc:"Name"`

	FileCategoryId *string                 `json:"fileCategoryId" doc:"File Category Id"`
	FileCategory   *FileCategorySearchItem `json:"fileCategory" doc:"File Category"`
}
