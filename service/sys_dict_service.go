/**
 * @Author: Nan
 * @Date: 2024/5/23 下午2:59
 */

package service

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"go.uber.org/zap"
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
	if err == nil {
		return data, nil
	}
	if !errors.Is(err, inter.ErrCacheMiss) {
		return model.SysDict{}, err
	}
	// 尝试从数据库获取
	dict, err := s.sysDictRepo.GetSysDictById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.SysDict{}, nil
		}
		return model.SysDict{}, err
	}
	s.sysDictCache.Set(strconv.Itoa(id), dict)
	return dict, nil
}

func (s *SysDictService) GetSysDictList(basic request.Basic) (response.ListResult[model.SysDict], error) {
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
	if err != nil {
		return err
	}
	data, err := s.GetSysDictById(req.ID)
	if err != nil {
		return err
	}
	if data.ID != 0 {
		s.sysDictCache.Delete(strconv.Itoa(data.ID))
		s.sysDictCache.Delete(data.DictCode)
	}
	return nil
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
		zap.L().Error("Error during struct mapping:", zap.Error(err))
		return err
	}
	id, err := s.sf.GenerateUniqueID()
	if err != nil {
		return err
	}
	data.ID = int(id)
	err = s.sysDictRepo.InsertSysDictItem(data)
	if err != nil {
		zap.L().Error("InsertSysDictItem err:", zap.Error(err))
		return err
	}
	dict, err := s.GetSysDictById(req.DictID)
	if err != nil {
		zap.L().Error("InsertSysDictItem err:", zap.Error(err))
		return err
	}
	if dict.ID != 0 {
		s.sysDictCache.Delete(strconv.Itoa(dict.ID))
		s.sysDictCache.Delete(dict.DictCode)
	}
	return nil
}

func (s *SysDictService) UpdateSysDictItem(req request.DictItemUpdateReq) error {
	err := s.sysDictRepo.UpdateSysDictItem(req)
	if err != nil {
		zap.L().Error("UpdateSysDictItem err:", zap.Error(err))
		return err
	}
	dictItem, err := s.GetSysDictItemById(req.ID)
	if err != nil {
		zap.L().Error("UpdateSysDictItem err:", zap.Error(err))
		return err
	}
	dict, err := s.GetSysDictById(dictItem.DictID)
	if err != nil {
		zap.L().Error("UpdateSysDictItem err:", zap.Error(err))
		return err
	}
	if dict.ID != 0 {
		s.sysDictCache.Delete(strconv.Itoa(dict.ID))
		s.sysDictCache.Delete(dict.DictCode)
	}
	return nil
}

func (s *SysDictService) DeleteSysDictItemById(id int) error {
	err := s.sysDictRepo.DeleteSysDictItemById(id)
	if err != nil {
		zap.L().Error("DeleteSysDictItemById err:", zap.Error(err))
		return err
	}
	dictItem, err := s.GetSysDictItemById(id)
	if err != nil {
		zap.L().Error("DeleteSysDictItemById err:", zap.Error(err))
		return err
	}
	dict, err := s.GetSysDictById(dictItem.DictID)
	if err != nil {
		zap.L().Error("DeleteSysDictItemById err:", zap.Error(err))
		return err
	}
	if dict.ID != 0 {
		s.sysDictCache.Delete(strconv.Itoa(dict.ID))
		s.sysDictCache.Delete(dict.DictCode)
	}
	return nil
}
