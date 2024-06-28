/**
 * @Author: Nan
 * @Date: 2024/5/24 下午10:20
 */

package service

import (
	"sweet-cms/cache"
	"sweet-cms/model"
	"sweet-cms/repository"
	"sweet-cms/utils"
)

type SysUserService struct {
	sysUserRepo  repository.SysUserRepository
	sf           *utils.Snowflake
	sysUserCache *cache.SysUserCache
}

func NewSysUserService(sysUserRepo repository.SysUserRepository, sf *utils.Snowflake, sysUserCache *cache.SysUserCache) *SysUserService {
	return &SysUserService{
		sysUserRepo,
		sf,
		sysUserCache,
	}
}

// GetByUserName 根据username获取用户信息
func (s *SysUserService) GetByUserName(username string) (model.SysUser, error) {
	user, err := s.sysUserRepo.GetByUserName(username)
	return user, err
}

// GetByUserId 根据id获取用户信息
func (s *SysUserService) GetByUserId(id int) (model.SysUser, error) {
	user, err := s.sysUserRepo.GetByUserId(id)
	return user, err
}
