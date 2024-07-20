/**
 * @Author: Nan
 * @Date: 2024/7/19 下午5:07
 */

package impl

import (
	"gorm.io/gorm"
	"sweet-cms/model"
)

type SysRoleRepositoryImpl struct {
	db *gorm.DB
	*BasicImpl
}

func NewSysRoleRepositoryImpl(db *gorm.DB) *SysRoleRepositoryImpl {
	return &SysRoleRepositoryImpl{db, NewBasicImpl(db, &model.SysRole{})}
}

func (s *SysRoleRepositoryImpl) GetRoleById(roleId int) (model.SysRole, error) {
	var role model.SysRole
	err := s.db.Preload("Menus").Preload("Buttons").First(&role, roleId).Error
	return role, err
}

func (s *SysRoleRepositoryImpl) GetRoles() ([]model.SysRole, error) {
	var roles []model.SysRole
	err := s.db.Preload("Menus").Preload("Buttons").Find(&roles).Error
	return roles, err
}

func (s *SysRoleRepositoryImpl) GetRoleMenus(roleId int) ([]model.SysMenu, error) {
	var role model.SysRole
	err := s.db.Preload("Menus").First(&role, roleId).Error
	if err != nil {
		return nil, err
	}
	return role.Menus, nil
}

func (s *SysRoleRepositoryImpl) GetRoleButtons(roleId int) ([]model.SysMenuButton, error) {
	var role model.SysRole
	err := s.db.Preload("Buttons").First(&role, roleId).Error
	if err != nil {
		return nil, err
	}
	return role.Buttons, nil
}
