/**
 * @Author: Nan
 * @Date: 2024/7/25 下午11:13
 */

package cache

import (
	"sweet-cms/inter"
	"sweet-cms/model"
)

type SysUserMenuDataPermissionCache struct {
	*BasicCache[model.SysUserMenuDataPermission]
}

const UserMenuDataPermissionCacheKey = "USER_MENU_DATA_PERMISSION_CACHE_KEY_"

func NewSysUserMenuDataPermissionCache(cacheInterface inter.CacheInterface) *SysUserMenuDataPermissionCache {
	return &SysUserMenuDataPermissionCache{
		BasicCache: NewBasicCache[model.SysUserMenuDataPermission](cacheInterface, UserMenuDataPermissionCacheKey),
	}
}
