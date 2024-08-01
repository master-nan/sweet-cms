/**
 * @Author: Nan
 * @Date: 2024/7/19 下午5:59
 */

package impl

import (
	"gorm.io/gorm"
	"sweet-cms/model"
)

type SysRoleMenuRepositoryImpl struct {
	db *gorm.DB
	*BasicImpl
}

func NewSysRoleMenuRepositoryImpl(db *gorm.DB) *SysRoleMenuRepositoryImpl {
	return &SysRoleMenuRepositoryImpl{db, NewBasicImpl(db, &model.SysRoleMenu{})}
}

func (s *SysRoleMenuRepositoryImpl) GetRoleMenus(roleId int) ([]model.SysMenu, error) {
	var menus []model.SysMenu
	err := s.db.Preload("Roles", "Roles.id = ?", roleId).Find(&menus).Error
	return menus, err
}

// GetRoleMenusByRoleIds 获取角色的所有菜单
func (s *SysRoleMenuRepositoryImpl) GetRoleMenusByRoleIds(roleIds []int) ([]model.SysMenu, error) {
	var menus []model.SysMenu
	err := s.db.Preload("Roles", "Roles.id IN ?", roleIds).Find(&menus).Error
	return menus, err
}

func (s *SysRoleMenuRepositoryImpl) DeleteRoleMenuByRoleIdAndMenuId(tx *gorm.DB, roleId, menuId int) error {
	if tx == nil {
		tx = s.db
	}
	return tx.Where("role_id = ? and menu_id", roleId, menuId).Delete(&model.SysRoleMenu{}).Error
}
