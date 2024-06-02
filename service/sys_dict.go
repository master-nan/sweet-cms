/**
 * @Author: Nan
 * @Date: 2024/5/23 下午2:59
 */

package service

import (
	"sweet-cms/form/request"
	"sweet-cms/model"
	"sweet-cms/repository"
)

type SysDictService struct {
	sysDictRepo repository.SysDictRepository
}

func NewSysDictService(sysDictRepo repository.SysDictRepository) *SysDictService {
	return &SysDictService{
		sysDictRepo: sysDictRepo,
	}
}

func (s *SysDictService) Get(id int) (model.SysDict, error) {
	return s.sysDictRepo.GetSysDictById(id)
}

func (s *SysDictService) Query(basic request.Basic) (repository.SysDictListResult, error) {
	result, err := s.sysDictRepo.GetSysDictList(basic)
	return result, err
}

func (s *SysDictService) Insert(d *model.SysDict) error {
	err := s.sysDictRepo.InsertSysDict(d)
	return err
}

func (s *SysDictService) Delete(id int) error {
	err := s.sysDictRepo.DeleteSysDictById(id)
	return err
}
