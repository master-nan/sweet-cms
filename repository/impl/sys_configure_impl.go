/**
 * @Author: Nan
 * @Date: 2024/6/3 下午2:52
 */

package impl

import (
	"gorm.io/gorm"
	"sweet-cms/model"
)

type SysConfigureRepositoryImpl struct {
	db *gorm.DB
}

func NewSysConfigureRepositoryImpl(db *gorm.DB) *SysConfigureRepositoryImpl {
	return &SysConfigureRepositoryImpl{db: db}
}

func (c *SysConfigureRepositoryImpl) GetSysConfigure() (model.SysConfigure, error) {
	var data model.SysConfigure
	err := c.db.First(&data).Error
	return data, err
}

func (c *SysConfigureRepositoryImpl) UpdateSysConfigure(d model.SysConfigure) error {
	return c.db.Save(d).Error
}
