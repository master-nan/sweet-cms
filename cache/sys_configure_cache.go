/**
 * @Author: Nan
 * @Date: 2024/5/21 下午2:22
 */

package cache

import (
	"github.com/pkg/errors"
	"sweet-cms/inter"
	"sweet-cms/model"
	"sweet-cms/server/sys"
	"time"
)

type SysConfigureCache struct {
	ConfigureServer *sys.ConfigureServer
	CacheInterface  inter.CacheInterface
}

func NewSysConfigureCache(configureServer *sys.ConfigureServer, cacheInterface inter.CacheInterface) *SysConfigureCache {
	return &SysConfigureCache{
		ConfigureServer: configureServer,
		CacheInterface:  cacheInterface,
	}
}

func (sc *SysConfigureCache) Get(key string) (model.SysConfigure, error) {
	var sysConfigure model.SysConfigure
	err := sc.CacheInterface.Get(key, &sysConfigure)
	if err != nil {
		if errors.Is(err, inter.ErrCacheMiss) {
			sysConfigure, err := sc.ConfigureServer.Query()
			if err != nil {
				return model.SysConfigure{}, err
			}
			return sysConfigure, nil
		} else {
			return sysConfigure, err
		}
	}
	return sysConfigure, nil
}

func (sc *SysConfigureCache) Set(key string, value model.SysConfigure) error {
	err := sc.CacheInterface.Set(key, value, 7200*time.Second)
	if err != nil {
		return err
	}
	return nil
}
