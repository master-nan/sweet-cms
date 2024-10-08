/**
 * @Author: Nan
 * @Date: 2024/7/25 下午3:29
 */

package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"sweet-cms/form/request"
	"sweet-cms/model"
	"sweet-cms/repository"
	"sweet-cms/utils"
)

type SysMenuService struct {
	sysMenuRepo           repository.SysMenuRepository
	sysRoleMenuRepo       repository.SysRoleMenuRepository
	sysUserMenuPermRepo   repository.SysUserMenuDataPermissionRepository
	sysRoleMenuButtonRepo repository.SysRoleMenuButtonRepository
	sysUserRoleRepo       repository.SysUserRoleRepository
	sysMenuButtonRepo     repository.SysMenuButtonRepository
	sf                    *utils.Snowflake
}

func NewSysMenuService(sysMenuRepo repository.SysMenuRepository, sysRoleMenuRepo repository.SysRoleMenuRepository,
	sysUserMenuPermRepo repository.SysUserMenuDataPermissionRepository, sysRoleMenuButtons repository.SysRoleMenuButtonRepository,
	sysUserRoleRepo repository.SysUserRoleRepository, sysMenuButtonRepo repository.SysMenuButtonRepository, sf *utils.Snowflake) *SysMenuService {
	return &SysMenuService{
		sysMenuRepo,
		sysRoleMenuRepo,
		sysUserMenuPermRepo,
		sysRoleMenuButtons,
		sysUserRoleRepo,
		sysMenuButtonRepo,
		sf,
	}
}

func (s *SysMenuService) GetMenuById(id int) (model.SysMenu, error) {
	result, err := s.sysMenuRepo.WithPreload("MenuButtons").FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.SysMenu{}, nil
		}
		return model.SysMenu{}, err
	}
	return result.(model.SysMenu), nil
}

// CreateMenu 新增菜单
func (s *SysMenuService) CreateMenu(ctx *gin.Context, req request.MenuCreateReq) error {
	var data model.SysMenu
	err := mapstructure.Decode(req, &data)
	if err != nil {
		fmt.Println("Error during struct mapping:", err)
		return err
	}
	id, err := s.sf.GenerateUniqueID()
	if err != nil {
		return err
	}
	data.Id = int(id)
	return s.sysMenuRepo.Create(s.sysMenuRepo.DBWithContext(ctx), &data)
}

// UpdateMenu 更新菜单
func (s *SysMenuService) UpdateMenu(ctx *gin.Context, data request.MenuUpdateReq) error {
	return s.sysMenuRepo.Update(s.sysMenuRepo.DBWithContext(ctx), &data, data.Id)
}

// DeleteMenuById 删除菜单
func (s *SysMenuService) DeleteMenuById(ctx *gin.Context, id int) error {
	return s.sysMenuRepo.DeleteById(s.sysMenuRepo.DBWithContext(ctx), id)
}

// GetMenuTree 获取菜单列表并构建树结构
func (s *SysMenuService) GetMenuTree() ([]model.SysMenu, error) {
	menus, err := s.sysMenuRepo.GetMenus()
	if err != nil {
		return nil, err
	}
	return utils.BuildMenuTree(menus, 0), nil
}

// GetUserMenus 获取用户菜单权限
func (s *SysMenuService) GetUserMenus(userId int) ([]model.SysMenu, error) {
	roles, err := s.sysUserRoleRepo.GetUserRoles(userId)
	if err != nil {
		return nil, err
	}
	var roleIds []int
	for _, role := range roles {
		roleIds = append(roleIds, role.Id)
	}
	menus, err := s.sysRoleMenuRepo.GetRoleMenusByRoleIds(roleIds)
	if err != nil {
		return nil, err
	}
	return utils.BuildMenuTree(menus, 0), nil
}

// GetUserMenuPermissions 获取菜单的用户权限
func (s *SysMenuService) GetUserMenuPermissions(menuId int) ([]model.SysUserMenuDataPermission, error) {
	return s.sysUserMenuPermRepo.GetUserMenuPermissions(menuId)
}
