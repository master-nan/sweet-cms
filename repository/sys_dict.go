/**
 * @Author: Nan
 * @Date: 2024/5/25 下午12:01
 */

package repository

import (
	"sweet-cms/form/request"
	"sweet-cms/form/response"
	"sweet-cms/model"
)

type SysDictRepository interface {
	BasicRepository
	GetSysDictById(int) (model.SysDict, error)
	GetSysDictList(request.Basic) (response.ListResult[model.SysDict], error)
	InsertSysDict(model.SysDict) error
	UpdateSysDict(request.DictUpdateReq) error
	DeleteSysDictById(int) error
	GetSysDictByCode(string) (model.SysDict, error)

	GetSysDictItemById(int) (model.SysDictItem, error)
	GetSysDictItemsByDictId(int) ([]model.SysDictItem, error)
	UpdateSysDictItem(request.DictItemUpdateReq) error
	InsertSysDictItem(model.SysDictItem) error
	DeleteSysDictItemById(int) error
}
