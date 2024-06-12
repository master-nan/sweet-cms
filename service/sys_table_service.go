/**
 * @Author: Nan
 * @Date: 2024/5/17 上午11:30
 */

package service

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strconv"
	"sweet-cms/cache"
	"sweet-cms/form/request"
	"sweet-cms/inter"
	"sweet-cms/model"
	"sweet-cms/repository"
	"sweet-cms/utils"
)

type SysTableService struct {
	sysTableRepo       repository.SysTableRepository
	sf                 *utils.Snowflake
	sysTableCache      *cache.SysTableCache
	sysTableFieldCache *cache.SysTableFieldCache
}

func NewSysTableService(
	sysTableRepo repository.SysTableRepository,
	sf *utils.Snowflake,
	sysTableCache *cache.SysTableCache,
	sysTableFieldCache *cache.SysTableFieldCache,
) *SysTableService {
	return &SysTableService{
		sysTableRepo,
		sf,
		sysTableCache,
		sysTableFieldCache,
	}
}

func (s *SysTableService) GetTableById(id int) (model.SysTable, error) {
	data, err := s.sysTableCache.Get(strconv.Itoa(id))
	if err == nil {
		return data, nil
	}
	if !errors.Is(err, inter.ErrCacheMiss) {
		return model.SysTable{}, err
	}
	data, err = s.sysTableRepo.GetTableById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.SysTable{}, nil
		}
		return model.SysTable{}, err
	}
	s.sysTableCache.Set(string(id), data)
	return data, nil
}

func (s *SysTableService) GetTableList(basic request.Basic) (repository.SysTableListResult, error) {
	result, err := s.sysTableRepo.GetTableList(basic)
	return result, err
}

func (s *SysTableService) GetTableByTableCode(code string) (model.SysTable, error) {
	data, err := s.sysTableCache.Get(code)
	if err == nil {
		return data, nil
	}
	if !errors.Is(err, inter.ErrCacheMiss) {
		return model.SysTable{}, err
	}
	data, err = s.sysTableRepo.GetTableByTableCode(code)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.SysTable{}, nil
		}
		return model.SysTable{}, err
	}
	s.sysTableCache.Set(code, data)
	return data, nil
}

func (s *SysTableService) InsertTable(req request.TableCreateReq) error {
	var data model.SysTable
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
	return s.sysTableRepo.InsertTable(data)
}

func (s *SysTableService) UpdateTable(req request.TableUpdateReq) error {
	err := s.sysTableRepo.UpdateTable(req)
	if err != nil {
		return err
	}
	data, err := s.GetTableById(req.ID)
	if err != nil {
		return err
	}
	if data.ID != 0 {
		s.sysTableCache.Delete(strconv.Itoa(data.ID))
		s.sysTableCache.Delete(data.TableCode)
	}
	return nil
}

func (s *SysTableService) DeleteTableById(id int) error {
	err := s.sysTableRepo.DeleteTableById(id)
	if err != nil {
		return err
	}
	data, err := s.GetTableById(id)
	if err != nil {
		return err
	}
	if data.ID != 0 {
		s.sysTableCache.Delete(strconv.Itoa(data.ID))
		s.sysTableCache.Delete(data.TableCode)
	}
	return nil
}

func (s *SysTableService) GetTableFieldById(id int) (model.SysTableField, error) {
	data, err := s.sysTableFieldCache.Get(strconv.Itoa(id))
	if err == nil {
		return data, nil
	}
	if !errors.Is(err, inter.ErrCacheMiss) {
		return model.SysTableField{}, err
	}
	data, err = s.sysTableRepo.GetTableFieldById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.SysTableField{}, nil
		}
		return model.SysTableField{}, err
	}
	s.sysTableFieldCache.Set(strconv.Itoa(id), data)
	return data, nil
}

func (s *SysTableService) GetTableFieldsByTableId(tableId int) ([]model.SysTableField, error) {
	data, err := s.sysTableCache.Get(strconv.Itoa(tableId))
	if err == nil {
		return data.TableFields, nil
	}
	if !errors.Is(err, inter.ErrCacheMiss) {
		return nil, err
	}
	fields, err := s.sysTableRepo.GetTableFieldsByTableId(tableId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.SysTableField{}, nil
		}
		return nil, err
	}
	return fields, nil
}

func (s *SysTableService) UpdateTableField(req request.TableFieldUpdateReq) error {
	return s.sysTableRepo.UpdateTableField(req)
}

func (s *SysTableService) DeleteTableFieldById(id int) error {
	return s.sysTableRepo.DeleteTableFieldById(id)
}

func (s *SysTableService) InsertTableField(data model.SysTableField) error {
	return s.sysTableRepo.InsertTableField(data)
}
