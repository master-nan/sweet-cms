/**
 * @Author: Nan
 * @Date: 2024/6/12 上午9:46
 */

package cache

import (
	"sweet-cms/inter"
)

type SysTableCache struct {
	cacheInterface inter.CacheInterface
}

func NewSysTableCache(cacheInterface inter.CacheInterface) *SysTableCache {
	return &SysTableCache{
		cacheInterface,
	}
}
