package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RespData struct {
	Data  any    `json:"data"`
	Msg   string `json:"msg"`
	Total int    `json:"total"`
	Code  int    `json:"code"`
	c     *gin.Context
}

func NewRespData(ctx *gin.Context) *RespData {
	return &RespData{Msg: "", Code: http.StatusOK, c: ctx}
}

func (r *RespData) SetData(data any) *RespData {
	r.Data = data
	return r
}

func (r *RespData) SetTotal(total int) *RespData {
	r.Total = total
	return r
}

func (r *RespData) SetMsg(msg string) *RespData {
	r.Msg = msg
	return r
}

func (r *RespData) SetCode(code int) *RespData {
	r.Code = code
	return r
}

func (r *RespData) ReturnJson() {
	r.c.JSON(r.Code, r)
}

func (r *RespData) AbortStatusJson() {
	r.c.Abort()
	r.c.JSON(r.Code, r)
}
