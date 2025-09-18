package schemas

import "majinyao.cn/my-app/backend/pkg/api/schema"

type GroupUserItem struct {
	schema.EntityId
	schema.EntityTime

	GroupId string    `json:"groupId" doc:"Group Id"`
	Group   GroupItem `json:"group" doc:"Group"`

	UserId string   `json:"userId" doc:"User Id"`
	User   UserItem `json:"user" doc:"User"`
}

type GroupUserDetail struct {
	schema.EntityId
	schema.EntityTime

	GroupId string    `json:"groupId" doc:"Group Id"`
	Group   GroupItem `json:"group" doc:"Group"`

	UserId string   `json:"userId" doc:"User Id"`
	User   UserItem `json:"user" doc:"User"`
}

type GroupUserSave struct {
	schema.EntityId

	GroupId string `json:"groupId" doc:"Group Id"`

	UserId string `json:"userId" doc:"User Id"`
}
