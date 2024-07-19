/**
 * @Author: Nan
 * @Date: 2024/7/19 下午5:57
 */

package repository

import (
	"gorm.io/gorm"
	"sweet-cms/model"
)

type SysRoleMenuButtonRepository interface {
	BasicRepository
	CreateRoleMenuButton(*gorm.DB, model.SysRoleMenuButton) error
	DeleteRoleMenuButton(*gorm.DB, int) error
	GetRoleMenuButtons(roleId, menuId int) ([]model.SysMenuButton, error)
}
