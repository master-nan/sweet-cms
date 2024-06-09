package response

import "fmt"

type AdminError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *AdminError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

// Response 返回值参数
type Response struct {
	Success      bool        `json:"success"`
	Data         interface{} `json:"data,omitempty"`
	ErrorMessage string      `json:"error_message,omitempty"`
	ErrorCode    int         `json:"error_code,omitempty"`
	Total        int         `json:"total"`
}

func NewResponse() *Response {
	return &Response{}
}

func (r *Response) SetData(data interface{}) *Response {
	r.Data = data
	return r
}

func (r *Response) SetTotal(total int) *Response {
	r.Total = total
	return r
}

func (r *Response) SetErrorMessage(msg string) *Response {
	r.ErrorMessage = msg
	return r
}

func (r *Response) SetErrorCode(code int) *Response {
	r.ErrorCode = code
	return r
}
