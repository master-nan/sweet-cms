/**
 * @Author: Nan
 * @Date: 2024/6/28 下午3:00
 */

package cache

import (
	"sweet-cms/inter"
	"sweet-cms/model"
)

type SysUserCache struct {
	*BasicCache[model.SysUser]
}

const UserCacheKey = "USER_CACHE_KEY_"

func NewSysUserCache(cacheInterface inter.CacheInterface) *SysUserCache {
	return &SysUserCache{
		BasicCache: NewBasicCache[model.SysUser](cacheInterface, UserCacheKey),
	}
}
