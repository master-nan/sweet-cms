/**
 * @Author: Nan
 * @Date: 2024/5/23 下午2:57
 */

package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sweet-cms/form/request"
	"sweet-cms/form/response"
	"sweet-cms/service"
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
	id, err := strconv.Atoi(ctx.Param("id"))
	resp := response.NewResponse()
	ctx.Set("response", resp)
	if err != nil {
		ctx.Error(err)
		return
	}
	data, err := t.sysDictService.GetSysDictById(id)
	if err != nil {
		ctx.Error(err)
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
		e := &response.AdminError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		ctx.Error(e)
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
		e := &response.AdminError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		ctx.Error(e)
		return
	}
	result, err := t.sysDictService.GetSysDictList(data)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp.SetData(result.Data).SetTotal(result.Total)
	return
}

// InsertSysDict 新增字典
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
func (t *DictController) InsertSysDict(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	var dictCreateReq request.DictCreateReq
	translator, _ := t.translators["zh"]
	err := ctx.ShouldBindBodyWith(&dictCreateReq, binding.JSON)
	if err != nil {
		if err == io.EOF {
			// 客户端请求体为空
			e := &response.AdminError{
				Code:    http.StatusBadRequest,
				Message: "请求参数错误",
			}
			ctx.Error(e)
			return
		}
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			// 如果是验证错误，则翻译错误信息
			var errorMessages []string
			for _, e := range ve {
				errMsg := e.Translate(translator)
				errorMessages = append(errorMessages, errMsg)
			}
			e := &response.AdminError{
				Code:    http.StatusBadRequest,
				Message: strings.Join(errorMessages, ","),
			}
			ctx.Error(e)
			return
		}
		ctx.Error(err)
		return
	}
	err = t.sysDictService.InsertSysDict(dictCreateReq)
	if err != nil {
		ctx.Error(err)
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
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		ctx.Error(e)
		return
	}
	var dictUpdateReq request.DictUpdateReq
	dictUpdateReq.ID = id
	err = ctx.ShouldBindBodyWith(&dictUpdateReq, binding.JSON)
	translator, _ := t.translators["zh"]
	if err != nil {
		if err == io.EOF {
			// 客户端请求体为空
			e := &response.AdminError{
				Code:    http.StatusBadRequest,
				Message: "请求参数错误",
			}
			ctx.Error(e)
			return
		}
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			// 如果是验证错误，则翻译错误信息
			var errorMessages []string
			for _, e := range ve {
				errMsg := e.Translate(translator)
				errorMessages = append(errorMessages, errMsg)
			}
			e := &response.AdminError{
				Code:    http.StatusBadRequest,
				Message: strings.Join(errorMessages, ","),
			}
			ctx.Error(e)
			return
		}
		ctx.Error(err)
		return
	}
	err = t.sysDictService.UpdateSysDict(dictUpdateReq)
	if err != nil {
		ctx.Error(err)
		return
	}
	return
}

func (t *DictController) DeleteSysDictById(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		return
	}
	err = t.sysDictService.DeleteSysDictById(id)
	if err != nil {
		ctx.Error(err)
		return
	}
	return

}

func (t *DictController) GetSysDictItemById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	resp := response.NewResponse()
	ctx.Set("response", resp)
	if err != nil {
		ctx.Error(err)
		return
	}
	data, err := t.sysDictService.GetSysDictItemById(id)
	if err != nil {
		ctx.Error(err)
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
		ctx.Error(err)
		return
	}
	result, err := t.sysDictService.GetSysDictItemsByDictId(id)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp.SetData(result)
	return
}

func (t *DictController) InsertSysDictItem(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	var dictItemCreateReq request.DictItemCreateReq
	err := ctx.ShouldBindBodyWith(&dictItemCreateReq, binding.JSON)
	translator, _ := t.translators["zh"]
	if err != nil {
		if err == io.EOF {
			// 客户端请求体为空
			e := &response.AdminError{
				Code:    http.StatusBadRequest,
				Message: "请求参数错误",
			}
			ctx.Error(e)
			return
		}
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			// 如果是验证错误，则翻译错误信息
			var errorMessages []string
			for _, e := range ve {
				errMsg := e.Translate(translator)
				errorMessages = append(errorMessages, errMsg)
			}
			e := &response.AdminError{
				Code:    http.StatusBadRequest,
				Message: strings.Join(errorMessages, ","),
			}
			ctx.Error(e)
			return
		}
		ctx.Error(err)
		return
	}
	err = t.sysDictService.InsertSysDictItem(dictItemCreateReq)
	if err != nil {
		ctx.Error(err)
		return
	}
	return
}

func (t *DictController) UpdateSysDictItem(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	translator, _ := t.translators["zh"]
	var dictItemUpdateReq request.DictItemUpdateReq
	err := ctx.ShouldBindBodyWith(&dictItemUpdateReq, binding.JSON)
	if err != nil {
		if err == io.EOF {
			// 客户端请求体为空
			e := &response.AdminError{
				Code:    http.StatusBadRequest,
				Message: "请求参数错误",
			}
			ctx.Error(e)
			return
		}
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			// 如果是验证错误，则翻译错误信息
			var errorMessages []string
			for _, e := range ve {
				errMsg := e.Translate(translator)
				errorMessages = append(errorMessages, errMsg)
			}
			e := &response.AdminError{
				Code:    http.StatusBadRequest,
				Message: strings.Join(errorMessages, ","),
			}
			ctx.Error(e)
			return
		}
		ctx.Error(err)
		return
	}
	err = t.sysDictService.UpdateSysDictItem(dictItemUpdateReq)
	if err != nil {
		ctx.Error(err)
		return
	}
	return
}

func (t *DictController) DeleteSysDictItemById(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		return
	}
	err = t.sysDictService.DeleteSysDictItemById(id)
	if err != nil {
		ctx.Error(err)
		return
	}
	return
}
