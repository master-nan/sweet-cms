package response

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdminError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *AdminError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

type BufferedResponseWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

func (w *BufferedResponseWriter) Write(b []byte) (int, error) {
	w.Body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *BufferedResponseWriter) WriteString(s string) (int, error) {
	w.Body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

// Response 返回值参数
type Response struct {
	Success      bool        `json:"success"`
	Data         interface{} `json:"data,omitempty"`
	Message      string      `json:"message,omitempty"`
	Code         int         `json:"code,omitempty"`
	ErrorMessage string      `json:"error_message,omitempty"`
	ErrorCode    int         `json:"error_code,omitempty"`
	Total        int         `json:"total"`
}

func NewResponse() *Response {
	return &Response{
		Success: true,
		Code:    http.StatusOK,
		Message: "操作成功",
	}
}

func (r *Response) SetSuccess(success bool) *Response {
	r.Success = success
	return r
}

func (r *Response) SetData(data interface{}) *Response {
	r.Data = data
	return r
}

func (r *Response) SetTotal(total int) *Response {
	r.Total = total
	return r
}

func (r *Response) SetMessage(msg string) *Response {
	r.Message = msg
	return r
}

func (r *Response) SetCode(code int) *Response {
	r.Code = code
	return r
}
