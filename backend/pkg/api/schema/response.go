package schema

type Response[T any] struct {
	Body struct {
		Success bool   `json:"success" example:"true" doc:"Has the action done successfully?"`
		Data    T      `json:"data" doc:"response payload"`
		Total   int64  `json:"total" doc:"if request is list query, total represents the count of the list w/o any user query conditions but conditions built-in; or represents affected rows for other requests"`
		Code    int    `json:"code" example:"0" doc:"0 => success; otherwise, fail"`
		Message string `json:"message" example:"hint message" doc:"hint message"`
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
