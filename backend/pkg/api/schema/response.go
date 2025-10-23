package schema

type Response[T any] struct {
	Body struct {
		Success bool   `json:"success" example:"true" doc:"Has the action done successfully?"`
		Data    T      `json:"data" doc:"Response Payload"`
		Total   int64  `json:"total" doc:"If request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; Or represents affected rows for other requests"`
		Code    int    `json:"code" example:"0" doc:"0 => success; Otherwise, fail"`
		Message string `json:"message" example:"hint message" doc:"Hint Message"`
	}
}

func Fail[T any](code int, msg string) *Response[T] {
	var r Response[T]
	r.Body.Code = code
	r.Body.Message = msg
	return &r
}

func Succeed[T any](data T, total int64) *Response[T] {
	var r Response[T]
	r.Body.Success = true
	r.Body.Data = data
	r.Body.Total = total
	return &r
}
