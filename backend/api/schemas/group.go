package schemas

import "majinyao.cn/my-app/backend/pkg/api/schema"

type GroupItem struct {
	schema.EntityId
	schema.EntityTime
	Code        string `json:"code" doc:"Code"`
	Name        string `json:"name" doc:"Name"`
	Description string `json:"description" doc:"Description"`
}

type GroupDetail struct {
	schema.EntityId
	schema.EntityTime
	Code        string `json:"code" doc:"Code"`
	Name        string `json:"name" doc:"Name"`
	Description string `json:"description" doc:"Description"`

	GroupUsers []GroupUserItem `json:"groupUsers" doc:"Group Users"`
	Users      []UserItem      `json:"users" doc:"Users"`

	GroupRoles []GroupRoleItem `json:"groupRoles" doc:"Group Roles"`
	Roles      []RoleItem      `json:"roles" doc:"Roles"`

	FileGroups []FileGroupItem `json:"fileGroups" doc:"File Groups"`
	Files      []FileItem      `json:"files" doc:"Files"`
}

type GroupSave struct {
	schema.EntityId
	Code        string `json:"code" doc:"Code"`
	Name        string `json:"name" doc:"Name"`
	Description string `json:"description" doc:"Description"`
}
