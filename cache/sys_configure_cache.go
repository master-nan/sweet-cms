/**
 * @Author: Nan
 * @Date: 2024/5/21 下午2:22
 */

package cache

import (
	"go.uber.org/zap"
	"sweet-cms/inter"
	"sweet-cms/model"
	"time"
)

type SysConfigureCache struct {
	cacheInterface inter.CacheInterface
}

const ConfigureCacheKey = "CONFIGURE_CACHE_KEY_"

func NewSysConfigureCache(cacheInterface inter.CacheInterface) *SysConfigureCache {
	return &SysConfigureCache{
		cacheInterface: cacheInterface,
	}
}

func (sc *SysConfigureCache) Get(key string) (model.SysConfigure, error) {
	var sysConfigure model.SysConfigure
	err := sc.cacheInterface.Get(ConfigureCacheKey+key, &sysConfigure)
	if err != nil {
		zap.L().Error("Error setting key in cache", zap.String("key", key), zap.Error(err))
		return sysConfigure, err
	}
	return sysConfigure, nil
}

func (sc *SysConfigureCache) Set(key string, value model.SysConfigure) error {
	err := sc.cacheInterface.Set(ConfigureCacheKey+key, value, 7200*time.Second)
	if err != nil {
		return err
	}
	return nil
}

func (sc *SysConfigureCache) Delete(key string) error {
	err := sc.cacheInterface.Del(ConfigureCacheKey + key)
	if err != nil {
		return err
	}
	return nil
}
