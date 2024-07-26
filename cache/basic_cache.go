/**
 * @Author: Nan
 * @Date: 2024/7/26 下午3:14
 */

package cache

import (
	"go.uber.org/zap"
	"sweet-cms/inter"
	"time"
)

type BasicCache[T any] struct {
	cacheInterface inter.CacheInterface
	cacheKeyPrefix string
}

func NewBasicCache[T any](cacheInterface inter.CacheInterface, cacheKeyPrefix string) *BasicCache[T] {
	return &BasicCache[T]{
		cacheInterface,
		cacheKeyPrefix,
	}
}
func (c *BasicCache[T]) Get(key string) (T, error) {
	var data T
	err := c.cacheInterface.Get(c.cacheKeyPrefix+key, &data)
	if err != nil {
		zap.L().Error(c.cacheKeyPrefix+"Error getting key in cache", zap.String("key", key), zap.Error(err))
		return data, err
	}
	return data, nil
}

func (c *BasicCache[T]) Set(key string, data T) error {
	err := c.cacheInterface.Set(c.cacheKeyPrefix+key, &data, 7200*time.Second)
	if err != nil {
		zap.L().Error(c.cacheKeyPrefix+"Error setting key in cache", zap.String("key", key), zap.Error(err))
		return err
	}
	return nil
}

func (c *BasicCache[T]) Delete(key string) error {
	err := c.cacheInterface.Del(c.cacheKeyPrefix + key)
	if err != nil {
		zap.L().Error(c.cacheKeyPrefix+"Error delete key in cache", zap.String("key", key), zap.Error(err))
		return err
	}
	return nil
}

func (c *BasicCache[T]) Exists(key string) bool {
	exists, _ := c.cacheInterface.Exists(c.cacheKeyPrefix + key)
	return exists > 0
}
