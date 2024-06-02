/**
 * @Author: Nan
 * @Date: 2024/5/21 下午2:27
 */

package service

import (
	"sweet-cms/form/request"
	"sweet-cms/global"
	"sweet-cms/model"
)

type ConfigureService struct {
}

func NewConfigureService() *ConfigureService {
	return &ConfigureService{}
}

func (cs *ConfigureService) Query() (model.SysConfigure, error) {
	var data model.SysConfigure
	err := global.DB.First(&data).Error
	if err != nil {
		return model.SysConfigure{}, err
	}
	return data, nil
}

func (cs *ConfigureService) Update(id int, data request.ConfigureUpdateReq) error {
	err := global.DB.Model(model.SysConfigure{}).Where("id = ？", id).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}
