/**
 * @Author: Nan
 * @Date: 2024/6/10 上午12:14
 */

package repository

import (
	"sweet-cms/form/request"
	"sweet-cms/model"
)

type SysTableListResult struct {
	Data  []model.SysTable `json:"data"`
	Total int              `json:"total"`
}

type SysTableRepository interface {
	GetTableById(int) (model.SysTable, error)
	GetTableByTableCode(string) (model.SysTable, error)
	InsertTable(model.SysTable) error
	UpdateTable(request.TableUpdateReq) error
	DeleteTableById(int) error
	GetTableList(request.Basic) (SysTableListResult, error)

	GetTableFieldById(int) (model.SysTableField, error)
	GetTableFieldsByTableId(int) ([]model.SysTableField, error)
	UpdateTableField(request.TableFieldUpdateReq) error
	InsertTableField(model.SysTableField) error
	DeleteTableFieldById(int) error
}
