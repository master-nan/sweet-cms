/**
 * @Author: Nan
 * @Date: 2024/6/3 下午6:08
 */

package impl

import (
	"gorm.io/gorm"
	"sweet-cms/form/request"
	"sweet-cms/form/response"
	"sweet-cms/model"
	"sweet-cms/repository/util"
)

type SysUserRepositoryImpl struct {
	db *gorm.DB
	*BasicImpl
}

func NewSysUserRepositoryImpl(db *gorm.DB, basicImpl *BasicImpl) *SysUserRepositoryImpl {
	return &SysUserRepositoryImpl{db, basicImpl}
}

func (s *SysUserRepositoryImpl) GetByUserName(username string) (model.SysUser, error) {
	var user model.SysUser
	result := s.db.Where(&model.SysUser{UserName: username}).Or(&model.SysUser{PhoneNumber: username}).First(&user)
	return user, result.Error
}

func (s *SysUserRepositoryImpl) GetById(id int) (model.SysUser, error) {
	var user model.SysUser
	result := s.db.Where("id = ?", id).First(&user)
	return user, result.Error
}

func (s *SysUserRepositoryImpl) GetByEmployeeID(id int) (model.SysUser, error) {
	var user model.SysUser
	result := s.db.Where("employee_id = ?", id).First(&user)
	return user, result.Error
}

func (s *SysUserRepositoryImpl) Insert(tx *gorm.DB, d model.SysUser) error {
	return tx.Create(&d).Error
}

func (s *SysUserRepositoryImpl) Update(tx *gorm.DB, req request.UserUpdateReq) error {
	return tx.Model(model.SysUser{}).Where("id=?", req.Id).Updates(&req).Error
}

func (s *SysUserRepositoryImpl) DeleteById(tx *gorm.DB, i int) error {
	return tx.Where("id = ", i).Delete(model.SysUser{}).Error
}

func (s *SysUserRepositoryImpl) GetList(basic request.Basic) (response.ListResult[model.SysUser], error) {
	var repo response.ListResult[model.SysUser]
	query := util.ExecuteQuery(s.db, basic)
	var sysUserList []model.SysUser
	var total int64 = 0
	err := query.Find(&sysUserList).Limit(-1).Offset(-1).Count(&total).Error
	repo.Data = sysUserList
	repo.Total = int(total)
	return repo, err
}

// GetUserMenuPermissions 获取用户的菜单权限
func (s *SysUserRepositoryImpl) GetUserMenuPermissions(userId int) ([]model.SysUserMenuDataPermission, error) {
	var permissions []model.SysUserMenuDataPermission
	err := s.db.Where("user_id = ?", userId).Find(&permissions).Error
	return permissions, err
}

// GetUserMenus 获取用户的菜单
func (s *SysUserRepositoryImpl) GetUserMenus(userId int) ([]model.SysMenu, error) {
	var user model.SysUser
	err := s.db.Preload("Permissions").Preload("Permissions.Menu").First(&user, userId).Error
	if err != nil {
		return nil, err
	}
	var menus []model.SysMenu
	for _, permission := range user.Permissions {
		menus = append(menus, permission.Menu)
	}
	return menus, nil
}
