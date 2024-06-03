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
}

func NewLogRepositoryImpl(db *gorm.DB) *LogRepositoryImpl {
	return &LogRepositoryImpl{
		db,
	}
}

func (lr *LogRepositoryImpl) CreateLoginLog(log model.LoginLog) (int, error) {
	err := lr.db.Omit("gmt_delete").Create(&log).Error
	if err != nil {
		return 0, err
	}
	return log.ID, err
}

func (lr *LogRepositoryImpl) CreateAccessLog(log model.AccessLog) (int, error) {
	err := lr.db.Omit("gmt_delete").Create(&log).Error
	if err != nil {
		return 0, err
	}
	return log.ID, err
}
