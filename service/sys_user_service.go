/**
 * @Author: Nan
 * @Date: 2024/5/24 下午10:20
 */

package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
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

func (s *SysUserService) Insert(ctx *gin.Context, req request.UserCreateReq) error {
	var data model.SysUser
	//user, e := s.GetByUserName(req.UserName)
	user, e := s.GetByEmployeeId(req.EmployeeId)
	if e != nil {
		return e
	}
	if user.Id != 0 {
		e = &response.AdminError{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: "存在重复的用户",
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
	tx := s.sysUserRepo.DBWithContext(ctx)
	return s.sysUserRepo.Create(tx, data)
}

func (s *SysUserService) Update(ctx *gin.Context, req request.UserUpdateReq) error {
	tx := s.sysUserRepo.DBWithContext(ctx)
	err := s.sysUserRepo.Update(tx, req)
	if err != nil {
		return err
	}
	// 删除缓存
	s.DeleteCache(req.Id)
	return nil
}

func (s *SysUserService) Delete(ctx *gin.Context, id int) error {
	tx := s.sysUserRepo.DBWithContext(ctx)
	err := s.sysUserRepo.DeleteById(tx, id)
	if err != nil {
		return err
	}
	// 删除缓存
	s.DeleteCache(id)
	return nil
}

func (s *SysUserService) DeleteCache(userId int) {
	go func() {
		data, _ := s.GetById(userId)
		if data.Id != 0 {
			s.sysUserCache.Delete(strconv.Itoa(data.Id))
			s.sysUserCache.Delete(data.UserName)
			s.sysUserCache.Delete(data.PhoneNumber)
		}
	}()
}
