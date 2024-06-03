/**
 * @Author: Nan
 * @Date: 2024/5/21 下午2:27
 */

package service

import (
	"go.uber.org/zap"
	"sweet-cms/cache"
	"sweet-cms/form/request"
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
		data, err = cs.sysConfigureRepo.GetSysConfigure()
		if err != nil {
			return data, err
		}
		err = cs.sysConfigureCache.Set("", data)
		if err != nil {
			zap.S().Errorf("Failed to cache sysConfigure set: %s", err.Error())
		}
	}
	return data, err
}

func (cs *SysConfigureService) Update(id int, data request.ConfigureUpdateReq) error {
	var d model.SysConfigure
	d.ID = id
	err := cs.sysConfigureRepo.UpdateSysConfigure(d)
	if err != nil {
		return err
	}
	err = cs.sysConfigureCache.Delete("")
	if err != nil {
		zap.S().Errorf("Failed to cache sysConfigure delete: %s", err.Error())
	}
	return nil
}
