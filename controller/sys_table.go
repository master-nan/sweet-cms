/**
 * @Author: Nan
 * @Date: 2024/5/17 上午11:12
 */

package controller

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"sweet-cms/service"
)

type TableController struct {
	sysTableService *service.SysTableService
	translators     map[string]ut.Translator
}

func NewTableController() *TableController {
	return &TableController{}
}

func (t *TableController) GetSysTableByID(c *gin.Context) {

}

func (t *TableController) GetSysTableByCode(c *gin.Context) {

}

func (t *TableController) QuerySysTable(c *gin.Context) {

}

func (t *TableController) InsertSysTable(c *gin.Context) {

}

func (t *TableController) UpdateSysTable(c *gin.Context) {

}
func (t *TableController) DeleteSysTableById(c *gin.Context) {

}

func (t *TableController) GetSysTableFieldById(c *gin.Context)       {}
func (t *TableController) GetSysTableFieldsByTableId(c *gin.Context) {}
func (t *TableController) InsertSysTableField(c *gin.Context)        {}
func (t *TableController) UpdateSysTableField(c *gin.Context)        {}
func (t *TableController) DeleteSysTableFieldById(c *gin.Context)    {}
