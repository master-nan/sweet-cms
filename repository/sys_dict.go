/**
 * @Author: Nan
 * @Date: 2024/5/25 下午12:01
 */

package repository

import "sweet-cms/model"

type SysDictRepository interface {
	GetSysDictById(id int) (model.SysDict, error)
	GetSysDictList() ([]model.SysDict, int, error)
	UpdateSysDict(*model.SysDict) error
	InsertSysDict(*model.SysDict) error
	DeleteSysDictById(id int) error
	GetSysDictByCode(code int) (model.SysDict, error)

	GetSysDictItemById(id int) (model.SysDictItem, error)
	GetSysDictItemsByDictId(id int) ([]model.SysDictItem, int, error)
	UpdateSysDictItem(*model.SysDictItem) error
	InsertSysDictItem(*model.SysDictItem) error
	DeleteSysDictItemById(id int) error
}
