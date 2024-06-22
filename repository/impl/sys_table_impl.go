/**
 * @Author: Nan
 * @Date: 2024/6/10 上午12:16
 */

package impl

import (
	"fmt"
	"gorm.io/gorm"
	"reflect"
	"sweet-cms/enum"
	"sweet-cms/form/request"
	"sweet-cms/model"
	"sweet-cms/repository"
	"sweet-cms/utils"
)

type SysTableRepositoryImpl struct {
	db *gorm.DB
	sf *utils.Snowflake
}

func NewSysTableRepositoryImpl(db *gorm.DB, sf *utils.Snowflake) *SysTableRepositoryImpl {
	return &SysTableRepositoryImpl{
		db,
		sf,
	}
}

func (s *SysTableRepositoryImpl) GetTableById(i int) (model.SysTable, error) {
	var table model.SysTable
	err := s.db.Preload("TableFields", func(db *gorm.DB) *gorm.DB {
		return db.Order("sequence")
	}).Where("id = ", i).First(&table).Error
	return table, err
}

func (s *SysTableRepositoryImpl) GetTableByTableCode(code string) (model.SysTable, error) {
	var table model.SysTable
	err := s.db.Preload("TableFields", func(db *gorm.DB) *gorm.DB {
		return db.Order("sequence")
	}).Where("table_code=?", code).First(&table).Error
	return table, err
}

func (s *SysTableRepositoryImpl) InsertTable(table model.SysTable) (err error) {
	// 开始事务
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r) // 打印错误信息
			tx.Rollback()                               // 回滚事务
			// 设置返回的错误信息
			if e, ok := r.(error); ok {
				err = e // 如果 r 是 error 类型，直接返回
			} else {
				// 如果 r 不是 error 类型，转换为 error 后返回
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	err = tx.Create(&table).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//// 构建创建实际表的SQL语句，包含Basic结构体的字段
	//createSQL := fmt.Sprintf("CREATE TABLE `%s` (", table.TableCode)
	//createSQL += `
	//   	id INT AUTO_INCREMENT PRIMARY KEY,
	//    gmt_create DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	//    gmt_create_user INT COMMENT '创建者',
	//    gmt_modify DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
	//    gmt_modify_user INT COMMENT '修改者',
	//    gmt_delete DATETIME COMMENT '删除时间',
	//    gmt_delete_user INT COMMENT '删除者',
	//    state BOOLEAN DEFAULT TRUE COMMENT '状态'
	//);`
	// 自动在sys_table_field中为Basic结构体中的每个字段创建记录
	basicFields := []model.SysTableField{
		{TableID: table.ID, FieldName: "ID", FieldCode: "id", FieldType: enum.INT, IsPrimaryKey: true, IsNull: false, InputType: enum.INPUT_NUMBER, IsSort: true, Sequence: 1, IsListShow: true},
		{TableID: table.ID, FieldName: "创建时间", FieldCode: "gmt_create", FieldType: enum.DATETIME, IsNull: false, InputType: enum.DATETIME_PICKER, IsSort: true, Sequence: 2, IsListShow: true},
		{TableID: table.ID, FieldName: "创建者", FieldCode: "gmt_create_user", FieldType: enum.INT, IsNull: false, InputType: enum.INPUT_NUMBER, Sequence: 3, IsListShow: true},
		{TableID: table.ID, FieldName: "修改时间", FieldCode: "gmt_modify", FieldType: enum.DATETIME, IsNull: false, InputType: enum.DATETIME_PICKER, IsSort: true, Sequence: 4, IsListShow: true},
		{TableID: table.ID, FieldName: "修改者", FieldCode: "gmt_modify_user", FieldType: enum.INT, IsNull: false, InputType: enum.INPUT_NUMBER, Sequence: 5, IsListShow: true},
		{TableID: table.ID, FieldName: "删除时间", FieldCode: "gmt_delete", FieldType: enum.DATETIME, IsNull: true, InputType: enum.DATETIME_PICKER, Sequence: 6},
		{TableID: table.ID, FieldName: "删除者", FieldCode: "gmt_delete_user", FieldType: enum.INT, IsNull: true, InputType: enum.INPUT_NUMBER, Sequence: 7},
		{TableID: table.ID, FieldName: "状态", FieldCode: "state", FieldType: enum.BOOLEAN, IsNull: false, InputType: enum.SELECT, IsSort: true, DefaultValue: utils.StringPtr("true"), DictCode: utils.StringPtr("whether"), Sequence: 8, IsListShow: true},
	}
	// 动态创建结构体类型
	dynamicType := utils.CreateDynamicStruct(basicFields)
	// 创建实例
	dynamicModel := reflect.New(dynamicType).Interface()
	tableName := utils.GetTableName(tx, table.TableCode)
	err = tx.Table(tableName).AutoMigrate(dynamicModel)
	if err != nil {
		tx.Rollback()
		return err
	}
	//// 执行创建表的SQL语句
	//if err := tx.Exec(createSQL).Error; err != nil {
	//	tx.Rollback()
	//	return err
	//}

	for i := range basicFields {
		fieldID, err := s.sf.GenerateUniqueID()
		if err != nil {
			tx.Rollback()
			return err
		}
		basicFields[i].ID = int(fieldID)
	}
	if err := tx.Create(&basicFields).Error; err != nil {
		tx.Rollback()
		return err
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
	query := utils.ExecuteQuery(s.db, basic)
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
	err := s.db.Where("table_id = ?", id).Order("sequence").Find(&items).Error
	return items, err
}

func (s *SysTableRepositoryImpl) UpdateTableField(req request.TableFieldUpdateReq, tableCode string) (err error) {
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r) // 打印错误信息
			tx.Rollback()                               // 回滚事务
			// 设置返回的错误信息
			if e, ok := r.(error); ok {
				err = e // 如果 r 是 error 类型，直接返回
			} else {
				// 如果 r 不是 error 类型，转换为 error 后返回
				err = fmt.Errorf("%v", r)
			}
		}
	}()
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
		sqlType += " NOT NULL"
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
	//indexName := fmt.Sprintf("idx_%s_%s", tableCode, req.FieldCode)
	//if req.IsIndex {
	//	// 检查索引是否存在
	//	var count int64
	//	tx.Raw("SHOW INDEX FROM `"+tableCode+"` WHERE Key_name = ?", indexName).Count(&count)
	//	if count == 0 {
	//		// 创建索引
	//		createIndexSQL := fmt.Sprintf("CREATE INDEX `%s` ON `%s`(`%s`);", indexName, tableCode, req.FieldCode)
	//		if err := tx.Exec(createIndexSQL).Error; err != nil {
	//			tx.Rollback()
	//			return err
	//		}
	//	}
	//} else {
	//	// 删除索引
	//	dropIndexSQL := fmt.Sprintf("DROP INDEX `%s` ON `%s`;", indexName, tableCode)
	//	if err := tx.Exec(dropIndexSQL).Error; err != nil {
	//		tx.Rollback()
	//		return err
	//	}
	//}
	return tx.Commit().Error
}

