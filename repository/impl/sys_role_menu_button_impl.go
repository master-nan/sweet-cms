/**
 * @Author: Nan
 * @Date: 2024/7/19 下午6:00
 */

package impl

import (
	"gorm.io/gorm"
	"sweet-cms/model"
)

type SysRoleMenuButtonRepositoryImpl struct {
	db *gorm.DB
	*BasicImpl
}

func NewSysRoleMenuButtonRepositoryImpl(db *gorm.DB) *SysRoleMenuButtonRepositoryImpl {
	return &SysRoleMenuButtonRepositoryImpl{db, NewBasicImpl(db, &model.SysRoleMenuButton{})}
}

func (s *SysRoleMenuButtonRepositoryImpl) GetRoleMenuButtons(roleId, menuId int) ([]model.SysMenuButton, error) {
	var buttons []model.SysMenuButton
	err := s.db.Preload("Roles", "id = ?", roleId).
		Preload("Menus", "id = ?", menuId).
		Find(&buttons).Error
	return buttons, err
}
