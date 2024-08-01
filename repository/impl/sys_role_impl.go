/**
 * @Author: Nan
 * @Date: 2024/7/19 下午5:07
 */

package impl

import (
	"gorm.io/gorm"
	"sweet-cms/form/request"
	"sweet-cms/form/response"
	"sweet-cms/model"
)

type SysRoleRepositoryImpl struct {
	db *gorm.DB
	*BasicImpl
}

func NewSysRoleRepositoryImpl(db *gorm.DB) *SysRoleRepositoryImpl {
	return &SysRoleRepositoryImpl{db, NewBasicImpl(db, &model.SysRole{})}
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

func (s *SysRoleRepositoryImpl) GetRoleList(basic request.Basic) (response.ListResult[model.SysRole], error) {
	var repo response.ListResult[model.SysRole]
	var sysRoleList []model.SysRole
	total, err := s.PaginateAndCountAsync(basic, &sysRoleList)
	repo.Data = sysRoleList
	repo.Total = int(total)
	return repo, err
}