func (s *SysTableRepositoryImpl) InsertTableField(field model.SysTableField, tableCode string) (err error) {
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r) // 打印错误信息
			tx.Rollback()                               // 回滚事务
			// 设置返回的错误信息
			if e, ok := r.(error); ok {
				err = e // 如果 r 是 error 类型，直接返回
			} else {
				// 如果 r 不是 error 类型，转换为 error 后返回
				err = fmt.Errorf("%v", r)
			}
		}
	}()
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
	if field.IsNull {
		sqlType += " NULL"
	} else {
		sqlType += " NOT NULL"
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
	//if field.IsIndex {
	//	indexSQL := fmt.Sprintf("CREATE INDEX `idx_%s_%s` ON `%s`(`%s`);", tableCode, field.FieldCode, tableCode, field.FieldCode)
	//	if err := tx.Exec(indexSQL).Error; err != nil {
	//		tx.Rollback()
	//		return err
	//	}
	//}
	return tx.Commit().Error
}

func (s *SysTableRepositoryImpl) DeleteTableField(field model.SysTableField, tableCode string) (err error) {
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r) // 打印错误信息
			tx.Rollback()                               // 回滚事务
			// 设置返回的错误信息
			if e, ok := r.(error); ok {
				err = e // 如果 r 是 error 类型，直接返回
			} else {
				// 如果 r 不是 error 类型，转换为 error 后返回
				err = fmt.Errorf("%v", r)
			}
		}
	}()
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

func (s *SysTableRepositoryImpl) GetTableRelationById(i int) (model.SysTableRelation, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SysTableRepositoryImpl) GetTableRelationByTableId(i int) (model.SysTableRelation, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SysTableRepositoryImpl) UpdateTableRelation(req request.TableRelationUpdateReq, s2 string) error {
	//TODO implement me
	panic("implement me")
}

func (s *SysTableRepositoryImpl) InsertTableRelation(relation model.SysTableRelation, s2 string) error {
	//TODO implement me
	panic("implement me")
}

func (s *SysTableRepositoryImpl) DeleteTableRelation(relation model.SysTableRelation, s2 string) error {
	//TODO implement me
	panic("implement me")
}
