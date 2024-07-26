/**
 * @Author: Nan
 * @Date: 2024/5/21 下午2:22
 */

package cache

import (
	"sweet-cms/inter"
	"sweet-cms/model"
)

type SysConfigureCache struct {
	*BasicCache[model.SysConfigure]
}

const ConfigureCacheKey = "CONFIGURE_CACHE_KEY_"

func NewSysConfigureCache(cacheInterface inter.CacheInterface) *SysConfigureCache {
	return &SysConfigureCache{
		BasicCache: NewBasicCache[model.SysConfigure](cacheInterface, ConfigureCacheKey),
	}
}
