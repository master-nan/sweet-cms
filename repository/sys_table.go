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
	UpdateTableField(request.TableFieldUpdateReq, model.SysTableField, string) error
	InsertTableField(model.SysTableField, string) error
	DeleteTableField(model.SysTableField, string) error

	GetTableRelationById(int) (model.SysTableRelation, error)
	GetTableRelationByTableId(int) (model.SysTableRelation, error)
	InsertTableRelation(model.SysTableRelation, string) error
	UpdateTableRelation(request.TableRelationUpdateReq, string) error
	DeleteTableRelation(model.SysTableRelation, string) error

	GetTableIndexById(int) (model.SysTableIndex, error)
	GetTableIndexByTableId(int) (model.SysTableIndex, error)
	InsertTableIndex(model.SysTableIndex, string) error
	UpdateTableIndex(request.TableIndexUpdateReq, string) error
	DeleteTableIndex(model.SysTableIndex, string) error
}
