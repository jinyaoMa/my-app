package schemas

import (
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/api/schema"
)

type FileItem struct {
	schema.EntityId
	schema.EntityTime
	Oid   string `json:"oid" doc:"Oid"`
	IsDir bool   `json:"isDir" doc:"Is Directory or Not"`
	Name  string `json:"name" doc:"Name"`
	Size  uint64 `json:"size" doc:"Size"`

	Status   entity.FileStatus `json:"status" doc:"Status: 1-Uploading, 2-Persisted, 3-NotFound, 4-Forbidden"`
	Readonly bool              `json:"readonly" doc:"Is Read-only or Not"`
	Hidden   bool              `json:"hidden" doc:"Is Hidden or Visible"`

	FileExtensionId *string            `json:"fileExtensionId" doc:"File Extension Id"`
	FileExtension   *FileExtensionItem `json:"fileExtension" doc:"File Extension"`
}

type FileDetail struct {
	schema.EntityId
	schema.EntityTime
	Oid   string `json:"oid" doc:"Oid"`
	IsDir bool   `json:"isDir" doc:"Is Directory or Not"`
	Name  string `json:"name" doc:"Name"`
	Size  uint64 `json:"size" doc:"Size"`

	Status   entity.FileStatus `json:"status" doc:"Status: 1-Uploading, 2-Persisted, 3-NotFound, 4-Forbidden"`
	Readonly bool              `json:"readonly" doc:"Is Read-only or Not"`
	Hidden   bool              `json:"hidden" doc:"Is Hidden or Visible"`

	FileExtensionId *string            `json:"fileExtensionId" doc:"File Extension Id"`
	FileExtension   *FileExtensionItem `json:"fileExtension" doc:"File Extension"`

	FileUsers []FileUserItem `json:"fileUsers" doc:"File Users"`
	Users     []UserItem     `json:"users" doc:"Users"`

	FileGroups []FileGroupItem `json:"fileGroups" doc:"File Groups"`
	Groups     []GroupItem     `json:"groups" doc:"Groups"`
}

type FileSave struct {
	schema.EntityId
	Oid   string `json:"oid" doc:"Oid"`
	IsDir bool   `json:"isDir" doc:"Is Directory or Not"`
	Name  string `json:"name" doc:"Name"`
	Size  uint64 `json:"size" doc:"Size"`

	Status   entity.FileStatus `json:"status" doc:"Status: 1-Uploading, 2-Persisted, 3-NotFound, 4-Forbidden"`
	Readonly bool              `json:"readonly" doc:"Is Read-only or Not"`
	Hidden   bool              `json:"hidden" doc:"Is Hidden or Visible"`

	FileExtensionId *string `json:"fileExtensionId" doc:"File Extension Id"`
}
