/**
 * @Author: Nan
 * @Date: 2024/6/3 下午6:07
 */

package repository

import "sweet-cms/model"

type SysUserRepository interface {
	GetByUserName(string) (model.SysUser, error)
	GetByUserId(id int) (model.SysUser, error)
}
