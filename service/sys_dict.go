/**
 * @Author: Nan
 * @Date: 2024/5/23 下午2:59
 */

package service

import (
	"sweet-cms/model"
	"sweet-cms/repository"
)

type SysDictService struct {
	sysDictRepo repository.SysDictRepository
}

func NewSysDictServer(sysDictRepo repository.SysDictRepository) *SysDictService {
	return &SysDictService{
		sysDictRepo: sysDictRepo,
	}
}

func (s *SysDictService) Get(id int) (model.SysDict, error) {
	return s.sysDictRepo.GetSysDictById(id)
}

func (s *SysDictService) Query(name string) (model.SysDict, error) {
	return model.SysDict{}, nil
}
