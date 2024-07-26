/**
 * @Author: Nan
 * @Date: 2024/7/25 下午10:48
 */

package cache

import (
	"sweet-cms/inter"
	"sweet-cms/model"
)

type SysRoleMenuCache struct {
	*BasicCache[model.SysRoleMenu]
}

const RoleMenuCacheKey = "ROLE_MENU_CACHE_KEY_"

func NewSysRoleMenuCache(cacheInterface inter.CacheInterface) *SysRoleMenuCache {
	return &SysRoleMenuCache{
		BasicCache: NewBasicCache[model.SysRoleMenu](cacheInterface, RoleMenuCacheKey),
	}
}
