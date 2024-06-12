/**
 * @Author: Nan
 * @Date: 2024/5/17 上午11:12
 */

package controller

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"strconv"
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
	id, err := strconv.Atoi(ctx.Param("id"))
	resp := response.NewResponse()
	ctx.Set("response", resp)
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

}

func (t *TableController) QuerySysTable(ctx *gin.Context) {

}

func (t *TableController) InsertSysTable(ctx *gin.Context) {

}

func (t *TableController) UpdateSysTable(ctx *gin.Context) {

}
func (t *TableController) DeleteSysTableById(ctx *gin.Context) {

}

func (t *TableController) GetSysTableFieldById(ctx *gin.Context)       {}
func (t *TableController) GetSysTableFieldsByTableId(ctx *gin.Context) {}
func (t *TableController) InsertSysTableField(ctx *gin.Context)        {}
func (t *TableController) UpdateSysTableField(ctx *gin.Context)        {}
func (t *TableController) DeleteSysTableFieldById(ctx *gin.Context)    {}
