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
	*BasicImpl
}

func NewSysConfigureRepositoryImpl(db *gorm.DB, basicImpl *BasicImpl) *SysConfigureRepositoryImpl {
	return &SysConfigureRepositoryImpl{db, basicImpl}
}

func (c *SysConfigureRepositoryImpl) GetSysConfigure() (model.SysConfigure, error) {
	var data model.SysConfigure
	err := c.db.First(&data).Error
	return data, err
}

func (c *SysConfigureRepositoryImpl) UpdateSysConfigure(d model.SysConfigure) error {
	return c.db.Save(d).Error
}
