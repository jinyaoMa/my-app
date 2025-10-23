package schemas

import "majinyao.cn/my-app/backend/pkg/api/schema"

type GroupItem struct {
	schema.ModelId
	schema.ModelTime
	schema.ModelReserved
	Code        string `json:"code" doc:"Code"`
	Name        string `json:"name" doc:"Name"`
	Description string `json:"description" doc:"Description"`
}

type GroupSave struct {
	schema.ModelId
	Code        string `json:"code" doc:"Code"`
	Name        string `json:"name" doc:"Name"`
	Description string `json:"description" doc:"Description"`
}

type GroupDetail struct {
	GroupSave
	schema.ModelTime
	schema.ModelReserved

	GroupUsers []GroupUserItem `json:"groupUsers" doc:"Group Users"`
	Users      []UserItem      `json:"users" doc:"Users"`

	GroupRoles []GroupRoleItem `json:"groupRoles" doc:"Group Roles"`
	Roles      []RoleItem      `json:"roles" doc:"Roles"`

	FileGroups []FileGroupItem `json:"fileGroups" doc:"File Groups"`
	Files      []FileItem      `json:"files" doc:"Files"`
}

type GroupSearchItem struct {
	schema.ModelId
	Code string `json:"code" doc:"Code"`
	Name string `json:"name" doc:"Name"`
}
