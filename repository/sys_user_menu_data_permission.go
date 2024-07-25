/**
 * @Author: Nan
 * @Date: 2024/7/19 下午4:46
 */

package repository

import (
	"sweet-cms/model"
)

type SysUserMenuDataPermissionRepository interface {
	BasicRepository
	GetUserMenuPermissionsByUserId(int) ([]model.SysUserMenuDataPermission, error)
	GetUserMenuPermissions(int) ([]model.SysUserMenuDataPermission, error)
}
