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
	"sweet-cms/repository/util"
)

type SysDictRepositoryImpl struct {
	db *gorm.DB
	*BasicImpl
}

func NewSysDictRepositoryImpl(db *gorm.DB) *SysDictRepositoryImpl {
	return &SysDictRepositoryImpl{
		db,
		NewBasicImpl(db, &model.SysDict{}),
	}
}

func (i *SysDictRepositoryImpl) GetSysDictList(basic request.Basic) (response.ListResult[model.SysDict], error) {
	var repo response.ListResult[model.SysDict]
	query := util.ExecuteQuery(i.db, basic)
	var sysDictList []model.SysDict
	var total int64 = 0
	err := query.Find(&sysDictList).Limit(-1).Offset(-1).Count(&total).Error
	repo.Data = sysDictList
	repo.Total = int(total)
	return repo, err
}

func (i *SysDictRepositoryImpl) GetSysDictByCode(code string) (model.SysDict, error) {
	var sysDict model.SysDict
	err := i.db.Preload("DictItems").Where("dict_code = ?", code).First(&sysDict).Error
	return sysDict, err
}
