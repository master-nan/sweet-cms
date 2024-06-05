/**
 * @Author: Nan
 * @Date: 2024/5/23 下午2:59
 */

package service

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"sweet-cms/form/request"
	"sweet-cms/model"
	"sweet-cms/repository"
	"sweet-cms/utils"
)

type SysDictService struct {
	sysDictRepo repository.SysDictRepository
	sf          *utils.Snowflake
}

func NewSysDictService(sysDictRepo repository.SysDictRepository, sf *utils.Snowflake) *SysDictService {
	return &SysDictService{
		sysDictRepo,
		sf,
	}
}

func (s *SysDictService) GetSysDictById(id int) (model.SysDict, error) {
	return s.sysDictRepo.GetSysDictById(id)
}

func (s *SysDictService) GetSysDictList(basic request.Basic) (repository.SysDictListResult, error) {
	result, err := s.sysDictRepo.GetSysDictList(basic)
	return result, err
}

func (s *SysDictService) InsertSysDict(req request.DictCreateReq) error {
	var data model.SysDict
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
	return s.sysDictRepo.UpdateSysDict(req)
}

func (s *SysDictService) DeleteSysDictById(id int) error {
	return s.sysDictRepo.DeleteSysDictById(id)
}

func (s *SysDictService) GetSysDictByCode(code string) (*model.SysDict, error) {
	return s.sysDictRepo.GetSysDictByCode(code)
}

func (s *SysDictService) GetSysDictItemById(id int) (model.SysDictItem, error) {
	return s.sysDictRepo.GetSysDictItemById(id)
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
	return s.sysDictRepo.InsertSysDictItem(data)
}

func (s *SysDictService) UpdateSysDictItem(req request.DictItemUpdateReq) error {
	return s.sysDictRepo.UpdateSysDictItem(req)
}

func (s *SysDictService) DeleteSysDictItemById(id int) error {
	return s.sysDictRepo.DeleteSysDictItemById(id)
}
