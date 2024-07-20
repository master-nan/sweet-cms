/**
 * @Author: Nan
 * @Date: 2024/7/20 上午10:27
 */

package impl

import (
	"gorm.io/gorm"
	"sweet-cms/model"
)

type AccessLogRepositoryImpl struct {
	db *gorm.DB
	*BasicImpl
}

func NewAccessLogRepositoryImpl(db *gorm.DB) *AccessLogRepositoryImpl {
	return &AccessLogRepositoryImpl{
		db,
		NewBasicImpl(db, &model.AccessLog{}),
	}
}
