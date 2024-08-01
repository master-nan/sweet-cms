/**
 * @Author: Nan
 * @Date: 2024/7/25 下午11:05
 */

package service

import (
	"github.com/gin-gonic/gin"
	"sweet-cms/form/request"
	"sweet-cms/form/response"
	"sweet-cms/model"
	"sweet-cms/repository"
	"sweet-cms/utils"
)

type SysRoleService struct {
	sysRoleRepo           repository.SysRoleRepository
	sysRoleMenuRepo       repository.SysRoleMenuRepository
	sysRoleMenuButtonRepo repository.SysRoleMenuButtonRepository
	sf                    *utils.Snowflake
}

func NewSysRoleService(sysRoleRepo repository.SysRoleRepository, sysRoleMenuRepo repository.SysRoleMenuRepository, sysRoleMenuButtonRepo repository.SysRoleMenuButtonRepository, sf *utils.Snowflake) *SysRoleService {
	return &SysRoleService{
		sysRoleRepo,
		sysRoleMenuRepo,
		sysRoleMenuButtonRepo,
		sf,
	}
}

func (s *SysRoleService) GetRoleById(id int) (model.SysRole, error) {
	result, err := s.sysRoleRepo.WithPreload("Menus", "Buttons").FindById(id)
	if err != nil {
		return model.SysRole{}, err
	}
	data := result.(model.SysRole)
	return data, nil
}

func (s *SysRoleService) GetRoleList(basic request.Basic) (response.ListResult[model.SysRole], error) {
	return s.sysRoleRepo.GetRoleList(basic)
}

func (s *SysRoleService) CreateRole(ctx *gin.Context, role model.SysRole) error {
	id, err := s.sf.GenerateUniqueID()
	if err != nil {
		return err
	}
	role.Id = int(id)
	return s.sysRoleRepo.Create(s.sysRoleRepo.DBWithContext(ctx), role)
}

func (s *SysRoleService) UpdateRole(ctx *gin.Context, role model.SysRole) error {
	return s.sysRoleRepo.Update(s.sysRoleRepo.DBWithContext(ctx), role)
}

func (s *SysRoleService) DeleteRole(ctx *gin.Context, id int) error {
	return s.sysRoleRepo.DeleteById(s.sysRoleRepo.DBWithContext(ctx), id)
}

func (s *SysRoleService) GetRoleMenus(roleId int) ([]model.SysMenu, error) {
	menus, err := s.sysRoleMenuRepo.GetRoleMenus(roleId)
	if err != nil {
		return nil, err
	}
	return buildMenuTree(menus, 0), nil
}

func (s *SysRoleService) GetRoleMenuButtons(roleId, menuId int) ([]model.SysMenuButton, error) {
	return s.sysRoleMenuButtonRepo.GetRoleMenuButtons(roleId, menuId)
}
