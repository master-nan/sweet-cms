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
	GetByUserName(string) (model.SysUser, error)
	GetByUserId(int) (model.SysUser, error)
	UpdateUser(request.UserUpdateReq) error
	DeleteUserById(int) error
	GetList(request.Basic) (response.ListResult[model.SysUser], error)
	Insert(model.SysUser) error
	GetByEmployeeID(int) (model.SysUser, error)
}
