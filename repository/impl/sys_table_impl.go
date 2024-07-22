/**
 * @Author: Nan
 * @Date: 2024/6/10 上午12:16
 */

package impl

import (
	"fmt"
	"gorm.io/gorm"
	"reflect"
	"sweet-cms/form/request"
	"sweet-cms/form/response"
	"sweet-cms/model"
	"sweet-cms/repository/util"
)

type SysTableRepositoryImpl struct {
	db *gorm.DB
	*BasicImpl
}

func NewSysTableRepositoryImpl(db *gorm.DB) *SysTableRepositoryImpl {
	return &SysTableRepositoryImpl{
		db,
		NewBasicImpl(db, &model.SysTable{}),
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
	}).Where("table_code = ? ", code).First(&table).Error
	return table, err
}

func (s *SysTableRepositoryImpl) GetTableList(basic request.Basic) (response.ListResult[model.SysTable], error) {
	var repo response.ListResult[model.SysTable]
	var sysTableList []model.SysTable
	total, err := s.PaginateAndCountAsync(basic, &sysTableList)
	repo.Data = sysTableList
	repo.Total = int(total)
	return repo, err
}

// CreateTableIndex 删除实体表索引
func (s *SysTableRepositoryImpl) CreateTableIndex(tx *gorm.DB, isUnique bool, indexName string, tableCode string, fields string) error {
	var unique string
	if isUnique {
		unique = "UNIQUE"
	}
	tableName := util.GetTableName(tx, tableCode)
	createIndexSql := fmt.Sprintf("CREATE %s INDEX %s ON %s (%s)", unique, indexName, tableName, fields)
	return tx.Exec(createIndexSql).Error
}

// DropTableIndex 删除实体表索引
func (s *SysTableRepositoryImpl) DropTableIndex(tx *gorm.DB, indexName string, tableCode string) error {
	tableName := util.GetTableName(tx, tableCode)
	// 构建删除索引的SQL语句
	dropIndexSQL := fmt.Sprintf("DROP INDEX %s ON %s", indexName, tableName)
	return tx.Exec(dropIndexSQL).Error
}

func (s *SysTableRepositoryImpl) CreateTable(tx *gorm.DB, tableCode string, model any) error {
	tableName := util.GetTableName(tx, tableCode)
	// 检查表是否存在
	if !tx.Migrator().HasTable(tableName) {
		// 不存在则创建表
		err := tx.Table(tableName).AutoMigrate(model)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *SysTableRepositoryImpl) DropTable(tx *gorm.DB, tableCode string) error {
	tableName := util.GetTableName(tx, tableCode)
	// 检查表是否存在
	if tx.Migrator().HasTable(tableName) {
		// 删除表
		err := tx.Migrator().DropTable(tableName)
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateTableColumn 添加实体字段
func (s *SysTableRepositoryImpl) CreateTableColumn(tx *gorm.DB, tableCode string, fieldCode string, sqlType string) error {
	tableName := util.GetTableName(tx, tableCode)
	addColumnSQL := fmt.Sprintf("ALTER TABLE `%s` ADD COLUMN `%s` %s;", tableName, fieldCode, sqlType)
	return tx.Exec(addColumnSQL).Error
}

// DropTableColumn 删除实体字段
func (s *SysTableRepositoryImpl) DropTableColumn(tx *gorm.DB, tableCode string, fieldCode string) error {
	tableName := util.GetTableName(tx, tableCode)
	// 构建删除字段的SQL语句
	dropColumnSQL := fmt.Sprintf("ALTER TABLE `%s` DROP COLUMN `%s`;", tableName, fieldCode)
	return tx.Exec(dropColumnSQL).Error
}

func (s *SysTableRepositoryImpl) ModifyTableColumn(tx *gorm.DB, tableCode string, fieldCode string, sqlType string) error {
	tableName := util.GetTableName(tx, tableCode)
	alterColumnSQL := fmt.Sprintf("ALTER TABLE `%s` MODIFY `%s` %s;", tableName, fieldCode, sqlType)
	return tx.Exec(alterColumnSQL).Error
}

func (s *SysTableRepositoryImpl) ChangeTableColumn(tx *gorm.DB, tableCode string, originalFieldCode string, fieldCode string, sqlType string) error {
	tableName := util.GetTableName(tx, tableCode)
	alterColumnSQL := fmt.Sprintf("ALTER TABLE `%s` CHANGE `%s` `%s` %s;", tableName, originalFieldCode, fieldCode, sqlType)
	return tx.Exec(alterColumnSQL).Error
}

func (s *SysTableRepositoryImpl) FetchTableMetadata(tableSchema string, tableCode string) ([]model.TableColumnMate, error) {
	var columns []model.TableColumnMate
	tableName := util.GetTableName(s.db, tableCode)
	query := `SELECT *  FROM INFORMATION_SCHEMA.COLUMNS  WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?;`
	err := s.db.Raw(query, tableSchema, tableName).Scan(&columns).Error
	if err != nil {
		return []model.TableColumnMate{}, err
	}
	return columns, nil
}

func (s *SysTableRepositoryImpl) FetchTableIndexMetadata(tableSchema string, tableCode string) ([]model.TableIndexMate, error) {
	var indexes []model.TableIndexMate
	tableName := util.GetTableName(s.db, tableCode)
	query := `SELECT COLUMN_NAME, INDEX_NAME, NON_UNIQUE, INDEX_TYPE FROM
    INFORMATION_SCHEMA.STATISTICS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ? AND INDEX_NAME != 'PRIMARY';`
	err := s.db.Raw(query, tableSchema, tableName).Scan(&indexes).Error
	if err != nil {
		return []model.TableIndexMate{}, err
	}
	return indexes, nil
}

func (s *SysTableRepositoryImpl) Model(data []model.SysTableField) interface{} {
	// 动态创建结构体类型
	dynamicType := util.CreateDynamicStruct(data)
	// 创建实例
	dynamicModel := reflect.New(dynamicType).Interface()
	return dynamicModel
}
