/**
 * @Author: Nan
 * @Date: 2024/6/3 下午5:41
 */

package cache

import (
	"sweet-cms/inter"
	"sweet-cms/repository"
)

type SysDictCache struct {
	sysDictRepo    *repository.SysDictRepository
	cacheInterface inter.CacheInterface
}

func NewSysDictCache(sysDictRepo *repository.SysDictRepository, cacheInterface inter.CacheInterface) *SysDictCache {
	return &SysDictCache{sysDictRepo, cacheInterface}
}
