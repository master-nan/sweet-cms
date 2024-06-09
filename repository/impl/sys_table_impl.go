/**
 * @Author: Nan
 * @Date: 2024/6/10 上午12:16
 */

package impl

import "gorm.io/gorm"

type SysTableRepositoryImpl struct {
	db *gorm.DB
}

func NewSysTableRepositoryImpl(db *gorm.DB) *SysTableRepositoryImpl {
	return &SysTableRepositoryImpl{
		db,
	}
}
