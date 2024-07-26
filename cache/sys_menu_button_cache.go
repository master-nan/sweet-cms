/**
 * @Author: Nan
 * @Date: 2024/7/25 下午11:12
 */

package cache

import (
	"sweet-cms/inter"
	"sweet-cms/model"
)

type SysMenuButtonCache struct {
	*BasicCache[model.SysMenuButton]
}

const MenuButtonCacheKey = "MENU_BUTTON_CACHE_KEY_"

func NewSysMenuButtonCache(cacheInterface inter.CacheInterface) *SysMenuButtonCache {
	return &SysMenuButtonCache{
		BasicCache: NewBasicCache[model.SysMenuButton](cacheInterface, MenuButtonCacheKey),
	}
}
