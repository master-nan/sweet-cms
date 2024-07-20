/**
 * @Author: Nan
 * @Date: 2024/7/19 上午11:24
 */

package repository

import (
	"sweet-cms/model"
)

type SysMenuRepository interface {
	BasicRepository
	GetMenuById(int) (model.SysMenu, error)
	GetMenus() ([]model.SysMenu, error)
}
