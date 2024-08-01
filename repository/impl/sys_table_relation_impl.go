/**
 * @Author: Nan
 * @Date: 2024/7/20 下午3:50
 */

package impl

import (
	"gorm.io/gorm"
	"sweet-cms/model"
)

type SysTableRelationRepositoryImpl struct {
	db *gorm.DB
	*BasicImpl
}

func NewSysTableRelationRepositoryImpl(db *gorm.DB) *SysTableRelationRepositoryImpl {
	return &SysTableRelationRepositoryImpl{
		db,
		NewBasicImpl(db, &model.SysTableRelation{}),
	}
}

func (s *SysTableRelationRepositoryImpl) GetTableRelationsByTableId(i int) ([]model.SysTableRelation, error) {
	var relations []model.SysTableRelation
	err := s.db.Where("table_id = ?", i).First(&relations).Error
	return relations, err
}
