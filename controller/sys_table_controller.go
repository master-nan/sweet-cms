/**
 * @Author: Nan
 * @Date: 2024/5/17 上午11:12
 */

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sweet-cms/form/request"
	"sweet-cms/form/response"
	"sweet-cms/service"
)

type TableController struct {
	sysTableService *service.SysTableService
	translators     map[string]ut.Translator
}

func NewTableController(sysTableService *service.SysTableService, translators map[string]ut.Translator) *TableController {
	return &TableController{
		sysTableService,
		translators,
	}
}

func (t *TableController) GetSysTableByID(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		return
	}
	data, err := t.sysTableService.GetTableById(id)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp.SetData(data)
	return
}

func (t *TableController) GetSysTableByCode(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	code := ctx.Param("code")
	data, err := t.sysTableService.GetTableByTableCode(code)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp.SetData(data)
	return
}

func (t *TableController) QuerySysTable(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	var data request.Basic
	translator, _ := t.translators["zh"]
	if err := ctx.ShouldBindQuery(&data); err != nil {
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
		e := &response.AdminError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		ctx.Error(e)
		return
	}
	result, err := t.sysTableService.GetTableList(data)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp.SetData(result.Data).SetTotal(result.Total)
	return
}

func (t *TableController) InsertSysTable(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	var data request.TableCreateReq
	translator, _ := t.translators["zh"]
	if err := ctx.ShouldBindBodyWith(&data, binding.JSON); err != nil {
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
	err := t.sysTableService.InsertTable(data)
	if err != nil {
		ctx.Error(err)
		return
	}
	return
}

func (t *TableController) UpdateSysTable(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	var data request.TableUpdateReq
	translator, _ := t.translators["zh"]
	if err := ctx.ShouldBindBodyWith(&data, binding.JSON); err != nil {
		if err == io.EOF {
			e := &response.AdminError{
				Code:    http.StatusBadRequest,
				Message: "请求参数错误",
			}
			ctx.Error(e)
			return
		}
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
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
	err := t.sysTableService.UpdateTable(data)
	if err != nil {
		ctx.Error(err)
		return
	}
	return
}

func (t *TableController) DeleteSysTableById(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		return
	}
	err = t.sysTableService.DeleteTableById(id)
	if err != nil {
		ctx.Error(err)
		return
	}
	return
}

func (t *TableController) GetSysTableFieldsByTableId(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		return
	}
	data, err := t.sysTableService.GetTableFieldsByTableId(id)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp.SetData(data)
	resp.SetData(len(data))
	return
}

func (t *TableController) GetSysTableFieldById(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		return
	}
	data, err := t.sysTableService.GetTableFieldById(id)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp.SetData(data)
	return
}

func (t *TableController) InsertSysTableField(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	var data request.TableFieldCreateReq
	translator, _ := t.translators["zh"]
	if err := ctx.ShouldBindBodyWith(&data, binding.JSON); err != nil {
		if err == io.EOF {
			e := &response.AdminError{
				Code:    http.StatusBadRequest,
				Message: "请求参数错误",
			}
			ctx.Error(e)
			return
		}
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
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
	err := t.sysTableService.InsertTableField(data)
	if err != nil {
		ctx.Error(err)
		return
	}
	return
}

func (t *TableController) UpdateSysTableField(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	var data request.TableFieldUpdateReq
	translator, _ := t.translators["zh"]
	if err := ctx.ShouldBindBodyWith(&data, binding.JSON); err != nil {
		if err == io.EOF {
			e := &response.AdminError{
				Code:    http.StatusBadRequest,
				Message: "请求参数错误",
			}
			ctx.Error(e)
			return
		}
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
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
	err := t.sysTableService.UpdateTableField(data)
	if err != nil {
		ctx.Error(err)
		return
	}
	return
}

func (t *TableController) DeleteSysTableFieldById(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Error(err)
		return
	}
	err = t.sysTableService.DeleteTableFieldById(id)
	if err != nil {
		ctx.Error(err)
		return
	}
	return
}
