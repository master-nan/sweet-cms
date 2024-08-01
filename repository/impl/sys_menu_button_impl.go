/**
 * @Author: Nan
 * @Date: 2024/8/1 下午10:37
 */

package impl

import (
	"gorm.io/gorm"
	"sweet-cms/model"
)

type SysMenuButtonRepositoryImpl struct {
	db *gorm.DB
	*BasicImpl
}

func NewSysMenuButtonRepositoryImpl(db *gorm.DB) *SysMenuButtonRepositoryImpl {
	return &SysMenuButtonRepositoryImpl{
		db,
		NewBasicImpl(db, &model.SysMenuButton{}),
	}
}

func (s SysMenuButtonRepositoryImpl) GetMenuButtonsByMenuId(id int) ([]model.SysMenuButton, error) {
	var items []model.SysMenuButton
	err := s.db.Where("dict_id = ?", id).Find(&items).Error
	return items, err
}
