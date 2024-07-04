/**
 * @Author: Nan
 * @Date: 2024/5/24 下午10:20
 */

package service

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"sweet-cms/cache"
	"sweet-cms/form/request"
	"sweet-cms/form/response"
	"sweet-cms/inter"
	"sweet-cms/model"
	"sweet-cms/repository"
	"sweet-cms/utils"
)

type SysUserService struct {
	sysUserRepo  repository.SysUserRepository
	sf           *utils.Snowflake
	sysUserCache *cache.SysUserCache
}

func NewSysUserService(
	sysUserRepo repository.SysUserRepository,
	sf *utils.Snowflake,
	sysUserCache *cache.SysUserCache,
) *SysUserService {
	return &SysUserService{
		sysUserRepo,
		sf,
		sysUserCache,
	}
}

// GetByUserName 根据username获取用户信息
func (s *SysUserService) GetByUserName(username string) (model.SysUser, error) {
	data, err := s.sysUserCache.Get(username)
	if err == nil {
		return data, nil
	}
	if !errors.Is(err, inter.ErrCacheMiss) {
		return model.SysUser{}, err
	}
	data, err = s.sysUserRepo.GetByUserName(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.SysUser{}, nil
		}
		return model.SysUser{}, err
	}
	// 将用户按照id、username以及手机号缓存
	s.sysUserCache.Set(strconv.Itoa(data.Id), data)
	s.sysUserCache.Set(data.UserName, data)
	s.sysUserCache.Set(data.PhoneNumber, data)
	return data, nil
}

// GetById 根据id获取用户信息
func (s *SysUserService) GetById(id int) (model.SysUser, error) {
	data, err := s.sysUserCache.Get(strconv.Itoa(id))
	if err == nil {
		return data, nil
	}
	if !errors.Is(err, inter.ErrCacheMiss) {
		return model.SysUser{}, err
	}
	data, err = s.sysUserRepo.GetById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.SysUser{}, nil
		}
		return model.SysUser{}, err
	}
	// 将用户按照id、username以及手机号缓存
	s.sysUserCache.Set(string(data.Id), data)
	s.sysUserCache.Set(data.UserName, data)
	s.sysUserCache.Set(data.PhoneNumber, data)
	return data, nil
}

func (s *SysUserService) GetByEmployeeId(id int) (model.SysUser, error) {
	data, err := s.sysUserRepo.GetByEmployeeID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.SysUser{}, nil
		}
		return model.SysUser{}, err
	}
	return data, nil
}

func (s *SysUserService) GetList(basic request.Basic) (response.ListResult[model.SysUser], error) {
	result, err := s.sysUserRepo.GetList(basic)
	return result, err
}

func (s *SysUserService) Insert(req request.UserCreateReq) error {
	var data model.SysUser
	//user, e := s.GetByUserName(req.UserName)
	user, e := s.GetByEmployeeId(req.EmployeeId)
	if e != nil {
		return e
	}
	if user.Id != 0 {
		e = &response.AdminError{
			Code:    http.StatusBadRequest,
			Message: "存在重复的用户",
		}
		return e
	}
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
	return s.sysUserRepo.Insert(data)
}

func (s *SysUserService) Update(req request.UserUpdateReq) error {
	err := s.sysUserRepo.Update(req)
	if err != nil {
		return err
	}
	data, err := s.GetById(req.Id)
	if err != nil {
		return err
	}
	if data.Id != 0 {
		s.sysUserCache.Delete(strconv.Itoa(data.Id))
		s.sysUserCache.Delete(data.UserName)
		s.sysUserCache.Delete(data.PhoneNumber)
	}
	return nil
}

func (s *SysUserService) Delete(id int) error {
	data, err := s.GetById(id)
	if err != nil {
		return err
	}
	err = s.sysUserRepo.DeleteById(id)
	if err != nil {
		return err
	}
	if data.Id != 0 {
		s.sysUserCache.Delete(strconv.Itoa(data.Id))
		s.sysUserCache.Delete(data.UserName)
		s.sysUserCache.Delete(data.PhoneNumber)
	}
	return nil
}
