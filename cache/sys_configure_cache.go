/**
 * @Author: Nan
 * @Date: 2024/5/21 下午2:22
 */

package cache

import (
	"sweet-cms/inter"
	"sweet-cms/model"
)

type SysConfigureCache struct {
	*BasicCache[model.SysConfigure]
}

const ConfigureCacheKey = "CONFIGURE_CACHE_KEY_"

func NewSysConfigureCache(cacheInterface inter.CacheInterface) *SysConfigureCache {
	return &SysConfigureCache{
		BasicCache: NewBasicCache[model.SysConfigure](cacheInterface, ConfigureCacheKey),
	}
}

//func (sc *SysConfigureCache) Get(key string) (model.SysConfigure, error) {
//	var sysConfigure model.SysConfigure
//	err := sc.cacheInterface.Get(ConfigureCacheKey+key, &sysConfigure)
//	if err != nil {
//		zap.L().Error(ConfigureCacheKey+"Error getting key in cache", zap.String("key", key), zap.Error(err))
//		return sysConfigure, err
//	}
//	return sysConfigure, nil
//}
//
//func (sc *SysConfigureCache) Set(key string, value model.SysConfigure) error {
//	err := sc.cacheInterface.Set(ConfigureCacheKey+key, value, 7200*time.Second)
//	if err != nil {
//		zap.L().Error(ConfigureCacheKey+"Error setting key in cache", zap.String("key", key), zap.Error(err))
//		return err
//	}
//	return nil
//}
//
//func (sc *SysConfigureCache) Delete(key string) error {
//	return sc.cacheInterface.Del(ConfigureCacheKey + key)
//}
