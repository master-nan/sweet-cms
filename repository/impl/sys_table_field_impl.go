/**
 * @Author: Nan
 * @Date: 2024/7/20 下午3:34
 */

package impl

import (
	"gorm.io/gorm"
	"sweet-cms/model"
)

type SysTableFieldRepositoryImpl struct {
	db *gorm.DB
	*BasicImpl
}

func NewSysTableFieldRepositoryImpl(db *gorm.DB) *SysTableFieldRepositoryImpl {
	return &SysTableFieldRepositoryImpl{
		db:        db,
		BasicImpl: NewBasicImpl(db, &model.SysTableField{}),
	}
}

func (s *SysTableFieldRepositoryImpl) GetTableFieldById(i int) (model.SysTableField, error) {
	var tableField model.SysTableField
	err := s.db.Where("id = ? ", i).First(&tableField).Error
	return tableField, err
}

func (s *SysTableFieldRepositoryImpl) GetTableFieldsByTableId(id int) ([]model.SysTableField, error) {
	var items []model.SysTableField
	err := s.db.Where("table_id = ?", id).Order("sequence").Find(&items).Error
	return items, err
}
