package schemas

import "majinyao.cn/my-app/backend/pkg/api/schema"

type RolePermissionItem struct {
	schema.ModelId
	schema.ModelTime

	RoleId string   `json:"roleId" doc:"Role Id"`
	Role   RoleItem `json:"role" doc:"Role"`

	PermissionId string         `json:"permissionId" doc:"Permission Id"`
	Permission   PermissionItem `json:"permission" doc:"Permission"`
}

type RolePermissionDetail struct {
	schema.ModelId
	schema.ModelTime

	RoleId string   `json:"roleId" doc:"Role Id"`
	Role   RoleItem `json:"role" doc:"Role"`

	PermissionId string         `json:"permissionId" doc:"Permission Id"`
	Permission   PermissionItem `json:"permission" doc:"Permission"`
}

type RolePermissionSave struct {
	schema.ModelId

	RoleId string `json:"roleId" doc:"Role Id"`

	PermissionId string `json:"permissionId" doc:"Permission Id"`
}
