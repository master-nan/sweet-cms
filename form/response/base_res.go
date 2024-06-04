package response

import (
	"net/http"
)

// Response 返回值参数
type Response struct {
	Data  interface{} `json:"data"`
	Msg   string      `json:"msg"`
	Total int         `json:"total"`
	Code  int         `json:"code"`
}

func NewResponse() *Response {
	return &Response{Code: http.StatusOK, Msg: ""}
}

func (r *Response) SetData(data interface{}) *Response {
	r.Data = data
	return r
}

func (r *Response) SetTotal(total int) *Response {
	r.Total = total
	return r
}

func (r *Response) SetMsg(msg string) *Response {
	r.Msg = msg
	return r
}

func (r *Response) SetCode(code int) *Response {
	r.Code = code
	return r
}
