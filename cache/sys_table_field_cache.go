/**
 * @Author: Nan
 * @Date: 2024/6/12 上午9:46
 */

package cache

import (
	"sweet-cms/inter"
	"sweet-cms/model"
)

type SysTableFieldCache struct {
	*BasicCache[model.SysTableField]
}

const TableFieldCacheKey = "TABLE_FIELD_CACHE_KEY_"

func NewSysTableFieldCache(cacheInterface inter.CacheInterface) *SysTableFieldCache {
	return &SysTableFieldCache{
		BasicCache: NewBasicCache[model.SysTableField](cacheInterface, TableFieldCacheKey),
	}
}
