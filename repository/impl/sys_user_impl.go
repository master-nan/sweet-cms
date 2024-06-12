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

func (s *SysUserRepositoryImpl) GetByUserName(username string) (model.SysUser, error) {
	var user model.SysUser
	result := s.db.Where(&model.SysUser{UserName: username}).First(&user)
	return user, result.Error
}

func (s *SysUserRepositoryImpl) GetByUserId(id int) (model.SysUser, error) {
	var user model.SysUser
	result := s.db.Where("id = ?", id).First(&user)
	return user, result.Error
}
