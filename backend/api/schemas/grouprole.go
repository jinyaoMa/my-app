package schemas

import "majinyao.cn/my-app/backend/pkg/api/schema"

type GroupRoleItem struct {
	schema.ModelId
	schema.ModelTime
	schema.ModelReserved

	GroupId string    `json:"groupId" doc:"Group Id"`
	Group   GroupItem `json:"group" doc:"Group"`

	RoleId string   `json:"roleId" doc:"Role Id"`
	Role   RoleItem `json:"role" doc:"Role"`
}

type GroupRoleSave struct {
	schema.ModelId

	GroupId string `json:"groupId" doc:"Group Id"`

	RoleId string `json:"roleId" doc:"Role Id"`
}

type GroupRoleDetail struct {
	GroupRoleSave
	schema.ModelTime
	schema.ModelReserved

	Group GroupItem `json:"group" doc:"Group"`

	Role RoleItem `json:"role" doc:"Role"`
}

type GroupRoleSearchItem struct {
	schema.ModelId

	GroupId string          `json:"groupId" doc:"Group Id"`
	Group   GroupSearchItem `json:"group" doc:"Group"`

	RoleId string         `json:"roleId" doc:"Role Id"`
	Role   RoleSearchItem `json:"role" doc:"Role"`
}
