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

type SysTableFieldCache struct {
	cacheInterface inter.CacheInterface
}

const TableFieldCacheKey = "TABLE_FIELD_CACHE_KEY_"

func NewSysTableFieldCache(cacheInterface inter.CacheInterface) *SysTableFieldCache {
	return &SysTableFieldCache{
		cacheInterface,
	}
}

func (c *SysTableFieldCache) Get(key string) (model.SysTableField, error) {
	var data model.SysTableField
	err := c.cacheInterface.Get(TableFieldCacheKey+key, &data)
	if err != nil {
		zap.L().Error(TableFieldCacheKey+"Error getting key in cache", zap.String("key", key), zap.Error(err))
		return data, err
	}
	return data, nil
}

func (c *SysTableFieldCache) Set(key string, data model.SysTableField) error {
	err := c.cacheInterface.Set(TableFieldCacheKey+key, &data, 7200*time.Second)
	if err != nil {
		zap.L().Error(TableFieldCacheKey+"Error setting key in cache", zap.String("key", key), zap.Error(err))
		return err
	}
	return nil
}

func (c *SysTableFieldCache) Delete(key string) error {
	err := c.cacheInterface.Del(TableFieldCacheKey + key)
	if err != nil {
		zap.L().Error(TableFieldCacheKey+"Error delete key in cache", zap.String("key", key), zap.Error(err))
		return err
	}
	return nil
}
