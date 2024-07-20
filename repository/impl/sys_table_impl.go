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

func (s *SysTableRepositoryImpl) InsertTable(tx *gorm.DB, table model.SysTable) error {
	return tx.Create(&table).Error
}

func (s *SysTableRepositoryImpl) UpdateTable(req request.TableUpdateReq) error {
	return s.db.Model(model.SysTable{Basic: model.Basic{Id: req.Id}}).Updates(&req).Error
}

// DeleteTableById 删除表信息数据
func (s *SysTableRepositoryImpl) DeleteTableById(tx *gorm.DB, i int) error {
	return tx.Where("id = ? ", i).Delete(model.SysTable{}).Error
}

func (s *SysTableRepositoryImpl) GetTableList(basic request.Basic) (response.ListResult[model.SysTable], error) {
	var repo response.ListResult[model.SysTable]
	var sysTableList []model.SysTable
	total, err := s.PaginateAndCountAsync(basic, &sysTableList)
	repo.Data = sysTableList
	repo.Total = int(total)
	return repo, err
}

func (s *SysTableRepositoryImpl) GetTableFieldById(i int) (model.SysTableField, error) {
	var tableField model.SysTableField
	err := s.db.Where("id = ? ", i).First(&tableField).Error
	return tableField, err
}

func (s *SysTableRepositoryImpl) GetTableFieldsByTableId(id int) ([]model.SysTableField, error) {
	var items []model.SysTableField
	err := s.db.Where("table_id = ?", id).Order("sequence").Find(&items).Error
	return items, err
}

func (s *SysTableRepositoryImpl) UpdateTableField(tx *gorm.DB, req request.TableFieldUpdateReq) error {
	return tx.Model(&model.SysTableField{}).Where("id = ?", req.Id).Updates(req).Error
}

func (s *SysTableRepositoryImpl) InsertTableField(tx *gorm.DB, field model.SysTableField) error {
	return tx.Create(&field).Error
}

func (s *SysTableRepositoryImpl) DeleteTableField(tx *gorm.DB, id int) error {
	return tx.Where("id = ?", id).Delete(model.SysTableField{}).Error
}

func (s *SysTableRepositoryImpl) DeleteTableFieldByTableId(tx *gorm.DB, tableId int) error {
	return tx.Where("table_id = ?", tableId).Delete(model.SysTableField{}).Error
}

func (s *SysTableRepositoryImpl) GetTableRelationById(i int) (model.SysTableRelation, error) {
	var relation model.SysTableRelation
	err := s.db.Where("id = ?", i).First(&relation).Error
	return relation, err
}

func (s *SysTableRepositoryImpl) GetTableRelationsByTableId(i int) ([]model.SysTableRelation, error) {
	var relations []model.SysTableRelation
	err := s.db.Where("table_id = ?", i).First(&relations).Error
	return relations, err
}

func (s *SysTableRepositoryImpl) InsertTableRelation(tx *gorm.DB, relation model.SysTableRelation) error {
	return tx.Create(&relation).Error
}

func (s *SysTableRepositoryImpl) DeleteTableRelation(tx *gorm.DB, id int) error {
	return tx.Where("id = ?", id).Delete(model.SysTableRelation{}).Error
}

func (s *SysTableRepositoryImpl) GetTableIndexesByTableId(id int) ([]model.SysTableIndex, error) {
	var indexes []model.SysTableIndex
	err := s.db.Where("table_id = ?", id).Find(&indexes).Error
	return indexes, err
}

func (s *SysTableRepositoryImpl) GetTableIndexById(id int) (model.SysTableIndex, error) {
	var index model.SysTableIndex
	err := s.db.Where("id = ?", id).Find(&index).Error
	return index, err
}

// InsertTableIndex 新增表索引
func (s *SysTableRepositoryImpl) InsertTableIndex(tx *gorm.DB, index model.SysTableIndex) error {
	// 创建索引表数据
	return tx.Create(&index).Error
}

// UpdateTableIndex 修改表索引
func (s *SysTableRepositoryImpl) UpdateTableIndex(tx *gorm.DB, req request.TableIndexUpdateReq) error {
	// 修改表数据
	return tx.Model(model.SysTableIndex{}).Where("id = ?", req.Id).Updates(&req).Error
}

func (s *SysTableRepositoryImpl) DeleteTableIndex(tx *gorm.DB, id int) error {
	return tx.Where("id = ?", id).Delete(model.SysTableIndex{}).Error
}

func (s *SysTableRepositoryImpl) DeleteTableIndexByTableId(tx *gorm.DB, tableId int) error {
	return tx.Where("table_id = ?", tableId).Delete(model.SysTableIndex{}).Error
}

func (s *SysTableRepositoryImpl) InsertTableIndexFields(tx *gorm.DB, indexFields []model.SysTableIndexField) error {
	return tx.Create(&indexFields).Error
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

// DeleteTableIndexFieldByIndexId 根据单个indexId删除中间表字段
func (s *SysTableRepositoryImpl) DeleteTableIndexFieldByIndexId(tx *gorm.DB, id int) error {
	return tx.Where("index_id = ?", id).Delete(model.SysTableIndexField{}).Error
}

// DeleteTableIndexFieldByIndexIds 根据所有indexId删除中间表字段
func (s *SysTableRepositoryImpl) DeleteTableIndexFieldByIndexIds(tx *gorm.DB, ids []int) error {
	return tx.Where("index_id in ?", ids).Delete(model.SysTableIndexField{}).Error
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
