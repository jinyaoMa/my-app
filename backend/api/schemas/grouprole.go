package schemas

import "majinyao.cn/my-app/backend/pkg/api/schema"

type GroupRoleItem struct {
	schema.EntityId
	schema.EntityTime

	GroupId string    `json:"groupId" doc:"Group Id"`
	Group   GroupItem `json:"group" doc:"Group"`

	RoleId string   `json:"roleId" doc:"Role Id"`
	Role   RoleItem `json:"role" doc:"Role"`
}

type GroupRoleDetail struct {
	schema.EntityId
	schema.EntityTime

	GroupId string    `json:"groupId" doc:"Group Id"`
	Group   GroupItem `json:"group" doc:"Group"`

	RoleId string   `json:"roleId" doc:"Role Id"`
	Role   RoleItem `json:"role" doc:"Role"`
}

type GroupRoleSave struct {
	schema.EntityId

	GroupId string `json:"groupId" doc:"Group Id"`

	RoleId string `json:"roleId" doc:"Role Id"`
}
