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
	"sweet-cms/form/response"
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

// GetSysDictById 根据ID获取字典详情
// @Summary 字典详情
// @Description 根据ID获取字典详情
// @Tags 字典
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path int true "字典ID"
// @Success 200 {object} response.Response "请求成功"
// @Router /dict/id/{id} [get]
func (t *DictController) GetSysDictById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	resp := response.NewResponse()
	ctx.Set("response", resp)
	if err != nil {
		resp.SetErrorMessage(err.Error()).SetErrorCode(http.StatusBadRequest)
		return
	}
	data, err := t.sysDictService.GetSysDictById(id)
	if err != nil {
		resp.SetErrorMessage(err.Error()).SetErrorCode(http.StatusInternalServerError)
		return
	}
	resp.SetData(data)
	return
}

// GetSysDictByCode 根据CODE获取字典详情
// @Summary 字典详情
// @Description 根据CODE获取字典详情
// @Tags 字典
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param code path string true "字典CODE"
// @Success 200 {object} response.Response "请求成功"
// @Router /dict/code/{code} [get]
func (t *DictController) GetSysDictByCode(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	code := ctx.Param("code")
	data, err := t.sysDictService.GetSysDictByCode(code)
	if err != nil {
		resp.SetErrorMessage(err.Error()).SetErrorCode(http.StatusInternalServerError)
		return
	}
	resp.SetData(data)
	return
}

// QuerySysDict 字典列表
// @Summary 字典列表
// @Description 根据查询条件查询字段列表
// @Tags 字典
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param q query request.Basic false "请求参数"
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "参数错误"
// @Failure 500 {object} response.Response  "内部错误"
// @Router /dict/query [get]
func (t *DictController) QuerySysDict(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	var data request.Basic
	if err := ctx.ShouldBindQuery(&data); err != nil {
		resp.SetErrorCode(http.StatusInternalServerError).SetErrorMessage(err.Error())
		return
	}
	result, err := t.sysDictService.GetSysDictList(data)
	if err != nil {
		resp.SetErrorCode(http.StatusInternalServerError).SetErrorMessage(err.Error())
		return
	}
	resp.SetData(result.Data).SetTotal(result.Total)
	return
}

func (t *DictController) InsertSysDict(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	var dictCreateReq request.DictCreateReq
	err := ctx.ShouldBindJSON(&dictCreateReq)
	if err != nil {
		resp.SetErrorCode(http.StatusBadRequest).SetErrorMessage(err.Error())
		return
	}
	err = t.sysDictService.InsertSysDict(dictCreateReq)
	if err != nil {
		resp.SetErrorCode(http.StatusBadRequest).SetErrorMessage(err.Error())
		return
	}
	return
}

// UpdateSysDict 更新字典
// @Summary 更新字典
// @Description 根据ID更新字典信息
// @Tags 字典
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path int true "字典ID"
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "请求错误"
// @Failure 500 {object} response.Response "内部错误"
// @Router /dict/{id} [put]
func (t *DictController) UpdateSysDict(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		resp.SetErrorCode(http.StatusBadRequest).SetErrorMessage(err.Error())
		return
	}
	var dictUpdateReq request.DictUpdateReq
	dictUpdateReq.ID = id
	err = ctx.ShouldBindJSON(&dictUpdateReq)
	if err != nil {
		resp.SetErrorCode(http.StatusBadRequest).SetErrorMessage(err.Error())
		return
	}
	err = t.sysDictService.UpdateSysDict(dictUpdateReq)
	if err != nil {
		resp.SetErrorCode(http.StatusBadRequest).SetErrorMessage(err.Error())
		return
	}
	return
}

func (t *DictController) DeleteSysDictById(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		resp.SetErrorCode(http.StatusInternalServerError).SetErrorMessage(err.Error())
		return
	}
	err = t.sysDictService.DeleteSysDictById(id)
	if err != nil {
		resp.SetErrorCode(http.StatusInternalServerError).SetErrorMessage(err.Error())
		return
	}
	return

}

func (t *DictController) GetSysDictItemById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	resp := response.NewResponse()
	ctx.Set("response", resp)
	if err != nil {
		resp.SetErrorMessage(err.Error()).SetErrorCode(http.StatusBadRequest)
		return
	}
	data, err := t.sysDictService.GetSysDictItemById(id)
	if err != nil {
		resp.SetErrorMessage(err.Error()).SetErrorCode(http.StatusInternalServerError)
		return
	}
	resp.SetData(data)
	return
}

func (t *DictController) GetSysDictItemsByDictId(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		resp.SetErrorMessage(err.Error()).SetErrorCode(http.StatusBadRequest)
		return
	}
	result, err := t.sysDictService.GetSysDictItemsByDictId(id)
	if err != nil {
		resp.SetErrorCode(http.StatusInternalServerError).SetErrorMessage(err.Error())
		return
	}
	resp.SetData(result)
	return
}

func (t *DictController) InsertSysDictItem(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	var dictItemCreateReq request.DictItemCreateReq
	err := ctx.ShouldBindJSON(&dictItemCreateReq)
	if err != nil {
		resp.SetErrorCode(http.StatusBadRequest).SetErrorMessage(err.Error())
		return
	}
	err = t.sysDictService.InsertSysDictItem(dictItemCreateReq)
	if err != nil {
		resp.SetErrorCode(http.StatusBadRequest).SetErrorMessage(err.Error())
		return
	}
	return
}

func (t *DictController) UpdateSysDictItem(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	var dictItemUpdateReq request.DictItemUpdateReq
	err := ctx.ShouldBindJSON(&dictItemUpdateReq)
	if err != nil {
		resp.SetErrorCode(http.StatusBadRequest).SetErrorMessage(err.Error())
		return
	}
	err = t.sysDictService.UpdateSysDictItem(dictItemUpdateReq)
	if err != nil {
		resp.SetErrorCode(http.StatusBadRequest).SetErrorMessage(err.Error())
		return
	}
	return
}

func (t *DictController) DeleteSysDictItemById(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		resp.SetErrorCode(http.StatusInternalServerError).SetErrorMessage(err.Error())
		return
	}
	err = t.sysDictService.DeleteSysDictItemById(id)
	if err != nil {
		resp.SetErrorCode(http.StatusInternalServerError).SetErrorMessage(err.Error())
		return
	}
	return
}
