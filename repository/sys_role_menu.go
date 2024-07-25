/**
 * @Author: Nan
 * @Date: 2024/7/19 下午5:56
 */

package repository

import (
	"sweet-cms/model"
)

type SysRoleMenuRepository interface {
	BasicRepository
	GetRoleMenus(int) ([]model.SysMenu, error)
	GetRoleMenusByRoleIds([]int) ([]model.SysMenu, error)
}
