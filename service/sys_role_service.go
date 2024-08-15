/**
 * @Author: Nan
 * @Date: 2024/7/25 下午11:05
 */

package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"gorm.io/gorm"
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.SysRole{}, nil
		}
		return model.SysRole{}, err
	}
	data := result.(model.SysRole)
	return data, nil
}

func (s *SysRoleService) GetRoleList(basic request.Basic) (response.ListResult[model.SysRole], error) {
	return s.sysRoleRepo.GetRoleList(basic)
}

func (s *SysRoleService) CreateRole(ctx *gin.Context, req request.RoleCreateReq) error {
	var role model.SysRole
	err := mapstructure.Decode(req, &role)
	if err != nil {
		fmt.Println("Error during struct mapping:", err)
		return err
	}
	id, err := s.sf.GenerateUniqueID()
	if err != nil {
		return err
	}
	role.Id = int(id)
	return s.sysRoleRepo.Create(s.sysRoleRepo.DBWithContext(ctx), &role)
}

func (s *SysRoleService) UpdateRole(ctx *gin.Context, req request.RoleUpdateReq) error {
	return s.sysRoleRepo.Update(s.sysRoleRepo.DBWithContext(ctx), &req)
}

func (s *SysRoleService) DeleteRole(ctx *gin.Context, id int) error {
	return s.sysRoleRepo.DeleteById(s.sysRoleRepo.DBWithContext(ctx), id)
}

func (s *SysRoleService) GetRoleMenus(roleId int) ([]model.SysMenu, error) {
	menus, err := s.sysRoleMenuRepo.GetRoleMenus(roleId)
	if err != nil {
		return nil, err
	}
	return utils.BuildMenuTree(menus, 0), nil
}

func (s *SysRoleService) GetRoleMenuButtons(roleId, menuId int) ([]model.SysMenuButton, error) {
	return s.sysRoleMenuButtonRepo.GetRoleMenuButtons(roleId, menuId)
}

// GetRoleMenuButtons 获取角色菜单按钮权限
func (s *SysMenuService) GetRoleMenuButtons(roleId, menuId int) ([]model.SysMenuButton, error) {
	return s.sysRoleMenuButtonRepo.GetRoleMenuButtons(roleId, menuId)
}

// CreateRoleMenu 新增角色菜单
func (s *SysRoleService) CreateRoleMenu(ctx *gin.Context, req request.RoleMenuCreateReq) error {
	var data model.SysRoleMenu
	err := mapstructure.Decode(req, &data)
	if err != nil {
		fmt.Println("Error during struct mapping:", err)
		return err
	}
	return s.sysRoleMenuRepo.Create(s.sysRoleMenuRepo.DBWithContext(ctx), &data)
}

// DeleteRoleMenu 删除角色菜单
func (s *SysRoleService) DeleteRoleMenu(ctx *gin.Context, roleId, menuId int) error {
	return s.sysRoleMenuRepo.DeleteRoleMenuByRoleIdAndMenuId(s.sysRoleMenuRepo.DBWithContext(ctx), roleId, menuId)
}
