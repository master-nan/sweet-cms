/**
 * @Author: Nan
 * @Date: 2024/5/23 下午2:57
 */

package controller

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"net/http"
	"strconv"
	"sweet-cms/form/request"
	"sweet-cms/form/response"
	"sweet-cms/service"
	"sweet-cms/utils"
)

type DictController struct {
	sysDictService *service.SysDictService
	translators    map[string]ut.Translator
}

func NewDictController(sysDictService *service.SysDictService, translators map[string]ut.Translator) *DictController {
	return &DictController{
		sysDictService,
		translators,
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
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	data, err := t.sysDictService.GetSysDictById(id)
	if err != nil {
		_ = ctx.Error(err)
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
	code := utils.SanitizeInput(ctx.Param("code"))
	data, err := t.sysDictService.GetSysDictByCode(code)
	if err != nil {
		e := &response.AdminError{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: err.Error(),
		}
		_ = ctx.Error(e)
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
// @Param q body request.Basic false "请求参数"
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "参数错误"
// @Failure 500 {object} response.Response  "内部错误"
// @Router /dict/query [get]
func (t *DictController) QuerySysDict(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	var data request.Basic
	if err := ctx.ShouldBindQuery(&data); err != nil {
		e := &response.AdminError{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: err.Error(),
		}
		_ = ctx.Error(e)
		return
	}
	result, err := t.sysDictService.GetSysDictList(data)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	resp.SetData(result.Data).SetTotal(result.Total)
	return
}

// CreateSysDict 新增字典
// @Summary 新增字典
// @Description 新增字典主体
// @Tags 字典
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param b body request.DictCreateReq  true "请求参数"
// @Success 200 {object} response.Response "请求成功"
// @Failure 400 {object} response.Response "参数错误"
// @Failure 500 {object} response.Response  "内部错误"
// @Router /dict [post]
func (t *DictController) CreateSysDict(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	var data request.DictCreateReq
	translator, _ := t.translators["zh"]
	err := utils.ValidatorBody[request.DictCreateReq](ctx, &data, translator)
	err = t.sysDictService.CreateSysDict(ctx, data)
	if err != nil {
		_ = ctx.Error(err)
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
		e := &response.AdminError{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: err.Error(),
		}
		_ = ctx.Error(e)
		return
	}
	var data request.DictUpdateReq
	data.Id = id
	translator, _ := t.translators["zh"]
	err = utils.ValidatorBody[request.DictUpdateReq](ctx, &data, translator)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = t.sysDictService.UpdateSysDict(ctx, data)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	return
}

func (t *DictController) DeleteSysDictById(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = t.sysDictService.DeleteSysDictById(ctx, id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	return

}

func (t *DictController) GetSysDictItemsByDictId(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	result, err := t.sysDictService.GetSysDictItemsByDictId(id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	resp.SetData(result)
	return
}

func (t *DictController) GetSysDictItemById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	resp := response.NewResponse()
	ctx.Set("response", resp)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	data, err := t.sysDictService.GetSysDictItemById(id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	resp.SetData(data)
	return
}

func (t *DictController) CreateSysDictItem(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	var data request.DictItemCreateReq
	translator, _ := t.translators["zh"]
	err := utils.ValidatorBody[request.DictItemCreateReq](ctx, &data, translator)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = t.sysDictService.CreateSysDictItem(ctx, data)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	return
}

func (t *DictController) UpdateSysDictItem(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		e := &response.AdminError{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: err.Error(),
		}
		_ = ctx.Error(e)
		return
	}
	translator, _ := t.translators["zh"]
	var data request.DictItemUpdateReq
	data.Id = id
	err = utils.ValidatorBody[request.DictItemUpdateReq](ctx, &data, translator)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = t.sysDictService.UpdateSysDictItem(ctx, data)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	return
}

func (t *DictController) DeleteSysDictItemById(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = t.sysDictService.DeleteSysDictItemById(ctx, id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	return
}
