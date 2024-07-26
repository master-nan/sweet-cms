/**
 * @Author: Nan
 * @Date: 2024/7/25 下午10:49
 */

package cache

import (
	"sweet-cms/inter"
	"sweet-cms/model"
)

type SysRoleCache struct {
	*BasicCache[model.SysRole]
}

const RoleCacheKey = "ROLE_CACHE_KEY_"

func NewSysRoleCache(cacheInterface inter.CacheInterface) *SysRoleCache {
	return &SysRoleCache{
		BasicCache: NewBasicCache[model.SysRole](cacheInterface, RoleCacheKey),
	}
}
