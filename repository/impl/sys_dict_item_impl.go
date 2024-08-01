/**
 * @Author: Nan
 * @Date: 2024/7/20 下午2:52
 */

package impl

import (
	"gorm.io/gorm"
	"sweet-cms/model"
)

type SysDictItemRepositoryImpl struct {
	db *gorm.DB
	*BasicImpl
}

func NewSysDictItemRepositoryImpl(db *gorm.DB) *SysDictItemRepositoryImpl {
	return &SysDictItemRepositoryImpl{
		db,
		NewBasicImpl(db, &model.SysDictItem{}),
	}
}

func (i *SysDictItemRepositoryImpl) GetSysDictItemsByDictId(id int) ([]model.SysDictItem, error) {
	var items []model.SysDictItem
	err := i.db.Where("dict_id = ?", id).Find(&items).Error
	return items, err
}
