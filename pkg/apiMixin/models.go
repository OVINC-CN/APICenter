package apiMixin

type APIError struct {
	Code int
	Msg  string
}

func (e *APIError) Error() string {
	return e.Msg
}

type ResponseModel struct {
	TraceID string      `json:"trace_id"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
