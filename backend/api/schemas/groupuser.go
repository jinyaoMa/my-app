package schemas

import "majinyao.cn/my-app/backend/pkg/api/schema"

type GroupUserItem struct {
	schema.ModelId
	schema.ModelTime

	GroupId string    `json:"groupId" doc:"Group Id"`
	Group   GroupItem `json:"group" doc:"Group"`

	UserId string   `json:"userId" doc:"User Id"`
	User   UserItem `json:"user" doc:"User"`
}

type GroupUserDetail struct {
	schema.ModelId
	schema.ModelTime

	GroupId string    `json:"groupId" doc:"Group Id"`
	Group   GroupItem `json:"group" doc:"Group"`

	UserId string   `json:"userId" doc:"User Id"`
	User   UserItem `json:"user" doc:"User"`
}

type GroupUserSave struct {
	schema.ModelId

	GroupId string `json:"groupId" doc:"Group Id"`

	UserId string `json:"userId" doc:"User Id"`
}
