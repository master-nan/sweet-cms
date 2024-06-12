/**
 * @Author: Nan
 * @Date: 2024/6/12 上午9:46
 */

package cache

import (
	"go.uber.org/zap"
	"sweet-cms/inter"
	"sweet-cms/model"
	"time"
)

type SysTableCache struct {
	cacheInterface inter.CacheInterface
}

const TableCacheKey = "TABLE_CACHE_KEY_"

func NewSysTableCache(cacheInterface inter.CacheInterface) *SysTableCache {
	return &SysTableCache{
		cacheInterface,
	}
}

func (c *SysTableCache) Get(key string) (model.SysTable, error) {
	var data model.SysTable
	err := c.cacheInterface.Get(TableCacheKey+key, &data)
	if err != nil {
		zap.L().Error(TableCacheKey+"Error getting key in cache", zap.String("key", key), zap.Error(err))
		return data, err
	}
	return data, nil
}

func (c *SysTableCache) Set(key string, data model.SysTable) error {
	err := c.cacheInterface.Set(TableCacheKey+key, &data, 7200*time.Second)
	if err != nil {
		zap.L().Error(TableCacheKey+"Error setting key in cache", zap.String("key", key), zap.Error(err))
		return err
	}
	return nil
}

func (c *SysTableCache) Delete(key string) error {
	err := c.cacheInterface.Del(TableCacheKey + key)
	if err != nil {
		zap.L().Error(TableCacheKey+"Error delete key in cache", zap.String("key", key), zap.Error(err))
		return err
	}
	return nil
}
