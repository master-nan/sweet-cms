/**
 * @Author: Nan
 * @Date: 2024/6/3 下午6:08
 */

package impl

import (
	"gorm.io/gorm"
	"sweet-cms/model"
)

type SysUserRepositoryImpl struct {
	db *gorm.DB
}

func NewSysUserRepositoryImpl(db *gorm.DB) *SysUserRepositoryImpl {
	return &SysUserRepositoryImpl{db}
}

func (s *SysUserRepositoryImpl) GetSysUserByUserName(username string) (model.SysUser, error) {
	var user model.SysUser
	result := s.db.Where(&model.SysUser{UserName: username}).First(&user)
	return user, result.Error
}
