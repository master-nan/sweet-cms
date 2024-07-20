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
	GetUserMenuPermissionsByUserId(userId int) ([]model.SysUserMenuDataPermission, error)
}
