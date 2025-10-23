package schema

import "majinyao.cn/my-app/backend/pkg/db/crud"

type CrudSearchRequest struct {
	Keyword  string        `json:"keyword" doc:"Search Keyword"`
	Fields   []string      `json:"fields" doc:"Search Fields"`
	Limit    int           `json:"limit" required:"true" minimum:"1" maximum:"20" doc:"Search Limit"`
	Exact    bool          `json:"exact" required:"false" doc:"Is Exact Search"`
	Includes crud.Includes `json:"includes" required:"false" doc:"Included Associations"`
}
