/**
 * @Author: Nan
 * @Date: 2024/7/19 下午5:06
 */

package repository

import (
	"sweet-cms/form/request"
	"sweet-cms/form/response"
	"sweet-cms/model"
)

type SysRoleRepository interface {
	BasicRepository
	GetRoles() ([]model.SysRole, error)
	GetRoleMenus(roleId int) ([]model.SysMenu, error)
	GetRoleButtons(roleId int) ([]model.SysMenuButton, error)
	GetRoleList(basic request.Basic) (response.ListResult[model.SysRole], error)
}
