package schemas

import "majinyao.cn/my-app/backend/pkg/api/schema"

type FileGroupItem struct {
	schema.ModelId
	schema.ModelTime

	FileId string   `json:"fileId" doc:"File Id"`
	File   FileItem `json:"file" doc:"File"`

	GroupId string    `json:"groupId" doc:"Group Id"`
	Group   GroupItem `json:"group" doc:"Group"`

	NoCreate bool `json:"noCreate" doc:"No Create"`
	NoRead   bool `json:"noRead" doc:"No Read"`
	NoUpdate bool `json:"noUpdate" doc:"No Update"`
	NoDelete bool `json:"noDelete" doc:"No Delete"`
}

type FileGroupDetail struct {
	schema.ModelId
	schema.ModelTime

	FileId string   `json:"fileId" doc:"File Id"`
	File   FileItem `json:"file" doc:"File"`

	GroupId string    `json:"groupId" doc:"Group Id"`
	Group   GroupItem `json:"group" doc:"Group"`

	NoCreate bool `json:"noCreate" doc:"No Create"`
	NoRead   bool `json:"noRead" doc:"No Read"`
	NoUpdate bool `json:"noUpdate" doc:"No Update"`
	NoDelete bool `json:"noDelete" doc:"No Delete"`
}

type FileGroupSave struct {
	schema.ModelId

	FileId string `json:"fileId" doc:"File Id"`

	GroupId string `json:"groupId" doc:"Group Id"`

	NoCreate bool `json:"noCreate" doc:"No Create"`
	NoRead   bool `json:"noRead" doc:"No Read"`
	NoUpdate bool `json:"noUpdate" doc:"No Update"`
	NoDelete bool `json:"noDelete" doc:"No Delete"`
}
