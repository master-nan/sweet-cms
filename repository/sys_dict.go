/**
 * @Author: Nan
 * @Date: 2024/5/25 下午12:01
 */

package repository

import (
	"gorm.io/gorm"
	"sweet-cms/form/request"
	"sweet-cms/form/response"
	"sweet-cms/model"
)

type SysDictRepository interface {
	BasicRepository
	GetSysDictById(int) (model.SysDict, error)
	GetSysDictList(request.Basic) (response.ListResult[model.SysDict], error)
	InsertSysDict(*gorm.DB, model.SysDict) error
	UpdateSysDict(*gorm.DB, request.DictUpdateReq) error
	DeleteSysDictById(*gorm.DB, int) error
	GetSysDictByCode(string) (model.SysDict, error)

	GetSysDictItemById(int) (model.SysDictItem, error)
	GetSysDictItemsByDictId(int) ([]model.SysDictItem, error)
	UpdateSysDictItem(*gorm.DB, request.DictItemUpdateReq) error
	InsertSysDictItem(*gorm.DB, model.SysDictItem) error
	DeleteSysDictItemById(*gorm.DB, int) error
}
