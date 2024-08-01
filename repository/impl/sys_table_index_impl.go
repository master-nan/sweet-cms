/**
 * @Author: Nan
 * @Date: 2024/7/20 下午3:55
 */

package impl

import (
	"gorm.io/gorm"
	"sweet-cms/model"
)

type SysTableIndexRepositoryImpl struct {
	db *gorm.DB
	*BasicImpl
}

func NewSysTableIndexRepositoryImpl(db *gorm.DB) *SysTableIndexRepositoryImpl {
	return &SysTableIndexRepositoryImpl{
		db,
		NewBasicImpl(db, &model.SysTableIndex{}),
	}
}

func (s *SysTableIndexRepositoryImpl) GetTableIndexesByTableId(id int) ([]model.SysTableIndex, error) {
	var indexes []model.SysTableIndex
	err := s.db.Where("table_id = ?", id).Find(&indexes).Error
	return indexes, err
}
