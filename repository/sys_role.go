/**
 * @Author: Nan
 * @Date: 2024/7/19 下午5:06
 */

package repository

import (
	"sweet-cms/model"
)

type SysRoleRepository interface {
	BasicRepository
	GetRoleById(roleId int) (model.SysRole, error)
	GetRoles() ([]model.SysRole, error)
	GetRoleMenus(roleId int) ([]model.SysMenu, error)
	GetRoleButtons(roleId int) ([]model.SysMenuButton, error)
}