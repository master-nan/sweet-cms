/**
 * @Author: Nan
 * @Date: 2024/6/3 下午4:31
 */

package impl

import (
	"gorm.io/gorm"
	"sweet-cms/model"
)

type LoginLogRepositoryImpl struct {
	db *gorm.DB
	*BasicImpl
}

func NewLoginLogRepositoryImpl(db *gorm.DB) *LoginLogRepositoryImpl {
	return &LoginLogRepositoryImpl{
		db,
		NewBasicImpl(db, &model.LoginLog{}),
	}
}
