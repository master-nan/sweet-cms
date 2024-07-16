/**
 * @Author: Nan
 * @Date: 2024/6/10 上午12:14
 */

package repository

import (
	"gorm.io/gorm"
	"sweet-cms/form/request"
	"sweet-cms/form/response"
	"sweet-cms/model"
)

type SysTableRepository interface {
	BasicRepository
	GetTableById(int) (model.SysTable, error)
	GetTableByTableCode(string) (model.SysTable, error)
	InsertTable(*gorm.DB, model.SysTable) error

	UpdateTable(request.TableUpdateReq) error
	DeleteTableById(*gorm.DB, int) error
	GetTableList(request.Basic) (response.ListResult[model.SysTable], error)

	GetTableFieldById(int) (model.SysTableField, error)
	GetTableFieldsByTableId(int) ([]model.SysTableField, error)
	UpdateTableField(*gorm.DB, request.TableFieldUpdateReq) error
	InsertTableField(*gorm.DB, model.SysTableField) error
	DeleteTableField(*gorm.DB, int) error

	DeleteTableFieldByTableId(*gorm.DB, int) error

	GetTableRelationById(int) (model.SysTableRelation, error)
	GetTableRelationsByTableId(int) ([]model.SysTableRelation, error)
	InsertTableRelation(*gorm.DB, model.SysTableRelation) error
	DeleteTableRelation(*gorm.DB, int) error

	GetTableIndexesByTableId(int) ([]model.SysTableIndex, error)
	GetTableIndexById(int) (model.SysTableIndex, error)
	InsertTableIndex(*gorm.DB, model.SysTableIndex) error
	UpdateTableIndex(*gorm.DB, request.TableIndexUpdateReq) error
	DeleteTableIndex(*gorm.DB, int) error
	DeleteTableIndexByTableId(*gorm.DB, int) error

	InsertTableIndexFields(*gorm.DB, []model.SysTableIndexField) error
	DeleteTableIndexFieldByIndexId(*gorm.DB, int) error
	DeleteTableIndexFieldByIndexIds(*gorm.DB, []int) error

	FetchTableMetadata(string, string) ([]model.TableColumnMate, error)
	FetchTableIndexMetadata(string, string) ([]model.TableIndexMate, error)
	DropTableIndex(*gorm.DB, string, string) error
	DropTable(*gorm.DB, string) error
	DropTableColumn(*gorm.DB, string, string) error
	ModifyTableColumn(*gorm.DB, string, string, string) error
	// ChangeTableColumn 修改字段
	ChangeTableColumn(*gorm.DB, string, string, string, string) error
	CreateTableColumn(*gorm.DB, string, string, string) error
	CreateTable(*gorm.DB, string, any) error
	CreateTableIndex(*gorm.DB, bool, string, string, string) error

	Model([]model.SysTableField) interface{}
}
