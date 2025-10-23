package schemas

import (
	"majinyao.cn/my-app/backend/internal/entity"
	"majinyao.cn/my-app/backend/pkg/api/schema"
)

type FileItem struct {
	schema.ModelId
	schema.ModelTime
	Oid   string `json:"oid" doc:"Oid"`
	IsDir bool   `json:"isDir" doc:"Is Directory or Not"`
	Name  string `json:"name" doc:"Name"`
	Size  int64  `json:"size" doc:"Size"`

	Status   entity.FileStatus `json:"status" doc:"Status: uploading, persisted, lost, forbidden"`
	Readonly bool              `json:"readonly" doc:"Is Read-only or Not"`
	Hidden   bool              `json:"hidden" doc:"Is Hidden or Visible"`

	FileExtensionId *string            `json:"fileExtensionId" doc:"File Extension Id"`
	FileExtension   *FileExtensionItem `json:"fileExtension" doc:"File Extension"`
}

type FileSave struct {
	schema.ModelId
	Oid       string   `json:"oid" doc:"Oid"`
	IsDir     bool     `json:"isDir" doc:"Is Directory or Not"`
	Name      string   `json:"name" doc:"Name"`
	Checksums []string `json:"checksums" doc:"Checksums"`
	Size      int64    `json:"size" doc:"Size"`

	Status   entity.FileStatus `json:"status" doc:"Status: uploading, persisted, lost, forbidden"`
	Readonly bool              `json:"readonly" doc:"Is Read-only or Not"`
	Hidden   bool              `json:"hidden" doc:"Is Hidden or Visible"`

	FileExtensionId *string `json:"fileExtensionId" doc:"File Extension Id"`
}

type FileDetail struct {
	FileSave
	schema.ModelTime

	FileExtension *FileExtensionItem `json:"fileExtension" doc:"File Extension"`

	FileUsers []FileUserItem `json:"fileUsers" doc:"File Users"`
	Users     []UserItem     `json:"users" doc:"Users"`

	FileGroups []FileGroupItem `json:"fileGroups" doc:"File Groups"`
	Groups     []GroupItem     `json:"groups" doc:"Groups"`
}

type FileSearchItem struct {
	schema.ModelId
	IsDir bool   `json:"isDir" doc:"Is Directory or Not"`
	Name  string `json:"name" doc:"Name"`
	Size  int64  `json:"size" doc:"Size"`

	Status   entity.FileStatus `json:"status" doc:"Status: uploading, persisted, lost, forbidden"`
	Readonly bool              `json:"readonly" doc:"Is Read-only or Not"`

	FileExtensionId *string                  `json:"fileExtensionId" doc:"File Extension Id"`
	FileExtension   *FileExtensionSearchItem `json:"fileExtension" doc:"File Extension"`
}
