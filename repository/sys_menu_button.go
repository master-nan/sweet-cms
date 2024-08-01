/**
 * @Author: Nan
 * @Date: 2024/8/1 下午10:36
 */

package repository

import "sweet-cms/model"

type SysMenuButtonRepository interface {
	GetMenuButtonsByMenuId(int) ([]model.SysMenuButton, error)
}
