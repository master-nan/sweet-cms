/**
 * @Author: Nan
 * @Date: 2024/7/25 下午3:29
 */

package service

import (
	"github.com/gin-gonic/gin"
	"sweet-cms/model"
	"sweet-cms/repository"
	"sweet-cms/utils"
)

type SysMenuService struct {
	sysMenuRepo         repository.SysMenuRepository
	sysRoleMenuRepo     repository.SysRoleMenuRepository
	sysUserMenuPermRepo repository.SysUserMenuDataPermissionRepository
	sysRoleMenuButtons  repository.SysRoleMenuButtonRepository
	sysUserRoleRepo     repository.SysUserRoleRepository
	sf                  *utils.Snowflake
}

func NewSysMenuService(sysMenuRepo repository.SysMenuRepository, sysRoleMenuRepo repository.SysRoleMenuRepository,
	sysUserMenuPermRepo repository.SysUserMenuDataPermissionRepository, sysRoleMenuButtons repository.SysRoleMenuButtonRepository,
	sysUserRoleRepo repository.SysUserRoleRepository, sf *utils.Snowflake) *SysMenuService {
	return &SysMenuService{
		sysMenuRepo,
		sysRoleMenuRepo,
		sysUserMenuPermRepo,
		sysRoleMenuButtons,
		sysUserRoleRepo,
		sf,
	}
}

// CreateMenu 新增菜单
func (s *SysMenuService) CreateMenu(ctx *gin.Context, menu model.SysMenu) error {
	return s.sysMenuRepo.Create(s.sysMenuRepo.DBWithContext(ctx), menu)
}

// UpdateMenu 更新菜单
func (s *SysMenuService) UpdateMenu(ctx *gin.Context, menu model.SysMenu) error {
	return s.sysMenuRepo.Update(s.sysMenuRepo.DBWithContext(ctx), menu)
}

// DeleteMenu 删除菜单
func (s *SysMenuService) DeleteMenu(ctx *gin.Context, id int) error {
	return s.sysMenuRepo.DeleteById(s.sysMenuRepo.DBWithContext(ctx), id)
}

// GetMenuTree 获取菜单列表并构建树结构
func (s *SysMenuService) GetMenuTree() ([]model.SysMenu, error) {
	menus, err := s.sysMenuRepo.GetMenus()
	if err != nil {
		return nil, err
	}
	return buildMenuTree(menus, 0), nil
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
	return buildMenuTree(menus, 0), nil
}

// GetRoleMenus 获取角色菜单权限
func (s *SysMenuService) GetRoleMenus(roleId int) ([]model.SysMenu, error) {
	menus, err := s.sysRoleMenuRepo.GetRoleMenus(roleId)
	if err != nil {
		return nil, err
	}
	return buildMenuTree(menus, 0), nil
}

// GetUserMenuPermissions 获取菜单的用户权限
func (s *SysMenuService) GetUserMenuPermissions(menuId int) ([]model.SysUserMenuDataPermission, error) {
	return s.sysUserMenuPermRepo.GetUserMenuPermissions(menuId)
}

// GetRoleMenuButtons 获取角色菜单按钮权限
func (s *SysMenuService) GetRoleMenuButtons(roleId, menuId int) ([]model.SysMenuButton, error) {
	return s.sysRoleMenuButtons.GetRoleMenuButtons(roleId, menuId)
}

// 递归构建树形结构
func buildMenuTree(menus []model.SysMenu, pid int) []model.SysMenu {
	var tree []model.SysMenu
	for _, menu := range menus {
		if menu.Pid == pid {
			menu.Children = buildMenuTree(menus, menu.Id)
			tree = append(tree, menu)
		}
	}
	return tree
}
