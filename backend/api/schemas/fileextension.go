package schemas

import "majinyao.cn/my-app/backend/pkg/api/schema"

type FileExtensionItem struct {
	schema.EntityId
	schema.EntityTime
	Ext  string `json:"ext" doc:"Ext"`
	Name string `json:"name" doc:"Name"`
	Mime string `json:"mime" doc:"Mime"`

	FileCategoryId *string           `json:"fileCategoryId" doc:"File Category Id"`
	FileCategory   *FileCategoryItem `json:"fileCategory" doc:"File Category"`
}

type FileExtensionDetail struct {
	schema.EntityId
	schema.EntityTime
	Ext  string `json:"ext" doc:"Ext"`
	Name string `json:"name" doc:"Name"`
	Mime string `json:"mime" doc:"Mime"`

	FileCategoryId *string           `json:"fileCategoryId" doc:"File Category Id"`
	FileCategory   *FileCategoryItem `json:"fileCategory" doc:"File Category"`

	Files []FileItem `json:"files" doc:"Files"`
}

type FileExtensionSave struct {
	schema.EntityId
	Ext  string `json:"ext" doc:"Ext"`
	Name string `json:"name" doc:"Name"`
	Mime string `json:"mime" doc:"Mime"`

	FileCategoryId *string `json:"fileCategoryId" doc:"File Category Id"`
}
