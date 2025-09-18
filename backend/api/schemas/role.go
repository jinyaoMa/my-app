package schemas

import "majinyao.cn/my-app/backend/pkg/api/schema"

type RoleItem struct {
	schema.EntityId
	schema.EntityTime
	Code        string `json:"code" doc:"Code"`
	Name        string `json:"name" doc:"Name"`
	Description string `json:"description" doc:"Description"`
}

type RoleDetail struct {
	schema.EntityId
	schema.EntityTime
	Code        string `json:"code" doc:"Code"`
	Name        string `json:"name" doc:"Name"`
	Description string `json:"description" doc:"Description"`

	RolePermissions []RolePermissionItem `json:"rolePermissions" doc:"Role Permissions"`
	Permissions     []PermissionItem     `json:"permissions" doc:"Permissions"`

	UserRoles []UserRoleItem `json:"userRoles" doc:"User Roles"`
	Users     []UserItem     `json:"users" doc:"Users"`

	GroupRoles []GroupRoleItem `json:"groupRoles" doc:"Group Roles"`
	Groups     []GroupItem     `json:"groups" doc:"Groups"`
}

type RoleSave struct {
	schema.EntityId
	Code        string `json:"code" doc:"Code"`
	Name        string `json:"name" doc:"Name"`
	Description string `json:"description" doc:"Description"`
}
