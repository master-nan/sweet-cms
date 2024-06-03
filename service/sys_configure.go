/**
 * @Author: Nan
 * @Date: 2024/5/21 下午2:27
 */

package service

import (
	"sweet-cms/form/request"
	"sweet-cms/model"
	"sweet-cms/repository"
)

type SysConfigureService struct {
	sysConfigureRepo repository.SysConfigureRepository
}

func NewSysConfigureService(s repository.SysConfigureRepository) *SysConfigureService {
	return &SysConfigureService{
		sysConfigureRepo: s,
	}
}

func (cs *SysConfigureService) Query() (model.SysConfigure, error) {
	var data model.SysConfigure
	data, err := cs.sysConfigureRepo.GetSysConfigure()
	return data, err
}

func (cs *SysConfigureService) Update(id int, data request.ConfigureUpdateReq) error {
	var d model.SysConfigure
	d.ID = id
	err := cs.sysConfigureRepo.UpdateSysConfigure(d)
	return err
}
