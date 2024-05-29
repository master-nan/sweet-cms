/**
 * @Author: Nan
 * @Date: 2024/5/21 下午2:22
 */

package cache

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"sweet-cms/inter"
	"sweet-cms/model"
	"sweet-cms/service"
	"time"
)

type SysConfigureCache struct {
	ConfigureServer *service.ConfigureService
	CacheInterface  inter.CacheInterface
}

const ConfigureCacheKey = "CONFIGURE_CACHE_KEY_"

func NewSysConfigureCache(configureServer *service.ConfigureService, cacheInterface inter.CacheInterface) *SysConfigureCache {
	return &SysConfigureCache{
		ConfigureServer: configureServer,
		CacheInterface:  cacheInterface,
	}
}

func (sc *SysConfigureCache) Get(key string) (model.SysConfigure, error) {
	var sysConfigure model.SysConfigure
	err := sc.CacheInterface.Get(ConfigureCacheKey+key, &sysConfigure)
	if err != nil {
		if errors.Is(err, inter.ErrCacheMiss) {
			sysConfigure, err := sc.ConfigureServer.Query()
			if err != nil {
				return model.SysConfigure{}, err
			}
			err = sc.Set(key, sysConfigure)
			if err != nil {
				zap.S().Error("Error setting key in cache", zap.String("key", key), zap.Error(err))
			}
			return sysConfigure, nil
		} else {
			return sysConfigure, err
		}
	}
	return sysConfigure, nil
}

func (sc *SysConfigureCache) Set(key string, value model.SysConfigure) error {
	err := sc.CacheInterface.Set(ConfigureCacheKey+key, value, 7200*time.Second)
	if err != nil {
		return err
	}
	return nil
}
