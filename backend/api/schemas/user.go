package schemas

import (
	"majinyao.cn/my-app/backend/pkg/api/schema"
	"majinyao.cn/my-app/backend/pkg/db/datatype"
)

type UserData struct {
	VisitorId string
	UserId    int64
}

func (u UserData) GetIdentity() string {
	return u.GetUserId().HexString() + "_" + u.VisitorId
}

func (u UserData) GetUserId() datatype.Id {
	return datatype.Id(u.UserId)
}

type UserItem struct {
	schema.EntityId
	schema.EntityTime
	Account     string `json:"account" doc:"Account"`
	Password    string `json:"password" doc:"Password"`
	Name        string `json:"name" doc:"Name"`
	Description string `json:"description" doc:"Description"`
}

type UserDetail struct {
	schema.EntityId
	schema.EntityTime
	Account     string `json:"account" doc:"Account"`
	Password    string `json:"password" doc:"Password"`
	Name        string `json:"name" doc:"Name"`
	Description string `json:"description" doc:"Description"`

	UserPasswords []UserPasswordItem `json:"userPasswords" doc:"User Passwords"`

	UserRoles []UserRoleItem `json:"userRoles" doc:"User Roles"`
	Roles     []RoleItem     `json:"roles" doc:"Roles"`

	GroupUsers []GroupUserItem `json:"groupUsers" doc:"Group Users"`
	Groups     []GroupItem     `json:"groups" doc:"Groups"`

	FileUsers []FileUserItem `json:"fileUsers" doc:"File Users"`
	Files     []FileItem     `json:"files" doc:"Files"`
}

type UserSave struct {
	schema.EntityId
	Account     string `json:"account" doc:"Account"`
	Password    string `json:"password" doc:"Password"`
	Name        string `json:"name" doc:"Name"`
	Description string `json:"description" doc:"Description"`
}
