/**
 * @Author: Nan
 * @Date: 2024/7/19 上午11:24
 */

package repository

import (
	"gorm.io/gorm"
	"sweet-cms/model"
)

type SysMenuRepository interface {
	BasicRepository
	GetMenuById(int) (model.SysMenu, error)
	CreateMenu(*gorm.DB, model.SysMenu) error
	UpdateMenu(*gorm.DB, model.SysMenu) error
	DeleteMenu(*gorm.DB, int) error
	GetMenus() ([]model.SysMenu, error)
	GetMenuUserPermissionsByMenuId(int) ([]model.SysUserMenuDataPermission, error)
	GetMenuUsersByMenuId(int) ([]model.SysUser, error)
}
