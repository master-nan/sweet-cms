/**
 * @Author: Nan
 * @Date: 2024/5/21 下午2:27
 */

package sys

import (
	"sweet-cms/form/request/sys"
	"sweet-cms/global"
	"sweet-cms/model"
)

type ConfigureServer struct {
}

func NewConfigureServer() *ConfigureServer {
	return &ConfigureServer{}
}

func (cs *ConfigureServer) Query() (model.SysConfigure, error) {
	var data model.SysConfigure
	err := global.DB.First(&data).Error
	if err != nil {
		return model.SysConfigure{}, err
	}
	return data, nil
}

func (cs *ConfigureServer) Update(id int, data sys.ConfigureUpdateReq) error {
	err := global.DB.Model(model.SysConfigure{}).Where("id = ？", id).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}
