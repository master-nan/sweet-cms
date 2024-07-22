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
