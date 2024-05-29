/**
 * @Author: Nan
 * @Date: 2024/5/24 下午10:20
 */

package service

import (
	"sweet-cms/global"
	"sweet-cms/model"
)

type SysUserService struct {
}

func NewSysUserService() *SysUserService {
	return &SysUserService{}
}

// Get 根据username获取用户信息
func (s *SysUserService) Get(username string) (model.SysUser, error) {
	var user model.SysUser
	result := global.DB.Where(&model.SysUser{UserName: username}).First(&user)
	return user, result.Error
}
