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
	GetList(request.Basic) (response.ListResult[model.SysUser], error)
	GetByEmployeeID(int) (model.SysUser, error)
	GetUserMenuPermissions(userId int) ([]model.SysUserMenuDataPermission, error)
	GetUserMenus(userId int) ([]model.SysMenu, error)
}
