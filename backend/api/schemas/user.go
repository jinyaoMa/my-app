package schemas

import (
	"time"

	"majinyao.cn/my-app/backend/pkg/api/schema"
)

type UserItem struct {
	schema.ModelId
	schema.ModelTime
	schema.ModelReserved
	Account                    string    `json:"account" doc:"Account"`
	Name                       string    `json:"name" doc:"Name"`
	TemporaryPassword          string    `json:"temporaryPassword" doc:"Temporary Password"`
	TemporaryPasswordExpiredAt time.Time `json:"temporaryPasswordExpiredAt" doc:"Temporary Password Expired At"`
}

type UserSave struct {
	schema.ModelId
	Account                    string    `json:"account" doc:"Account"`
	Password                   string    `json:"password" doc:"Password"`
	Name                       string    `json:"name" doc:"Name"`
	Description                string    `json:"description" doc:"Description"`
	TemporaryPassword          string    `json:"temporaryPassword" doc:"Temporary Password"`
	TemporaryPasswordExpiredAt time.Time `json:"temporaryPasswordExpiredAt" doc:"Temporary Password Expired At"`
}

type UserDetail struct {
	UserSave
	schema.ModelTime
	schema.ModelReserved

	UserPasswords []UserPasswordItem `json:"userPasswords" doc:"User Passwords"`

	UserRoles []UserRoleItem `json:"userRoles" doc:"User Roles"`
	Roles     []RoleItem     `json:"roles" doc:"Roles"`

	GroupUsers []GroupUserItem `json:"groupUsers" doc:"Group Users"`
	Groups     []GroupItem     `json:"groups" doc:"Groups"`

	FileUsers []FileUserItem `json:"fileUsers" doc:"File Users"`
	Files     []FileItem     `json:"files" doc:"Files"`
}

type UserSearchItem struct {
	schema.ModelId
	Account string `json:"account" doc:"Account"`
	Name    string `json:"name" doc:"Name"`
}
