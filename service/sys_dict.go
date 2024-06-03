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

func (s *SysDictService) Get(id int) (model.SysDict, error) {
	return s.sysDictRepo.GetSysDictById(id)
}

func (s *SysDictService) Query(basic request.Basic) (repository.SysDictListResult, error) {
	result, err := s.sysDictRepo.GetSysDictList(basic)
	return result, err
}

func (s *SysDictService) Insert(req request.DictCreateReq) error {
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

func (s *SysDictService) Update(req request.DictUpdateReq) error {
	return s.sysDictRepo.UpdateSysDict(req)
}

func (s *SysDictService) Delete(id int) error {
	return s.sysDictRepo.DeleteSysDictById(id)
}

func (s *SysDictService) GetByCode(code string) (model.SysDict, error) {
	return s.sysDictRepo.GetSysDictByCode(code)
}
