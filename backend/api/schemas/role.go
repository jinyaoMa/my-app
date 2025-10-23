package schemas

import "majinyao.cn/my-app/backend/pkg/api/schema"

type RoleItem struct {
	schema.ModelId
	schema.ModelTime
	schema.ModelReserved
	Code        string `json:"code" doc:"Code"`
	Name        string `json:"name" doc:"Name"`
	Description string `json:"description" doc:"Description"`
}

type RoleSave struct {
	schema.ModelId
	Code        string `json:"code" doc:"Code"`
	Name        string `json:"name" doc:"Name"`
	Description string `json:"description" doc:"Description"`
}

type RoleDetail struct {
	RoleSave
	schema.ModelTime
	schema.ModelReserved

	RolePermissions []RolePermissionItem `json:"rolePermissions" doc:"Role Permissions"`
	Permissions     []PermissionItem     `json:"permissions" doc:"Permissions"`

	UserRoles []UserRoleItem `json:"userRoles" doc:"User Roles"`
	Users     []UserItem     `json:"users" doc:"Users"`

	GroupRoles []GroupRoleItem `json:"groupRoles" doc:"Group Roles"`
	Groups     []GroupItem     `json:"groups" doc:"Groups"`
}

type RoleSearchItem struct {
	schema.ModelId
	Code string `json:"code" doc:"Code"`
	Name string `json:"name" doc:"Name"`
}
