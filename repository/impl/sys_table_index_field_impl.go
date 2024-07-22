/**
 * @Author: Nan
 * @Date: 2024/7/22 上午10:19
 */

package impl

import (
	"gorm.io/gorm"
	"sweet-cms/model"
)

type SysTableIndexFieldRepositoryImpl struct {
	db *gorm.DB
	*BasicImpl
}

func NewSysTableIndexFieldRepositoryImpl(db *gorm.DB) *SysTableIndexFieldRepositoryImpl {
	return &SysTableIndexFieldRepositoryImpl{
		db,
		NewBasicImpl(db, &model.SysTableIndexField{}),
	}
}

// DeleteTableIndexFieldByIndexId 根据单个indexId删除中间表字段
func (s *SysTableIndexFieldRepositoryImpl) DeleteTableIndexFieldByIndexId(tx *gorm.DB, id int) error {
	return tx.Where("index_id = ?", id).Delete(model.SysTableIndexField{}).Error
}

// DeleteTableIndexFieldByIndexIds 根据所有indexId删除中间表字段
func (s *SysTableIndexFieldRepositoryImpl) DeleteTableIndexFieldByIndexIds(tx *gorm.DB, ids []int) error {
	return tx.Where("index_id in ?", ids).Delete(model.SysTableIndexField{}).Error
}
