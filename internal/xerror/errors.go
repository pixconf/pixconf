package xerror

import "encoding/json"

type Response struct {
	Code   int       `json:"code"`
	Errors []Message `json:"errors"`
}

type Message struct {
	Message string `json:"message"`
}

func ErrorSingle(code int, message string) Response {
	return Response{
		Code: code,
		Errors: []Message{
			{Message: message},
		},
	}
}

func (r *Response) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
