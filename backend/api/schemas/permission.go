package schemas

import "majinyao.cn/my-app/backend/pkg/api/schema"

type PermissionItem struct {
	schema.ModelId
	schema.ModelTime
	Code        string `json:"code" doc:"Code"`
	Name        string `json:"name" doc:"Name"`
	Description string `json:"description" doc:"Description"`
	// Flag        []byte `json:"flag" doc:"Flag"`
}

type PermissionDetail struct {
	schema.ModelId
	schema.ModelTime
	Code        string `json:"code" doc:"Code"`
	Name        string `json:"name" doc:"Name"`
	Description string `json:"description" doc:"Description"`
	// Flag        []byte `json:"flag" doc:"Flag"`

	RolePermissions []RolePermissionItem `json:"rolePermissions" doc:"Role Permissions"`
	Roles           []RoleItem           `json:"roles" doc:"Roles"`
}

type PermissionSave struct {
	schema.ModelId
	Code        string `json:"code" doc:"Code"`
	Name        string `json:"name" doc:"Name"`
	Description string `json:"description" doc:"Description"`
	// Flag        []byte `json:"flag" doc:"Flag"`
}
