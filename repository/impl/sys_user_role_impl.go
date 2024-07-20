/**
 * @Author: Nan
 * @Date: 2024/7/19 下午5:57
 */

package impl

import (
	"gorm.io/gorm"
	"sweet-cms/model"
)

type SysUserRoleRepositoryImpl struct {
	db *gorm.DB
	*BasicImpl
}

func NewSysUserRoleRepositoryImpl(db *gorm.DB) *SysUserRoleRepositoryImpl {
	return &SysUserRoleRepositoryImpl{db, NewBasicImpl(db, &model.SysUserRole{})}
}

func (s *SysUserRoleRepositoryImpl) GetUserRoles(userId int) ([]model.SysRole, error) {
	var roles []model.SysRole
	err := s.db.Preload("Users", "id = ?", userId).Find(&roles).Error
	return roles, err
}
