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

func NewSysUserMenuDataPermissionRepositoryImpl(db *gorm.DB) *SysUserMenuDataPermissionRepositoryImpl {
	return &SysUserMenuDataPermissionRepositoryImpl{db, NewBasicImpl(db, &model.SysUserMenuDataPermission{})}
}

func (s *SysUserMenuDataPermissionRepositoryImpl) GetUserMenuPermissionsByUserId(userId int) ([]model.SysUserMenuDataPermission, error) {
	var permissions []model.SysUserMenuDataPermission
	err := s.db.Where("user_id = ?", userId).Find(&permissions).Error
	return permissions, err
}

// GetUserMenuPermissions 获取用户在指定菜单下的数据权限
func (s *SysUserMenuDataPermissionRepositoryImpl) GetUserMenuPermissions(menuId int) ([]model.SysUserMenuDataPermission, error) {
	var permissions []model.SysUserMenuDataPermission
	err := s.db.Where("menu_id = ?", menuId).Find(&permissions).Error
	return permissions, err
}
