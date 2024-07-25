/**
 * @Author: Nan
 * @Date: 2024/7/25 下午11:19
 */

package cache

import (
	"go.uber.org/zap"
	"sweet-cms/inter"
	"sweet-cms/model"
	"time"
)

type SysUserRoleCache struct {
	cacheInterface inter.CacheInterface
}

const UserRoleCacheKey = "USER_ROLE_CACHE_KEY_"

func NewSysUserRoleCache(cacheInterface inter.CacheInterface) *SysUserRoleCache {
	return &SysUserRoleCache{cacheInterface: cacheInterface}
}

func (c *SysUserRoleCache) Get(key string) (model.SysUserRole, error) {
	var data model.SysUserRole
	err := c.cacheInterface.Get(UserRoleCacheKey+key, &data)
	if err != nil {
		zap.L().Error(UserRoleCacheKey+"Error getting key in cache", zap.String("key", key), zap.Error(err))
		return data, err
	}
	return data, nil
}

func (c *SysUserRoleCache) Set(key string, data model.SysUserRole) error {
	err := c.cacheInterface.Set(UserRoleCacheKey+key, &data, 7200*time.Second)
	if err != nil {
		zap.L().Error(UserRoleCacheKey+"Error setting key in cache", zap.String("key", key), zap.Error(err))
		return err
	}
	return nil
}

func (c *SysUserRoleCache) Delete(key string) error {
	err := c.cacheInterface.Del(UserRoleCacheKey + key)
	if err != nil {
		zap.L().Error(UserRoleCacheKey+"Error delete key in cache", zap.String("key", key), zap.Error(err))
		return err
	}
	return nil
}
