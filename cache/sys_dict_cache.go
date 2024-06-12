/**
 * @Author: Nan
 * @Date: 2024/6/3 下午5:41
 */

package cache

import (
	"sweet-cms/inter"
)

type SysDictCache struct {
	cacheInterface inter.CacheInterface
}

func NewSysDictCache(cacheInterface inter.CacheInterface) *SysDictCache {
	return &SysDictCache{cacheInterface}
}
