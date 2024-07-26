/**
 * @Author: Nan
 * @Date: 2024/6/3 下午5:41
 */

package cache

import (
	"sweet-cms/inter"
	"sweet-cms/model"
)

type SysDictCache struct {
	*BasicCache[model.SysDict]
}

const DictCacheKey = "DICT_CACHE_KEY_"

func NewSysDictCache(cacheInterface inter.CacheInterface) *SysDictCache {
	return &SysDictCache{BasicCache: NewBasicCache[model.SysDict](cacheInterface, DictCacheKey)}
}
