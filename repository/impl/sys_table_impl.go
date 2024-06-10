/**
 * @Author: Nan
 * @Date: 2024/6/10 上午12:16
 */

package impl

import (
	"gorm.io/gorm"
	"sweet-cms/form/request"
	"sweet-cms/model"
	"sweet-cms/repository"
	"sweet-cms/utils"
)

type SysTableRepositoryImpl struct {
	db *gorm.DB
}

func NewSysTableRepositoryImpl(db *gorm.DB) *SysTableRepositoryImpl {
	return &SysTableRepositoryImpl{
		db,
	}
}

func (s *SysTableRepositoryImpl) GetTableById(i int) (model.SysTable, error) {
	var table model.SysTable
	err := s.db.Preload("TableFields").Where("id = ", i).First(&table).Error
	return table, err
}

func (s *SysTableRepositoryImpl) GetTableByTableCode(code string) (model.SysTable, error) {
	var table model.SysTable
	err := s.db.Preload("TableFields").Where("table_code=?", code).First(&table).Error
	return table, err
}

func (s *SysTableRepositoryImpl) InsertTable(table model.SysTable) error {
	return s.db.Create(&table).Error
}

func (s *SysTableRepositoryImpl) UpdateTable(req request.TableUpdateReq) error {
	return s.db.Model(model.SysTable{}).Updates(&req).Error
}

func (s *SysTableRepositoryImpl) DeleteTableById(i int) error {
	return s.db.Where("id = ", i).Delete(model.SysTable{}).Error
}

func (s *SysTableRepositoryImpl) GetTableList(basic request.Basic) (repository.SysTableListResult, error) {
	var repo repository.SysTableListResult
	query := utils.BuildQuery(s.db, basic)
	var sysTableList []model.SysTable
	var total int64 = 0
	err := query.Find(&sysTableList).Limit(-1).Offset(-1).Count(&total).Error
	repo.Data = sysTableList
	repo.Total = int(total)
	return repo, err
}

func (s *SysTableRepositoryImpl) GetTableFieldById(i int) (model.SysTableField, error) {
	var tableField model.SysTableField
	err := s.db.Where("id = ", i).First(&tableField).Error
	return tableField, err
}

func (s *SysTableRepositoryImpl) GetTableFieldsByTableId(id int) ([]model.SysTableField, error) {
	var items []model.SysTableField
	err := s.db.Where("table_id = ?", id).Find(&items).Error
	return items, err
}

func (s *SysTableRepositoryImpl) UpdateTableField(req request.TableFieldUpdateReq) error {
	return s.db.Model(model.SysTableField{}).Updates(&req).Error
}

func (s *SysTableRepositoryImpl) InsertTableField(field model.SysTableField) error {
	return s.db.Create(field).Error
}

func (s *SysTableRepositoryImpl) DeleteTableFieldById(i int) error {
	return s.db.Where("id = ", i).Delete(model.SysTableField{}).Error
}
