/**
 * @Author: Nan
 * @Date: 2024/5/25 下午2:24
 */

package impl

import (
	"sweet-cms/global"
	"sweet-cms/model"
)

type SysDictRepositoryImpl struct {
}

func NewSysDictRepositoryImpl() *SysDictRepositoryImpl {
	return &SysDictRepositoryImpl{}
}

func (i *SysDictRepositoryImpl) GetSysDictById(id int) (model.SysDict, error) {
	var sysDict model.SysDict
	err := global.DB.Preload("DictItems").Where("id = ?", id).First(&sysDict).Error
	return sysDict, err
}

func (i *SysDictRepositoryImpl) GetSysDictList() ([]model.SysDict, int, error) {
	return nil, 0, nil
}

func (i *SysDictRepositoryImpl) UpdateSysDict(*model.SysDict) error {
	return nil
}

func (i *SysDictRepositoryImpl) InsertSysDict(*model.SysDict) error {
	return nil
}

func (i *SysDictRepositoryImpl) DeleteSysDictById(id int) error {
	return nil
}

func (i *SysDictRepositoryImpl) GetSysDictByCode(code int) (model.SysDict, error) {
	return model.SysDict{}, nil
}

func (i *SysDictRepositoryImpl) GetSysDictItemById(id int) (model.SysDictItem, error) {
	return model.SysDictItem{}, nil
}

func (i *SysDictRepositoryImpl) GetSysDictItemsByDictId(id int) ([]model.SysDictItem, int, error) {
	return nil, 0, nil
}
func (i *SysDictRepositoryImpl) UpdateSysDictItem(*model.SysDictItem) error {
	return nil
}
func (i *SysDictRepositoryImpl) InsertSysDictItem(*model.SysDictItem) error {
	return nil
}
func (i *SysDictRepositoryImpl) DeleteSysDictItemById(id int) error {
	return nil
}
