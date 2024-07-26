/**
 * @Author: Nan
 * @Date: 2024/7/25 下午10:49
 */

package cache

import (
	"sweet-cms/inter"
	"sweet-cms/model"
)

type SysRoleMenuButtonCache struct {
	*BasicCache[model.SysRoleMenuButton]
}

const RoleMenuButtonCacheKey = "ROLE_MENU_BUTTON_CACHE_KEY_"

func NewSysRoleMenuButtonCache(cacheInterface inter.CacheInterface) *SysRoleMenuButtonCache {
	return &SysRoleMenuButtonCache{
		BasicCache: NewBasicCache[model.SysRoleMenuButton](cacheInterface, RoleMenuButtonCacheKey),
	}
}
