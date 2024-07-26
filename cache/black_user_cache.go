/**
 * @Author: Nan
 * @Date: 2024/7/4 下午5:21
 */

package cache

import (
	"sweet-cms/inter"
	"sweet-cms/model"
)

type BlackUserCache struct {
	*BasicCache[model.SysUser]
}

const BlackUserCacheKey = "BLACK_USER_CACHE_KEY_"

func NewBlackCache(cacheInterface inter.CacheInterface) *BlackUserCache {
	return &BlackUserCache{
		BasicCache: NewBasicCache[model.SysUser](cacheInterface, BlackUserCacheKey),
	}
}
