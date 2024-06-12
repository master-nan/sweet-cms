/**
 * @Author: Nan
 * @Date: 2024/5/23 下午2:59
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

type SysDictService struct {
	sysDictRepo  repository.SysDictRepository
	sf           *utils.Snowflake
	sysDictCache *cache.SysDictCache
}

func NewSysDictService(sysDictRepo repository.SysDictRepository, sf *utils.Snowflake, sysDictCache *cache.SysDictCache) *SysDictService {
	return &SysDictService{
		sysDictRepo,
		sf,
		sysDictCache,
	}
}

func (s *SysDictService) GetSysDictById(id int) (model.SysDict, error) {
	data, err := s.sysDictCache.Get(strconv.Itoa(id))
	if err != nil {
		d, e := s.sysDictRepo.GetSysDictById(id)
		if e != nil {
			if errors.Is(e, gorm.ErrRecordNotFound) {
				return d, nil
			}
			return d, e
		}
		if errors.As(err, inter.ErrCacheMiss) {
			s.sysDictCache.Set(strconv.Itoa(id), data)
		}
		return data, nil
	}
	return data, nil
}

func (s *SysDictService) GetSysDictList(basic request.Basic) (repository.SysDictListResult, error) {
	result, err := s.sysDictRepo.GetSysDictList(basic)
	return result, err
}

func (s *SysDictService) GetSysDictByCode(code string) (model.SysDict, error) {
	data, err := s.sysDictCache.Get(code)
	if err == nil {
		return data, nil
	}
	if !errors.Is(err, inter.ErrCacheMiss) { // 如果缓存错误不是因为未命中
		return model.SysDict{}, err
	}
	// 尝试从数据库获取
	dict, err := s.sysDictRepo.GetSysDictByCode(code)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.SysDict{}, nil
		}
		return model.SysDict{}, err
	}
	// 将结果设置回缓存
	s.sysDictCache.Set(code, dict)
	return dict, nil
}

func (s *SysDictService) InsertSysDict(req request.DictCreateReq) error {
	var data model.SysDict
	dict, e := s.GetSysDictByCode(req.DictCode)
	if e != nil {
		return e
	}
	if dict.ID != 0 {
		e = &response.AdminError{
			Code:    http.StatusBadRequest,
			Message: "存在重复的dict_code",
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
	data.ID = int(id)
	return s.sysDictRepo.InsertSysDict(data)
}

func (s *SysDictService) UpdateSysDict(req request.DictUpdateReq) error {
	err := s.sysDictRepo.UpdateSysDict(req)
	if err == nil {
		s.sysDictCache.Delete(strconv.Itoa(req.ID))
		s.sysDictCache.Delete(req.DictCode)
	}
	return err
}

func (s *SysDictService) DeleteSysDictById(id int) error {
	err := s.sysDictRepo.DeleteSysDictById(id)
	return err
}

func (s *SysDictService) GetSysDictItemById(id int) (model.SysDictItem, error) {
	data, err := s.sysDictRepo.GetSysDictItemById(id)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return data, nil
	}
	return data, err
}

func (s *SysDictService) GetSysDictItemsByDictId(id int) ([]model.SysDictItem, error) {
	result, err := s.sysDictRepo.GetSysDictItemsByDictId(id)
	return result, err
}

func (s *SysDictService) InsertSysDictItem(req request.DictItemCreateReq) error {
	var data model.SysDictItem
	err := mapstructure.Decode(req, &data)
	if err != nil {
		fmt.Println("Error during struct mapping:", err)
		return err
	}
	id, err := s.sf.GenerateUniqueID()
	if err != nil {
		return err
	}
	data.ID = int(id)
	err = s.sysDictRepo.InsertSysDictItem(data)
	if err == nil {
		s.sysDictCache.Delete(strconv.Itoa(req.DictID))
	}
	return err
}

func (s *SysDictService) UpdateSysDictItem(req request.DictItemUpdateReq) error {
	err := s.sysDictRepo.UpdateSysDictItem(req)
	if err == nil {
		dictItem, e := s.GetSysDictItemById(req.ID)
		if e == nil {
			s.sysDictCache.Delete(strconv.Itoa(*dictItem.DictID))
		}
	}
	return err
}

func (s *SysDictService) DeleteSysDictItemById(id int) error {
	err := s.sysDictRepo.DeleteSysDictItemById(id)
	if err == nil {
		dictItem, e := s.GetSysDictItemById(id)
		if e == nil {
			s.sysDictCache.Delete(strconv.Itoa(*dictItem.DictID))
		}
	}
	return err
}
