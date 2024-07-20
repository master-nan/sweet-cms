/**
 * @Author: Nan
 * @Date: 2024/7/19 下午5:56
 */

package repository

import (
	"sweet-cms/model"
)

type SysUserRoleRepository interface {
	BasicRepository
	GetUserRoles(userId int) ([]model.SysRole, error)
}
