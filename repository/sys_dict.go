/**
 * @Author: Nan
 * @Date: 2024/5/25 下午12:01
 */

package repository

import (
	"sweet-cms/form/request"
	"sweet-cms/model"
)

type SysDictListResult struct {
	Data  []model.SysDict `json:"data"`
	Total int             `json:"total"`
}

type SysDictRepository interface {
	GetSysDictById(int) (model.SysDict, error)
	GetSysDictList(request.Basic) (SysDictListResult, error)
	InsertSysDict(model.SysDict) error
	UpdateSysDict(request.DictUpdateReq) error
	DeleteSysDictById(int) error
	GetSysDictByCode(string) (model.SysDict, error)

	GetSysDictItemById(int) (model.SysDictItem, error)
	GetSysDictItemsByDictId(int) ([]model.SysDictItem, error)
	UpdateSysDictItem(*model.SysDictItem) error
	InsertSysDictItem(*model.SysDictItem) error
	DeleteSysDictItemById(int) error
}
