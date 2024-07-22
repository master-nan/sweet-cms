/**
 * @Author: Nan
 * @Date: 2024/5/17 上午11:12
 */

package controller

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	"strings"
	"sweet-cms/form/request"
	"sweet-cms/form/response"
	"sweet-cms/service"
	"sweet-cms/utils"
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
		_ = ctx.Error(err)
		return
	}
	data, err := t.sysTableService.GetTableById(id)
	if err != nil {
		_ = ctx.Error(err)
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
		_ = ctx.Error(err)
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
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: strings.Join(errorMessages, ","),
			}
			_ = ctx.Error(e)
			return
		}
		e := &response.AdminError{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: err.Error(),
		}
		_ = ctx.Error(e)
		return
	}
	result, err := t.sysTableService.GetTableList(data)
	if err != nil {
		_ = ctx.Error(err)
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
	err := utils.ValidatorBody[request.TableCreateReq](ctx, &data, translator)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = t.sysTableService.InsertTable(ctx, data)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	return
}

func (t *TableController) UpdateSysTable(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	var data request.TableUpdateReq
	translator, _ := t.translators["zh"]
	err := utils.ValidatorBody[request.TableUpdateReq](ctx, &data, translator)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = t.sysTableService.UpdateTable(ctx, data)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	return
}

func (t *TableController) DeleteSysTableById(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = t.sysTableService.DeleteTableById(ctx, id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	return
}

func (t *TableController) GetSysTableFieldsByTableId(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	data, err := t.sysTableService.GetTableFieldsByTableId(id)
	if err != nil {
		_ = ctx.Error(err)
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
		_ = ctx.Error(err)
		return
	}
	data, err := t.sysTableService.GetTableFieldById(id)
	if err != nil {
		_ = ctx.Error(err)
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
	err := utils.ValidatorBody[request.TableFieldCreateReq](ctx, &data, translator)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = t.sysTableService.InsertTableField(ctx, data)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	return
}

func (t *TableController) UpdateSysTableField(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	var data request.TableFieldUpdateReq
	translator, _ := t.translators["zh"]
	err := utils.ValidatorBody[request.TableFieldUpdateReq](ctx, &data, translator)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = t.sysTableService.UpdateTableField(ctx, data)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	return
}

func (t *TableController) DeleteSysTableFieldById(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = t.sysTableService.DeleteTableFieldById(ctx, id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	return
}

func (t *TableController) GetTableRelationsByTableId(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	data, err := t.sysTableService.GetTableRelationsByTableId(id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	resp.SetTotal(len(data))
	resp.SetData(data)
	return
}

func (t *TableController) GetTableRelationById(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	data, err := t.sysTableService.GetTableRelationById(id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	resp.SetData(data)
	return
}

func (t *TableController) InsertTableRelation(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	var data request.TableRelationCreateReq
	translator, _ := t.translators["zh"]
	err := utils.ValidatorBody[request.TableRelationCreateReq](ctx, &data, translator)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = t.sysTableService.InsertTableRelation(ctx, data)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	return
}

func (t *TableController) DeleteTableRelation(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = t.sysTableService.DeleteTableRelation(ctx, id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	return
}

func (t *TableController) GetTableIndexesByTableId(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	data, err := t.sysTableService.GetTableIndexesByTableId(id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	resp.SetData(data)
	return
}

func (t *TableController) InsertTableIndex(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	var data request.TableIndexCreateReq
	translator, _ := t.translators["zh"]
	err := utils.ValidatorBody[request.TableIndexCreateReq](ctx, &data, translator)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = t.sysTableService.InsertTableIndex(ctx, data)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	return
}

func (t *TableController) UpdateTableIndex(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	var data request.TableIndexUpdateReq
	translator, _ := t.translators["zh"]
	err := utils.ValidatorBody[request.TableIndexUpdateReq](ctx, &data, translator)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = t.sysTableService.UpdateTableIndex(ctx, data)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	return
}

func (t *TableController) DeleteTableIndex(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = t.sysTableService.DeleteTableIndex(ctx, id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	return
}

func (t *TableController) DeleteTableIndexByTableId(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = t.sysTableService.DeleteTableIndexByTableId(ctx, id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	return
}

func (t *TableController) InitTable(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	code := ctx.Param("code")
	err := t.sysTableService.InitTable(ctx, code)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	return
}
