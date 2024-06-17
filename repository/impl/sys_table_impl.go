/**
 * @Author: Nan
 * @Date: 2024/6/10 上午12:16
 */

package impl

import (
	"fmt"
	"gorm.io/gorm"
	"sweet-cms/enum"
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
	// 开始事务
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	err := s.db.Create(&table).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// 构建创建实际表的SQL语句，包含Basic结构体的字段
	createSQL := fmt.Sprintf("CREATE TABLE `%s` (", table.TableCode)
	createSQL += `
       	id INT AUTO_INCREMENT PRIMARY KEY,
        gmt_create DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        gmt_create_user INT COMMENT '创建者',
        gmt_modify DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
        gmt_modify_user INT COMMENT '修改者',
        gmt_delete DATETIME COMMENT '删除时间',
        gmt_delete_user INT COMMENT '删除者',
        state BOOLEAN DEFAULT TRUE COMMENT '状态'
    );`
	// 执行创建表的SQL语句
	if err := tx.Exec(createSQL).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 自动在sys_table_field中为Basic结构体中的每个字段创建记录
	basicFields := []model.SysTableField{
		{TableID: table.ID, FieldName: "ID", FieldCode: "id", FieldType: enum.INT, IsPrimaryKey: utils.BoolPtr(true), IsNull: utils.BoolPtr(false), InputType: enum.INPUT_NUMBER, IsSort: utils.BoolPtr(true)},
		{TableID: table.ID, FieldName: "创建时间", FieldCode: "gmt_create", FieldType: enum.DATETIME, IsNull: utils.BoolPtr(false), InputType: enum.DATETIME_PICKER, IsSort: utils.BoolPtr(true)},
		{TableID: table.ID, FieldName: "创建者", FieldCode: "gmt_create_user", FieldType: enum.INT, IsNull: utils.BoolPtr(true), InputType: enum.INPUT_NUMBER},
		{TableID: table.ID, FieldName: "修改时间", FieldCode: "gmt_modify", FieldType: enum.DATETIME, IsNull: utils.BoolPtr(false), InputType: enum.DATETIME_PICKER, IsSort: utils.BoolPtr(true)},
		{TableID: table.ID, FieldName: "修改者", FieldCode: "gmt_modify_user", FieldType: enum.INT, IsNull: utils.BoolPtr(true), InputType: enum.INPUT_NUMBER},
		{TableID: table.ID, FieldName: "删除时间", FieldCode: "gmt_delete", FieldType: enum.DATETIME, IsNull: utils.BoolPtr(true), InputType: enum.DATETIME_PICKER},
		{TableID: table.ID, FieldName: "删除者", FieldCode: "gmt_delete_user", FieldType: enum.INT, IsNull: utils.BoolPtr(true), InputType: enum.INPUT_NUMBER},
		{TableID: table.ID, FieldName: "状态", FieldCode: "state", FieldType: enum.BOOLEAN, IsNull: utils.BoolPtr(false), InputType: enum.SELECT, IsSort: utils.BoolPtr(true), DefaultValue: utils.StringPtr("true"), DictCode: utils.StringPtr("whether")},
	}
	for _, field := range basicFields {
		if err := tx.Create(&field).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	// 提交事务
	return tx.Commit().Error
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

func (s *SysTableRepositoryImpl) UpdateTableField(req request.TableFieldUpdateReq, tableCode string) error {
	tx := s.db.Begin()
	if err := tx.Model(&model.SysTableField{}).Where("id = ?", req.ID).Updates(req).Error; err != nil {
		tx.Rollback()
		return err
	}
	sqlType := utils.SqlTypeFromFieldType(req.FieldType)
	if req.FieldLength > 0 {
		sqlType += fmt.Sprintf("(%d)", req.FieldLength)
	}
	if req.DefaultValue != "" {
		sqlType += fmt.Sprintf(" DEFAULT '%s'", req.DefaultValue)
	}
	if req.IsNull {
		sqlType += " NULL"
	} else {
		sqlType += "NOT NULL"
	}
	if req.FieldName != "" {
		sqlType += fmt.Sprintf(" COMMENT '%s'", req.FieldName)
	}
	alterColumnSQL := fmt.Sprintf("ALTER TABLE `%s` MODIFY COLUMN `%s` %s;", tableCode, req.FieldCode, sqlType)
	if err := tx.Exec(alterColumnSQL).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 处理索引
	indexName := fmt.Sprintf("idx_%s_%s", tableCode, req.FieldCode)
	if req.IsIndex {
		// 检查索引是否存在
		var count int64
		tx.Raw("SHOW INDEX FROM `"+tableCode+"` WHERE Key_name = ?", indexName).Count(&count)
		if count == 0 {
			// 创建索引
			createIndexSQL := fmt.Sprintf("CREATE INDEX `%s` ON `%s`(`%s`);", indexName, tableCode, req.FieldCode)
			if err := tx.Exec(createIndexSQL).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	} else {
		// 删除索引
		dropIndexSQL := fmt.Sprintf("DROP INDEX `%s` ON `%s`;", indexName, tableCode)
		if err := tx.Exec(dropIndexSQL).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

func (s *SysTableRepositoryImpl) InsertTableField(field model.SysTableField, tableCode string) error {
	tx := s.db.Begin()
	// 创建字段记录
	if err := tx.Create(&field).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 构建SQL类型字符串，包括长度、默认值、是否可为空和备注
	sqlType := utils.SqlTypeFromFieldType(field.FieldType)
	if field.FieldLength > 0 {
		sqlType += fmt.Sprintf("(%d)", field.FieldLength)
	}
	if field.DefaultValue != nil {
		sqlType += fmt.Sprintf(" DEFAULT '%s'", field.DefaultValue)
	}
	if field.IsNull != nil && !*field.IsNull {
		sqlType += " NOT NULL"
	} else {
		sqlType += " NULL"
	}
	if field.FieldName != "" {
		sqlType += fmt.Sprintf(" COMMENT '%s'", field.FieldName)
	}
	// 添加字段
	addColumnSQL := fmt.Sprintf("ALTER TABLE `%s` ADD COLUMN `%s` %s;", tableCode, field.FieldCode, sqlType)
	if err := tx.Exec(addColumnSQL).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 如果字段需要索引
	if field.IsIndex {
		indexSQL := fmt.Sprintf("CREATE INDEX `idx_%s_%s` ON `%s`(`%s`);", tableCode, field.FieldCode, tableCode, field.FieldCode)
		if err := tx.Exec(indexSQL).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

func (s *SysTableRepositoryImpl) DeleteTableField(field model.SysTableField, tableCode string) error {
	tx := s.db.Begin()
	// 删除字段
	if err := tx.Where("id = ?", field.ID).Delete(model.SysTableField{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 构建删除字段的SQL语句
	dropColumnSQL := fmt.Sprintf("ALTER TABLE `%s` DROP COLUMN `%s`;", tableCode, field.FieldCode)
	if err := tx.Exec(dropColumnSQL).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
