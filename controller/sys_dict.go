/**
 * @Author: Nan
 * @Date: 2024/5/23 下午2:57
 */

package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"sweet-cms/form/request"
	"sweet-cms/middlewares"
	"sweet-cms/service"
)

type DictController struct {
	sysDictService *service.SysDictService
}

func NewDictController(sysDictService *service.SysDictService) *DictController {
	return &DictController{
		sysDictService: sysDictService,
	}
}

func (t *DictController) Get(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	resp := middlewares.NewResponse()
	ctx.Set("response", resp)
	if err != nil {
		resp.SetMsg(err.Error()).SetCode(http.StatusBadRequest)
		return
	}
	data, err := t.sysDictService.Get(id)
	if err != nil {
		resp.SetMsg(err.Error()).SetCode(http.StatusInternalServerError)
		return
	}
	resp.SetData(data)
	return
}

func (t *DictController) Query(ctx *gin.Context) {
	resp := middlewares.NewResponse()
	ctx.Set("response", resp)
	var data request.Basic
	if err := ctx.ShouldBindQuery(&data); err != nil {
		resp.SetCode(http.StatusInternalServerError).SetMsg(err.Error())
		return
	}
	result, err := t.sysDictService.Query(data)
	if err != nil {
		resp.SetCode(http.StatusInternalServerError).SetMsg(err.Error())
		return
	}
	resp.SetData(result.Data).SetTotal(result.Total)
	return
}

func (t *DictController) Insert(ctx *gin.Context) {
	resp := middlewares.NewResponse()
	ctx.Set("response", resp)
	return
}

func (t *DictController) Update(ctx *gin.Context) {
	resp := middlewares.NewResponse()
	ctx.Set("response", resp)
	return

}
func (t *DictController) Delete(ctx *gin.Context) {
	resp := middlewares.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		resp.SetCode(http.StatusInternalServerError).SetMsg(err.Error())
		return
	}
	err = t.sysDictService.Delete(id)
	if err != nil {
		resp.SetCode(http.StatusInternalServerError).SetMsg(err.Error())
		return
	}
	return

}

type DictItemController struct {
}

func NewDictItemController() *DictItemController {
	return &DictItemController{}
}

func (t *DictItemController) Get(ctx *gin.Context) {
	return
}

func (t *DictItemController) Query(ctx *gin.Context) {
	return
}
func (t *DictItemController) Insert(ctx *gin.Context) {

}

func (t *DictItemController) Update(ctx *gin.Context) {

}
func (t *DictItemController) Delete(ctx *gin.Context) {

}
