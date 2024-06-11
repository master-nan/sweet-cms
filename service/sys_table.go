/**
 * @Author: Nan
 * @Date: 2024/5/17 上午11:30
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

type SysTableService struct {
	sysTableRepo repository.SysTableRepository
	sf           *utils.Snowflake
}

func NewSysTableService(sysTableRepo repository.SysTableRepository, sf *utils.Snowflake) *SysTableService {
	return &SysTableService{
		sysTableRepo,
		sf,
	}
}

func (s *SysTableService) GetTableById(id int) (model.SysTable, error) {
	return s.sysTableRepo.GetTableById(id)
}

func (s *SysTableService) GetTableList(basic request.Basic) (repository.SysTableListResult, error) {
	result, err := s.sysTableRepo.GetTableList(basic)
	return result, err
}

func (s *SysTableService) GetTableByTableCode(code string) (model.SysTable, error) {
	return s.sysTableRepo.GetTableByTableCode(code)
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
	return s.sysTableRepo.UpdateTable(req)
}

func (s *SysTableService) DeleteTableById(id int) error {
	return s.sysTableRepo.DeleteTableById(id)
}
