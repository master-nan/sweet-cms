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
	err := s.db.Preload("Roles").Where("Roles.id = ?", roleId).Find(&menus).Error
	return menus, err
}
