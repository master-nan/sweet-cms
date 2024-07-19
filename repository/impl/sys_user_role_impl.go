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

func NewSysUserRoleRepositoryImpl(db *gorm.DB, basicImpl *BasicImpl) *SysUserRoleRepositoryImpl {
	return &SysUserRoleRepositoryImpl{db, basicImpl}
}

func (s *SysUserRoleRepositoryImpl) CreateUserRole(tx *gorm.DB, userRole model.SysUserRole) error {
	return tx.Create(&userRole).Error
}

func (s *SysUserRoleRepositoryImpl) DeleteUserRole(tx *gorm.DB, id int) error {
	return tx.Where("id = ?", id).Delete(&model.SysUserRole{}).Error
}

func (s *SysUserRoleRepositoryImpl) GetUserRoles(userId int) ([]model.SysRole, error) {
	var roles []model.SysRole
	err := s.db.Preload("Users", "id = ?", userId).Find(&roles).Error
	return roles, err
}
