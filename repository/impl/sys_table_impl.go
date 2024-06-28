/**
 * @Author: Nan
 * @Date: 2024/6/10 上午12:16
 */

package impl

import (
	"fmt"
	"gorm.io/gorm"
	"reflect"
	"strings"
	"sweet-cms/enum"
	"sweet-cms/form/request"
	"sweet-cms/form/response"
	"sweet-cms/model"
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
	}).Preload("TableRelations").Preload("TableIndexes.IndexFields").Where("id = ", i).First(&table).Error
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

func (s *SysTableRepositoryImpl) GetTableList(basic request.Basic) (response.ListResult[model.SysTable], error) {
	var repo response.ListResult[model.SysTable]
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

func (s *SysTableRepositoryImpl) UpdateTableField(req request.TableFieldUpdateReq, field model.SysTableField, tableCode string) (err error) {
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
	var sqlType string
	if req.FieldType != field.FieldType || (req.FieldLength > 0 && req.FieldLength != field.FieldLength) {
		sqlType += fmt.Sprintf("%s(%d)", utils.SqlTypeFromFieldType(req.FieldType), req.FieldLength)
	}
	if req.DefaultValue != "" && req.DefaultValue != *field.DefaultValue {
		sqlType += fmt.Sprintf(" DEFAULT '%s'", req.DefaultValue)
	}
	if req.IsNull != field.IsNull {
		if req.IsNull {
			sqlType += " NULL"
		} else {
			sqlType += " NOT NULL"
		}
	}
	if req.FieldName != "" && req.FieldName != field.FieldName {
		sqlType += fmt.Sprintf(" COMMENT '%s'", req.FieldName)
	}
	var alterColumnSQL string
	if req.FieldCode == field.FieldCode {
		if sqlType == "" {
			return tx.Commit().Error
		}
		alterColumnSQL = fmt.Sprintf("ALTER TABLE `%s` MODIFY `%s` %s;", tableCode, req.FieldCode, sqlType)
	} else {
		alterColumnSQL = fmt.Sprintf("ALTER TABLE `%s` CHANGE `%s` `%s` %s;", tableCode, field.FieldCode, req.FieldCode, sqlType)
	}
	if err := tx.Exec(alterColumnSQL).Error; err != nil {
		tx.Rollback()
		return err
	}
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
	var relation model.SysTableRelation
	err := s.db.Where("id = ", i).First(&relation).Error
	return relation, err
}

func (s *SysTableRepositoryImpl) GetTableRelationByTableId(i int) (model.SysTableRelation, error) {
	var relation model.SysTableRelation
	err := s.db.Where("table_id = ", i).First(&relation).Error
	return relation, err
}

func (s *SysTableRepositoryImpl) InsertTableRelation(relation model.SysTableRelation, tableCode string) error {
	//TODO 新增多对多关系，2张表都要调整
	panic("implement me")
}

func (s *SysTableRepositoryImpl) UpdateTableRelation(req request.TableRelationUpdateReq, tableCode string) error {
	// TODO 判断是否修改成多对多关系，检查多对多关系表是否存在，修改多对多关心，2张表都要调整
	return s.db.Model(model.SysTableRelation{}).Updates(&req).Error
}

func (s *SysTableRepositoryImpl) DeleteTableRelation(relation model.SysTableRelation, tableCode string) error {
	//TODO 判断是否修改表关系，是否需要删除多对多关联表，同时检查多对多关系表是否存在
	return s.db.Where("id = ", relation.ID).Delete(model.SysTableRelation{}).Error
}

func (s *SysTableRepositoryImpl) GetTableIndexById(i int) (model.SysTableIndex, error) {
	var index model.SysTableIndex
	err := s.db.Where("id = ", i).First(&index).Error
	return index, err
}

func (s *SysTableRepositoryImpl) GetTableIndexByTableId(i int) (model.SysTableIndex, error) {
	var index model.SysTableIndex
	err := s.db.Where("table_id = ", i).First(&index).Error
	return index, err
}

func (s *SysTableRepositoryImpl) InsertTableIndex(index model.SysTableIndex, tableCode string) (err error) {
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
	// 创建索引表数据
	if err := tx.Create(&index).Error; err != nil {
		tx.Rollback()
		return err
	}

	var indexFields []model.SysTableIndexField
	fieldCodeList := make([]string, len(index.IndexFields))
	for _, field := range index.IndexFields {
		fieldCodeList = append(fieldCodeList, field.FieldCode)
		indexField := model.SysTableIndexField{
			IndexID: index.ID,
			FieldID: field.ID,
		}
		indexFields = append(indexFields, indexField)
	}
	// 创建中间表数据
	if err := tx.Create(&indexFields).Error; err != nil {
		tx.Rollback()
		return err
	}
	var unique string
	if index.IsUnique {
		unique = "UNIQUE"
	}
	fields := strings.Join(fieldCodeList, ",")
	createIndexSql := fmt.Sprintf("CREATE %s INDEX %s ON %s (%s)", unique, index.IndexName, tableCode, fields)
	if err := tx.Exec(createIndexSql).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (s *SysTableRepositoryImpl) UpdateTableIndex(req request.TableIndexUpdateReq, data model.SysTableIndex, tableCode string) error {
	tx := s.db.Begin()
	// 删除中间表字段
	if err := tx.Where("index_id = ?", req.ID).Delete(model.SysTableIndexField{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 修改表数据
	if err := tx.Model(model.SysTableIndex{}).Where("id=?", req.ID).Updates(&req).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 使用原索引名称删除表索引
	dropIndexSQL := fmt.Sprintf("DROP INDEX `%s` ON `%s`;", data.IndexName, tableCode)
	if err := tx.Exec(dropIndexSQL).Error; err != nil {
		tx.Rollback()
		return err
	}
	var indexFields []model.SysTableIndexField
	fieldCodeList := make([]string, len(req.IndexFields))
	for _, field := range req.IndexFields {
		fieldCodeList = append(fieldCodeList, field.FieldCode)
		indexField := model.SysTableIndexField{
			IndexID: req.ID,
			FieldID: field.FieldID,
		}
		indexFields = append(indexFields, indexField)
	}
	// 创建中间表数据
	if err := tx.Create(&indexFields).Error; err != nil {
		tx.Rollback()
		return err
	}
	var unique string
	if req.IsUnique {
		unique = "UNIQUE"
	}
	fields := strings.Join(fieldCodeList, ",")
	// 创建表索引
	createIndexSql := fmt.Sprintf("CREATE %s INDEX %s ON %s (%s)", unique, req.IndexName, tableCode, fields)
	if err := tx.Exec(createIndexSql).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (s *SysTableRepositoryImpl) DeleteTableIndex(index model.SysTableIndex, tableCode string) error {
	tx := s.db.Begin()
	// 删除字段
	if err := tx.Where("id = ?", index.ID).Delete(model.SysTableIndex{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 删除中间表字段
	if err := tx.Where("index_id = ?", index.ID).Delete(model.SysTableIndexField{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 构建删除索引的SQL语句
	dropIndexSQL := fmt.Sprintf("DROP INDEX %s ON %s", index.IndexName, tableCode)
	if err := tx.Exec(dropIndexSQL).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
