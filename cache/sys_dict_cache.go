/**
 * @Author: Nan
 * @Date: 2024/6/3 下午5:41
 */

package cache

import (
	"go.uber.org/zap"
	"sweet-cms/inter"
	"sweet-cms/model"
	"time"
)

type SysDictCache struct {
	cacheInterface inter.CacheInterface
}

const DictCacheKey = "DICT_CACHE_KEY_"

func NewSysDictCache(cacheInterface inter.CacheInterface) *SysDictCache {
	return &SysDictCache{cacheInterface}
}

func (c *SysDictCache) Get(key string) (model.SysDict, error) {
	var data model.SysDict
	err := c.cacheInterface.Get(DictCacheKey+key, &data)
	if err != nil {
		zap.L().Error(DictCacheKey+"Error getting key in cache", zap.String("key", key), zap.Error(err))
		return data, err
	}
	return data, nil
}

func (c *SysDictCache) Set(key string, value model.SysDict) error {
	err := c.cacheInterface.Set(DictCacheKey+key, value, 7200*time.Second)
	if err != nil {
		zap.L().Error(DictCacheKey+"Error setting key in cache", zap.String("key", key), zap.Error(err))
		return err
	}
	return nil
}

func (c *SysDictCache) Delete(key string) error {
	err := c.cacheInterface.Del(DictCacheKey + key)
	if err != nil {
		zap.L().Error(DictCacheKey+"Error delete key in cache", zap.String("key", key), zap.Error(err))
		return err
	}
	return nil
}
