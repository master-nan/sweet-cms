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

func (i *SysDictItemRepositoryImpl) GetSysDictItemById(id int) (model.SysDictItem, error) {
	var item model.SysDictItem
	err := i.db.Where("id = ?", id).First(&item).Error
	return item, err
}

func (i *SysDictItemRepositoryImpl) GetSysDictItemsByDictId(id int) ([]model.SysDictItem, error) {
	var items []model.SysDictItem
	err := i.db.Where("dict_id = ?", id).Find(&items).Error
	return items, err
}
