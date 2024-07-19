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

func NewSysRoleMenuRepositoryImpl(db *gorm.DB, basicImpl *BasicImpl) *SysRoleMenuRepositoryImpl {
	return &SysRoleMenuRepositoryImpl{db: db, BasicImpl: basicImpl}
}

func (s *SysRoleMenuRepositoryImpl) CreateRoleMenu(tx *gorm.DB, roleMenu model.SysRoleMenu) error {
	return tx.Create(&roleMenu).Error
}

func (s *SysRoleMenuRepositoryImpl) DeleteRoleMenu(tx *gorm.DB, id int) error {
	return tx.Where("id = ? ", id).Delete(&model.SysRoleMenu{}).Error
}

func (s *SysRoleMenuRepositoryImpl) GetRoleMenus(roleId int) ([]model.SysMenu, error) {
	var menus []model.SysMenu
	err := s.db.Preload("Roles").Where("Roles.id = ?", roleId).Find(&menus).Error
	return menus, err
}
