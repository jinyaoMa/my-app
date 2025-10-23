package schemas

import "majinyao.cn/my-app/backend/pkg/api/schema"

type FileUserItem struct {
	schema.ModelId
	schema.ModelTime

	FileId string   `json:"fileId" doc:"File Id"`
	File   FileItem `json:"file" doc:"File"`

	UserId string   `json:"userId" doc:"User Id"`
	User   UserItem `json:"user" doc:"User"`

	NoCreate bool `json:"noCreate" doc:"No Create"`
	NoRead   bool `json:"noRead" doc:"No Read"`
	NoUpdate bool `json:"noUpdate" doc:"No Update"`
	NoDelete bool `json:"noDelete" doc:"No Delete"`
	IsAvatar bool `json:"isAvatar" doc:"Is Used as Avatar or Not"`
}

type FileUserSave struct {
	schema.ModelId

	FileId string `json:"fileId" doc:"File Id"`

	UserId string `json:"userId" doc:"User Id"`

	NoCreate bool `json:"noCreate" doc:"No Create"`
	NoRead   bool `json:"noRead" doc:"No Read"`
	NoUpdate bool `json:"noUpdate" doc:"No Update"`
	NoDelete bool `json:"noDelete" doc:"No Delete"`
	IsAvatar bool `json:"isAvatar" doc:"Is Used as Avatar or Not"`
}

type FileUserDetail struct {
	FileUserSave
	schema.ModelTime

	File FileItem `json:"file" doc:"File"`

	User UserItem `json:"user" doc:"User"`
}

type FileUserSearchItem struct {
	schema.ModelId

	FileId string         `json:"fileId" doc:"File Id"`
	File   FileSearchItem `json:"file" doc:"File"`

	UserId string         `json:"userId" doc:"User Id"`
	User   UserSearchItem `json:"user" doc:"User"`

	NoCreate bool `json:"noCreate" doc:"No Create"`
	NoRead   bool `json:"noRead" doc:"No Read"`
	NoUpdate bool `json:"noUpdate" doc:"No Update"`
	NoDelete bool `json:"noDelete" doc:"No Delete"`
	IsAvatar bool `json:"isAvatar" doc:"Is Used as Avatar or Not"`
}
