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

type GroupUserSave struct {
	schema.ModelId

	GroupId string `json:"groupId" doc:"Group Id"`

	UserId string `json:"userId" doc:"User Id"`
}

type GroupUserDetail struct {
	GroupUserSave
	schema.ModelTime

	Group GroupItem `json:"group" doc:"Group"`

	User UserItem `json:"user" doc:"User"`
}

type GroupUserSearchItem struct {
	schema.ModelId

	GroupId string          `json:"groupId" doc:"Group Id"`
	Group   GroupSearchItem `json:"group" doc:"Group"`

	UserId string         `json:"userId" doc:"User Id"`
	User   UserSearchItem `json:"user" doc:"User"`
}
