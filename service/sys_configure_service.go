/**
 * @Author: Nan
 * @Date: 2024/5/21 下午2:27
 */

package service

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"sweet-cms/cache"
	"sweet-cms/form/request"
	"sweet-cms/inter"
	"sweet-cms/model"
	"sweet-cms/repository"
)

type SysConfigureService struct {
	sysConfigureRepo  repository.SysConfigureRepository
	sysConfigureCache *cache.SysConfigureCache
}

func NewSysConfigureService(sysConfigureRepo repository.SysConfigureRepository, sysConfigureCache *cache.SysConfigureCache) *SysConfigureService {
	return &SysConfigureService{
		sysConfigureRepo,
		sysConfigureCache,
	}
}

func (cs *SysConfigureService) Query() (model.SysConfigure, error) {
	var data model.SysConfigure
	data, err := cs.sysConfigureCache.Get("")
	if err != nil {
		data, e := cs.sysConfigureRepo.GetSysConfigure()
		if e != nil {
			return data, e
		}
		if errors.Is(err, inter.ErrCacheMiss) {
			err = cs.sysConfigureCache.Set("", data)
			if err != nil {
				zap.L().Error("Failed to cache sysConfigure set: %s", zap.Error(err))
			}
		}
	}
	return data, err
}

func (cs *SysConfigureService) Update(id int, data request.ConfigureUpdateReq) error {
	var d model.SysConfigure
	d.Id = id
	err := cs.sysConfigureRepo.UpdateSysConfigure(d)
	if err != nil {
		return err
	}
	err = cs.sysConfigureCache.Delete("")
	if err != nil {
		zap.L().Error("Failed to cache sysConfigure delete: %s", zap.Error(err))
	}
	return nil
}
