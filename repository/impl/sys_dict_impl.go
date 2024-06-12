/**
 * @Author: Nan
 * @Date: 2024/5/25 下午2:24
 */

package impl

import (
	"gorm.io/gorm"
	"sweet-cms/form/request"
	"sweet-cms/model"
	"sweet-cms/repository"
	"sweet-cms/utils"
)

type SysDictRepositoryImpl struct {
	db *gorm.DB
}

func NewSysDictRepositoryImpl(db *gorm.DB) *SysDictRepositoryImpl {
	return &SysDictRepositoryImpl{
		db,
	}
}

func (i *SysDictRepositoryImpl) GetSysDictById(id int) (model.SysDict, error) {
	var sysDict model.SysDict
	err := i.db.Preload("DictItems").Where("id = ?", id).First(&sysDict).Error
	return sysDict, err
}

func (i *SysDictRepositoryImpl) GetSysDictList(basic request.Basic) (repository.SysDictListResult, error) {
	var repo repository.SysDictListResult
	query := utils.BuildQuery(i.db, basic)
	var sysDictList []model.SysDict
	var total int64 = 0
	err := query.Find(&sysDictList).Limit(-1).Offset(-1).Count(&total).Error
	repo.Data = sysDictList
	repo.Total = int(total)
	return repo, err
}

func (i *SysDictRepositoryImpl) InsertSysDict(d model.SysDict) error {
	result := i.db.Create(&d)
	return result.Error
}

func (i *SysDictRepositoryImpl) UpdateSysDict(d request.DictUpdateReq) error {
	return i.db.Model(model.SysDict{}).Updates(&d).Error
}

func (i *SysDictRepositoryImpl) DeleteSysDictById(id int) error {
	err := i.db.Where("id = ?", id).Delete(model.SysDict{}).Error
	return err
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

func (i *SysDictRepositoryImpl) UpdateSysDictItem(d request.DictItemUpdateReq) error {
	return i.db.Model(model.SysDictItem{}).Updates(&d).Error
}

func (i *SysDictRepositoryImpl) InsertSysDictItem(d model.SysDictItem) error {
	result := i.db.Create(&d)
	return result.Error
}

func (i *SysDictRepositoryImpl) DeleteSysDictItemById(id int) error {
	err := i.db.Delete(model.SysDictItem{}, id).Error
	return err
}
