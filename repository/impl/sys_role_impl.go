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

func NewSysRoleRepositoryImpl(db *gorm.DB, basicImpl *BasicImpl) *SysRoleRepositoryImpl {
	return &SysRoleRepositoryImpl{db: db, BasicImpl: basicImpl}
}

func (s *SysRoleRepositoryImpl) CreateRole(tx *gorm.DB, role model.SysRole) error {
	return tx.Create(&role).Error
}

func (s *SysRoleRepositoryImpl) UpdateRole(tx *gorm.DB, role model.SysRole) error {
	return tx.Save(&role).Error
}

func (s *SysRoleRepositoryImpl) DeleteRole(tx *gorm.DB, roleId int) error {
	return tx.Delete(&model.SysRole{}, roleId).Error
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
