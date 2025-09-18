package schemas

import "majinyao.cn/my-app/backend/pkg/api/schema"

type RolePermissionItem struct {
	schema.EntityId
	schema.EntityTime

	RoleId string   `json:"roleId" doc:"Role Id"`
	Role   RoleItem `json:"role" doc:"Role"`

	PermissionId string         `json:"permissionId" doc:"Permission Id"`
	Permission   PermissionItem `json:"permission" doc:"Permission"`
}

type RolePermissionDetail struct {
	schema.EntityId
	schema.EntityTime

	RoleId string   `json:"roleId" doc:"Role Id"`
	Role   RoleItem `json:"role" doc:"Role"`

	PermissionId string         `json:"permissionId" doc:"Permission Id"`
	Permission   PermissionItem `json:"permission" doc:"Permission"`
}

type RolePermissionSave struct {
	schema.EntityId

	RoleId string `json:"roleId" doc:"Role Id"`

	PermissionId string `json:"permissionId" doc:"Permission Id"`
}
