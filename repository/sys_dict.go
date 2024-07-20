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
	GetSysDictByCode(string) (model.SysDict, error)
}
