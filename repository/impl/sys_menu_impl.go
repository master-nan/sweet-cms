/**
 * @Author: Nan
 * @Date: 2024/7/19 上午11:27
 */

package impl

import (
	"gorm.io/gorm"
	"sweet-cms/model"
)

type SysMenuRepositoryImpl struct {
	db *gorm.DB
	*BasicImpl
}

func NewSysMenuRepositoryImpl(db *gorm.DB) *SysMenuRepositoryImpl {
	return &SysMenuRepositoryImpl{db, NewBasicImpl(db, &model.SysMenu{})}
}
func (s *SysMenuRepositoryImpl) GetMenuById(id int) (model.SysMenu, error) {
	var menu model.SysMenu
	err := s.db.Preload("MenuButtons").First(&menu, id).Error
	return menu, err
}

func (s *SysMenuRepositoryImpl) GetMenus() ([]model.SysMenu, error) {
	var menus []model.SysMenu
	err := s.db.Preload("MenuButtons").Find(&menus).Error
	return menus, err
}

// GetMenuUserPermissions 获取菜单的用户权限
func (s *SysMenuRepositoryImpl) GetMenuUserPermissions(menuId int) ([]model.SysUserMenuDataPermission, error) {
	var permissions []model.SysUserMenuDataPermission
	err := s.db.Where("menu_id = ?", menuId).Find(&permissions).Error
	return permissions, err
}

// GetMenuUsers 获取菜单的用户
func (s *SysMenuRepositoryImpl) GetMenuUsers(menuId int) ([]model.SysUser, error) {
	var menu model.SysMenu
	err := s.db.Preload("Permissions.User").First(&menu, menuId).Error
	if err != nil {
		return nil, err
	}
	var users []model.SysUser
	for _, permission := range menu.Permissions {
		users = append(users, permission.User)
	}
	return users, nil
}
