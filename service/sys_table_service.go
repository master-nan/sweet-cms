/**
 * @Author: Nan
 * @Date: 2024/5/17 上午11:30
 */

package service

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"sweet-cms/cache"
	"sweet-cms/form/request"
	"sweet-cms/model"
	"sweet-cms/repository"
	"sweet-cms/utils"
)

type SysTableService struct {
	sysTableRepo repository.SysTableRepository
	sf           *utils.Snowflake
	sysTableCode *cache.SysTableCache
}

func NewSysTableService(sysTableRepo repository.SysTableRepository, sf *utils.Snowflake, sysTableCode *cache.SysTableCache) *SysTableService {
	return &SysTableService{
		sysTableRepo,
		sf,
		sysTableCode,
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

func (s *SysTableService) GetTableFieldById(id int) (model.SysTableField, error) {
	return s.sysTableRepo.GetTableFieldById(id)
}

func (s *SysTableService) GetTableFieldsByTableId(tableId int) ([]model.SysTableField, error) {
	return s.sysTableRepo.GetTableFieldsByTableId(tableId)
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
