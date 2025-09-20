package schemas

import "majinyao.cn/my-app/backend/pkg/api/schema"

type GroupRoleItem struct {
	schema.ModelId
	schema.ModelTime

	GroupId string    `json:"groupId" doc:"Group Id"`
	Group   GroupItem `json:"group" doc:"Group"`

	RoleId string   `json:"roleId" doc:"Role Id"`
	Role   RoleItem `json:"role" doc:"Role"`
}

type GroupRoleDetail struct {
	schema.ModelId
	schema.ModelTime

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
