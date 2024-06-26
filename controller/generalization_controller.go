/**
 * @Author: Nan
 * @Date: 2024/6/13 下午11:30
 */

package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"sweet-cms/form/request"
	"sweet-cms/form/response"
	"sweet-cms/service"
)

type GeneralizationController struct {
	generalizationService *service.GeneralizationService
	sysTableService       *service.SysTableService
}

func NewGeneralizationController(generalizationService *service.GeneralizationService, sysTableService *service.SysTableService) *GeneralizationController {
	return &GeneralizationController{
		generalizationService,
		sysTableService,
	}
}

func (gc *GeneralizationController) Query(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	resp := response.NewResponse()
	ctx.Set("response", resp)
	if err != nil {
		ctx.Error(err)
		return
	}
	table, err := gc.sysTableService.GetTableById(id)
	if err != nil {
		ctx.Error(err)
		return
	}
	if table.ID == 0 {
		e := &response.AdminError{
			Code:    http.StatusBadRequest,
			Message: "ID资源错误",
		}
		ctx.Error(e)
		return
	}
	var data request.Basic
	if err := ctx.ShouldBindQuery(&data); err != nil {
		e := &response.AdminError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		ctx.Error(e)
		return
	}
	data.TableCode = table.TableCode
	result, err := gc.generalizationService.Query(data, table)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp.SetData(result.Data).SetTotal(result.Total)
	return
}
