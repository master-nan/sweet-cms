/**
 * @Author: Nan
 * @Date: 2024/6/3 下午6:07
 */

package repository

import (
	"sweet-cms/form/request"
	"sweet-cms/form/response"
	"sweet-cms/model"
)

type SysUserRepository interface {
	BasicRepository
	GetByUserName(string) (model.SysUser, error)
	GetById(int) (model.SysUser, error)
	Update(request.UserUpdateReq) error
	DeleteById(int) error
	GetList(request.Basic) (response.ListResult[model.SysUser], error)
	Insert(model.SysUser) error
	GetByEmployeeID(int) (model.SysUser, error)
}
