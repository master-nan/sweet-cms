/**
 * @Author: Nan
 * @Date: 2024/7/19 下午5:57
 */

package repository

import (
	"sweet-cms/model"
)

type SysRoleMenuButtonRepository interface {
	BasicRepository
	GetRoleMenuButtons(roleId, menuId int) ([]model.SysMenuButton, error)
}
