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
	GetRoleMenus(int) ([]model.SysMenu, error)
	GetRoleMenusByRoleIds([]int) ([]model.SysMenu, error)
	DeleteRoleMenuByRoleIdAndMenuId(*gorm.DB, int, int) error
}
