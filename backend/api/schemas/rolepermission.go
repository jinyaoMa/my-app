package schemas

import "majinyao.cn/my-app/backend/pkg/api/schema"

type RolePermissionItem struct {
	schema.ModelId
	schema.ModelTime
	schema.ModelReserved

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

type RolePermissionDetail struct {
	RolePermissionSave
	schema.ModelTime
	schema.ModelReserved

	Role RoleItem `json:"role" doc:"Role"`

	Permission PermissionItem `json:"permission" doc:"Permission"`
}

type RolePermissionSearchItem struct {
	schema.ModelId

	RoleId string         `json:"roleId" doc:"Role Id"`
	Role   RoleSearchItem `json:"role" doc:"Role"`

	PermissionId string               `json:"permissionId" doc:"Permission Id"`
	Permission   PermissionSearchItem `json:"permission" doc:"Permission"`
}
