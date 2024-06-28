/**
 * @Author: Nan
 * @Date: 2024/6/28 下午3:00
 */

package cache

import (
	"go.uber.org/zap"
	"sweet-cms/inter"
	"sweet-cms/model"
	"time"
)

type SysUserCache struct {
	cacheInterface inter.CacheInterface
}

const UserCacheKey = "USER_CACHE_KEY_"

func NewSysUserCache(cacheInterface inter.CacheInterface) *SysUserCache {
	return &SysUserCache{cacheInterface: cacheInterface}
}

func (c *SysUserCache) Get(key string) (model.SysTable, error) {
	var data model.SysTable
	err := c.cacheInterface.Get(UserCacheKey+key, &data)
	if err != nil {
		zap.L().Error(UserCacheKey+"Error getting key in cache", zap.String("key", key), zap.Error(err))
		return data, err
	}
	return data, nil
}

func (c *SysUserCache) Set(key string, data model.SysTable) error {
	err := c.cacheInterface.Set(UserCacheKey+key, &data, 7200*time.Second)
	if err != nil {
		zap.L().Error(UserCacheKey+"Error setting key in cache", zap.String("key", key), zap.Error(err))
		return err
	}
	return nil
}

func (c *SysUserCache) Delete(key string) error {
	err := c.cacheInterface.Del(UserCacheKey + key)
	if err != nil {
		zap.L().Error(UserCacheKey+"Error delete key in cache", zap.String("key", key), zap.Error(err))
		return err
	}
	return nil
}
