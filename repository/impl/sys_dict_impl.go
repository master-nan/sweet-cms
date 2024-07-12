/**
 * @Author: Nan
 * @Date: 2024/5/25 下午2:24
 */

package impl

import (
	"gorm.io/gorm"
	"sweet-cms/form/request"
	"sweet-cms/form/response"
	"sweet-cms/model"
	"sweet-cms/utils"
)

type SysDictRepositoryImpl struct {
	db *gorm.DB
	*BasicImpl
}

func NewSysDictRepositoryImpl(db *gorm.DB, basicImpl *BasicImpl) *SysDictRepositoryImpl {
	return &SysDictRepositoryImpl{
		db,
		basicImpl,
	}
}

func (i *SysDictRepositoryImpl) GetSysDictById(id int) (model.SysDict, error) {
	var sysDict model.SysDict
	err := i.db.Preload("DictItems").Where("id = ?", id).First(&sysDict).Error
	return sysDict, err
}

func (i *SysDictRepositoryImpl) GetSysDictList(basic request.Basic) (response.ListResult[model.SysDict], error) {
	var repo response.ListResult[model.SysDict]
	query := utils.ExecuteQuery(i.db, basic)
	var sysDictList []model.SysDict
	var total int64 = 0
	err := query.Find(&sysDictList).Limit(-1).Offset(-1).Count(&total).Error
	repo.Data = sysDictList
	repo.Total = int(total)
	return repo, err
}

func (i *SysDictRepositoryImpl) InsertSysDict(tx *gorm.DB, d model.SysDict) error {
	return tx.Create(&d).Error
}

func (i *SysDictRepositoryImpl) UpdateSysDict(tx *gorm.DB, d request.DictUpdateReq) error {
	return tx.Model(model.SysDict{}).Updates(&d).Error
}

func (i *SysDictRepositoryImpl) DeleteSysDictById(tx *gorm.DB, id int) error {
	return tx.Where("id = ?", id).Delete(model.SysDict{}).Error
}

func (i *SysDictRepositoryImpl) GetSysDictByCode(code string) (model.SysDict, error) {
	var sysDict model.SysDict
	err := i.db.Preload("DictItems").Where("dict_code = ?", code).First(&sysDict).Error
	return sysDict, err
}

func (i *SysDictRepositoryImpl) GetSysDictItemById(id int) (model.SysDictItem, error) {
	var item model.SysDictItem
	err := i.db.Where("id = ?", id).First(&item).Error
	return item, err
}

func (i *SysDictRepositoryImpl) GetSysDictItemsByDictId(id int) ([]model.SysDictItem, error) {
	var items []model.SysDictItem
	err := i.db.Where("dict_id = ?", id).Find(&items).Error
	return items, err
}

func (i *SysDictRepositoryImpl) UpdateSysDictItem(tx *gorm.DB, d request.DictItemUpdateReq) error {
	return tx.Model(model.SysDictItem{}).Updates(&d).Error
}

func (i *SysDictRepositoryImpl) InsertSysDictItem(tx *gorm.DB, d model.SysDictItem) error {
	return tx.Create(&d).Error
}

func (i *SysDictRepositoryImpl) DeleteSysDictItemById(tx *gorm.DB, id int) error {
	return tx.Delete(model.SysDictItem{}, id).Error
}
