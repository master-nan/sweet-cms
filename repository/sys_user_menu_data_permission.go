/**
 * @Author: Nan
 * @Date: 2024/7/19 下午4:46
 */

package repository

import (
	"gorm.io/gorm"
	"sweet-cms/model"
)

type SysUserMenuDataPermissionRepository interface {
	BasicRepository
	GetUserMenuPermissionsByUserId(userId int) ([]model.SysUserMenuDataPermission, error)
	CreateUserMenuPermission(tx *gorm.DB, permission model.SysUserMenuDataPermission) error
	UpdateUserMenuPermission(tx *gorm.DB, permission model.SysUserMenuDataPermission) error
	DeleteUserMenuPermission(tx *gorm.DB, id int) error
}
