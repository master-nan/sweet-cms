/**
 * @Author: Nan
 * @Date: 2024/7/25 下午10:48
 */

package cache

import (
	"sweet-cms/inter"
	"sweet-cms/model"
)

type SysMenuCache struct {
	*BasicCache[model.SysMenu]
}

const MenuCacheKey = "MENU_CACHE_KEY_"

func NewSysMenuCache(cacheInterface inter.CacheInterface) *SysMenuCache {
	return &SysMenuCache{
		BasicCache: NewBasicCache[model.SysMenu](cacheInterface, MenuCacheKey),
	}
}
