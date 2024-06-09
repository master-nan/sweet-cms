/**
 * @Author: Nan
 * @Date: 2024/5/17 上午11:12
 */

package controller

import (
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

func (t *TableController) Get() {

}

func (t *TableController) Query() {

}

func (t *TableController) Insert() {

}

func (t *TableController) Update() {

}
func (t *TableController) Delete() {

}

type TableFieldController struct {
}

func NewTableFieldController() *TableFieldController {
	return &TableFieldController{}
}

func (t *TableFieldController) Get()    {}
func (t *TableFieldController) Query()  {}
func (t *TableFieldController) Insert() {}
func (t *TableFieldController) Update() {}
func (t *TableFieldController) Delete() {}
