package ginUtils

type APIError struct {
	Code int
	Msg  string
}

func (e *APIError) Error() string {
	return e.Msg
}
