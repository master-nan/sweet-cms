/**
 * @Author: Nan
 * @Date: 2024/5/17 上午11:12
 */

package controller

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"strconv"
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

func (t *TableController) GetTableByID(ctx *gin.Context) {
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

func (t *TableController) GetTableByCode(ctx *gin.Context) {
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

func (t *TableController) QueryTable(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	var data request.Basic
	translator, _ := t.translators["zh"]
	err := utils.ValidatorQuery[request.Basic](ctx, &data, translator)
	if err != nil {
		_ = ctx.Error(err)
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

func (t *TableController) CreateTable(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	var data request.TableCreateReq
	translator, _ := t.translators["zh"]
	err := utils.ValidatorBody[request.TableCreateReq](ctx, &data, translator)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = t.sysTableService.CreateTable(ctx, data)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	return
}

func (t *TableController) UpdateTable(ctx *gin.Context) {
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

func (t *TableController) DeleteTableById(ctx *gin.Context) {
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

func (t *TableController) GetTableFieldsByTableId(ctx *gin.Context) {
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

func (t *TableController) GetTableFieldById(ctx *gin.Context) {
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

func (t *TableController) CreateTableField(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	var data request.TableFieldCreateReq
	translator, _ := t.translators["zh"]
	err := utils.ValidatorBody[request.TableFieldCreateReq](ctx, &data, translator)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = t.sysTableService.CreateTableField(ctx, data)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	return
}

func (t *TableController) UpdateTableField(ctx *gin.Context) {
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

func (t *TableController) DeleteTableFieldById(ctx *gin.Context) {
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

func (t *TableController) CreateTableRelation(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	var data request.TableRelationCreateReq
	translator, _ := t.translators["zh"]
	err := utils.ValidatorBody[request.TableRelationCreateReq](ctx, &data, translator)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = t.sysTableService.CreateTableRelation(ctx, data)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	return
}

func (t *TableController) UpdateTableRelation(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	var data request.TableRelationUpdateReq
	translator, _ := t.translators["zh"]
	err := utils.ValidatorBody[request.TableRelationUpdateReq](ctx, &data, translator)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = t.sysTableService.UpdateTableRelation(ctx, data)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	return
}

func (t *TableController) DeleteTableRelationById(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = t.sysTableService.DeleteTableRelationById(ctx, id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	return
}

func (t *TableController) GetTableIndexById(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	data, err := t.sysTableService.GetTableIndexById(id)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	resp.SetData(data)
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

func (t *TableController) CreateTableIndex(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	var data request.TableIndexCreateReq
	translator, _ := t.translators["zh"]
	err := utils.ValidatorBody[request.TableIndexCreateReq](ctx, &data, translator)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = t.sysTableService.CreateTableIndex(ctx, data)
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

func (t *TableController) DeleteTableIndexById(ctx *gin.Context) {
	resp := response.NewResponse()
	ctx.Set("response", resp)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = t.sysTableService.DeleteTableIndexById(ctx, id)
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
