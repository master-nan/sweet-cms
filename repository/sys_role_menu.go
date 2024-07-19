/**
 * @Author: Nan
 * @Date: 2024/7/19 下午5:56
 */

package repository

import (
	"gorm.io/gorm"
	"sweet-cms/model"
)

type SysRoleMenuRepository interface {
	BasicRepository
	CreateRoleMenu(*gorm.DB, model.SysRoleMenu) error
	DeleteRoleMenu(*gorm.DB, int) error
	GetRoleMenus(roleId int) ([]model.SysMenu, error)
}
