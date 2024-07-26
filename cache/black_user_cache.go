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

type BlackUserCache struct {
	cacheInterface inter.CacheInterface
}

const BlackUserCacheKey = "BLACK_USER_CACHE_KEY_"

func NewBlackCache(cacheInterface inter.CacheInterface) *BlackUserCache {
	return &BlackUserCache{cacheInterface: cacheInterface}
}

func (c *BlackUserCache) Get(key string) (model.SysUser, error) {
	var data model.SysUser
	err := c.cacheInterface.Get(BlackUserCacheKey+key, &data)
	if err != nil {
		zap.L().Error(BlackUserCacheKey+"Error getting key in cache", zap.String("key", key), zap.Error(err))
		return data, err
	}
	return data, nil
}

func (c *BlackUserCache) Set(key string, data model.SysUser) error {
	err := c.cacheInterface.Set(BlackUserCacheKey+key, &data, 7200*time.Second)
	if err != nil {
		zap.L().Error(BlackUserCacheKey+"Error setting key in cache", zap.String("key", key), zap.Error(err))
		return err
	}
	return nil
}

func (c *BlackUserCache) Delete(key string) error {
	err := c.cacheInterface.Del(BlackUserCacheKey + key)
	if err != nil {
		zap.L().Error(BlackUserCacheKey+"Error delete key in cache", zap.String("key", key), zap.Error(err))
		return err
	}
	return nil
}

func (c *BlackUserCache) Exists(key string) bool {
	exists, _ := c.cacheInterface.Exists(BlackUserCacheKey + key)
	return exists > 0
}
