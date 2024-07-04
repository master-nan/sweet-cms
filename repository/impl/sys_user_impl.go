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
	"sweet-cms/utils"
)

type SysUserRepositoryImpl struct {
	db *gorm.DB
}

func NewSysUserRepositoryImpl(db *gorm.DB) *SysUserRepositoryImpl {
	return &SysUserRepositoryImpl{db}
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

func (s *SysUserRepositoryImpl) Insert(d model.SysUser) error {
	result := s.db.Create(&d)
	return result.Error
}

func (s *SysUserRepositoryImpl) Update(req request.UserUpdateReq) error {
	return s.db.Model(model.SysUser{}).Where("id=?", req.ID).Updates(&req).Error
}

func (s *SysUserRepositoryImpl) DeleteById(i int) error {
	return s.db.Where("id = ", i).Delete(model.SysUser{}).Error
}

func (s *SysUserRepositoryImpl) GetList(basic request.Basic) (response.ListResult[model.SysUser], error) {
	var repo response.ListResult[model.SysUser]
	query := utils.ExecuteQuery(s.db, basic)
	var sysUserList []model.SysUser
	var total int64 = 0
	err := query.Find(&sysUserList).Limit(-1).Offset(-1).Count(&total).Error
	repo.Data = sysUserList
	repo.Total = int(total)
	return repo, err
}
