/**
 * @Author: Nan
 * @Date: 2024/7/19 下午5:56
 */

package repository

import (
	"gorm.io/gorm"
	"sweet-cms/model"
)

type SysUserRoleRepository interface {
	BasicRepository
	CreateUserRole(*gorm.DB, model.SysUserRole) error
	DeleteUserRole(*gorm.DB, int) error
	GetUserRoles(userId int) ([]model.SysRole, error)
}
