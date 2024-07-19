/**
 * @Author: Nan
 * @Date: 2024/7/19 下午4:49
 */

package impl

import (
	"gorm.io/gorm"
	"sweet-cms/model"
)

type SysUserMenuDataPermissionRepositoryImpl struct {
	db *gorm.DB
	*BasicImpl
}

func NewSysUserMenuDataPermissionRepositoryImpl(db *gorm.DB, basicImpl *BasicImpl) *SysUserMenuDataPermissionRepositoryImpl {
	return &SysUserMenuDataPermissionRepositoryImpl{db, basicImpl}
}

func (s *SysUserMenuDataPermissionRepositoryImpl) GetUserMenuPermissionsByUserId(userId int) ([]model.SysUserMenuDataPermission, error) {
	var permissions []model.SysUserMenuDataPermission
	err := s.db.Where("user_id = ?", userId).Find(&permissions).Error
	return permissions, err
}

func (s *SysUserMenuDataPermissionRepositoryImpl) CreateUserMenuPermission(tx *gorm.DB, permission model.SysUserMenuDataPermission) error {
	return tx.Create(&permission).Error
}

func (s *SysUserMenuDataPermissionRepositoryImpl) UpdateUserMenuPermission(tx *gorm.DB, permission model.SysUserMenuDataPermission) error {
	return tx.Save(&permission).Error
}

func (s *SysUserMenuDataPermissionRepositoryImpl) DeleteUserMenuPermission(tx *gorm.DB, id int) error {
	return tx.Delete(&model.SysUserMenuDataPermission{}, id).Error
}
