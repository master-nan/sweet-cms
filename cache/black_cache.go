/**
 * @Author: Nan
 * @Date: 2024/7/4 下午5:21
 */

package cache

import (
	"go.uber.org/zap"
	"sweet-cms/inter"
	"sweet-cms/model"
	"time"
)

type BlackCache struct {
	cacheInterface inter.CacheInterface
}

const BlackCacheKey = "BLACK_CACHE_KEY_"

func NewBlackCache(cacheInterface inter.CacheInterface) *BlackCache {
	return &BlackCache{cacheInterface: cacheInterface}
}

func (c *BlackCache) Get(key string) (model.SysUser, error) {
	var data model.SysUser
	err := c.cacheInterface.Get(BlackCacheKey+key, &data)
	if err != nil {
		zap.L().Error(BlackCacheKey+"Error getting key in cache", zap.String("key", key), zap.Error(err))
		return data, err
	}
	return data, nil
}

func (c *BlackCache) Set(key string, data model.SysUser) error {
	err := c.cacheInterface.Set(BlackCacheKey+key, &data, 7200*time.Second)
	if err != nil {
		zap.L().Error(BlackCacheKey+"Error setting key in cache", zap.String("key", key), zap.Error(err))
		return err
	}
	return nil
}

func (c *BlackCache) Delete(key string) error {
	err := c.cacheInterface.Del(BlackCacheKey + key)
	if err != nil {
		zap.L().Error(BlackCacheKey+"Error delete key in cache", zap.String("key", key), zap.Error(err))
		return err
	}
	return nil
}

func (c *BlackCache) Exists(key string) bool {
	exists, _ := c.cacheInterface.Exists(BlackCacheKey + key)
	return exists > 0
}
