/**
 * @Author: Nan
 * @Date: 2024/7/19 下午5:06
 */

package repository

import (
	"gorm.io/gorm"
	"sweet-cms/model"
)

type SysRoleRepository interface {
	BasicRepository
	CreateRole(*gorm.DB, model.SysRole) error
	UpdateRole(*gorm.DB, model.SysRole) error
	DeleteRole(*gorm.DB, int) error
	GetRoleById(roleId int) (model.SysRole, error)
	GetRoles() ([]model.SysRole, error)
	GetRoleMenus(roleId int) ([]model.SysMenu, error)
	GetRoleButtons(roleId int) ([]model.SysMenuButton, error)
}
