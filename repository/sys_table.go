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
	GetTableList(request.Basic) (response.ListResult[model.SysTable], error)

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
