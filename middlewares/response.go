/**
 * @Author: Nan
 * @Date: 2024/5/25 下午4:22
 */

package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Data  interface{} `json:"data"`
	Msg   string      `json:"msg"`
	Total int         `json:"total"`
	Code  int         `json:"code"`
}

func JSONResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if resp, exists := c.Get("response"); exists {
			response := resp.(*Response)
			c.JSON(response.Code, resp)
			c.Abort()
		}
	}
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
