/**
 * @Author: Nan
 * @Date: 2024/6/3 下午4:31
 */

package impl

import (
	"gorm.io/gorm"
	"sweet-cms/model"
)

type LogRepositoryImpl struct {
	db *gorm.DB
	*BasicImpl
}

func NewLogRepositoryImpl(db *gorm.DB, basicImpl *BasicImpl) *LogRepositoryImpl {
	return &LogRepositoryImpl{
		db,
		basicImpl,
	}
}

func (lr *LogRepositoryImpl) CreateLoginLog(log model.LoginLog) error {
	err := lr.db.Omit("gmt_delete").Create(&log).Error
	return err
}

func (lr *LogRepositoryImpl) CreateAccessLog(log model.AccessLog) error {
	err := lr.db.Omit("gmt_delete").Create(&log).Error
	return err
}
