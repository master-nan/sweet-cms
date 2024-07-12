/**
 * @Author: Nan
 * @Date: 2024/6/10 上午12:16
 */

package impl

import (
	"fmt"
	"gorm.io/gorm"
	"sweet-cms/form/request"
	"sweet-cms/form/response"
	"sweet-cms/model"
	"sweet-cms/utils"
)

type SysTableRepositoryImpl struct {
	db *gorm.DB
	*BasicImpl
}

func NewSysTableRepositoryImpl(db *gorm.DB, basicImpl *BasicImpl) *SysTableRepositoryImpl {
	return &SysTableRepositoryImpl{
		db,
		basicImpl,
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

func (s *SysTableRepositoryImpl) InsertTable(tx *gorm.DB, table model.SysTable) (err error) {
	if tx == nil {
		tx = s.db
	}
	err = tx.Create(&table).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *SysTableRepositoryImpl) UpdateTable(req request.TableUpdateReq) error {
	return s.db.Model(model.SysTable{Basic: model.Basic{Id: req.Id}}).Updates(&req).Error
}

// DeleteTableById 删除表信息数据
func (s *SysTableRepositoryImpl) DeleteTableById(tx *gorm.DB, i int) (err error) {
	if tx == nil {
		tx = s.db
	}
	if err = tx.Where("id = ? ", i).Delete(model.SysTable{}).Error; err != nil {
		return err
	}
	return nil
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
	err := s.db.Where("id = ? ", i).First(&tableField).Error
	return tableField, err
}

func (s *SysTableRepositoryImpl) GetTableFieldsByTableId(id int) ([]model.SysTableField, error) {
	var items []model.SysTableField
	err := s.db.Where("table_id = ?", id).Order("sequence").Find(&items).Error
	return items, err
}

func (s *SysTableRepositoryImpl) UpdateTableField(tx *gorm.DB, req request.TableFieldUpdateReq) error {
	if tx == nil {
		tx = s.db
	}
	if err := tx.Model(&model.SysTableField{}).Where("id = ?", req.Id).Updates(req).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysTableRepositoryImpl) ModifyTableColumn(tx *gorm.DB, tableCode string, fieldCode string, sqlType string) error {
	if tx == nil {
		tx = s.db
	}
	tableName := utils.GetTableName(tx, tableCode)
	alterColumnSQL := fmt.Sprintf("ALTER TABLE `%s` MODIFY `%s` %s;", tableName, fieldCode, sqlType)
	if err := tx.Exec(alterColumnSQL).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysTableRepositoryImpl) ChangeTableColumn(tx *gorm.DB, tableCode string, originalFieldCode string, fieldCode string, sqlType string) error {
	if tx == nil {
		tx = s.db
	}
	tableName := utils.GetTableName(tx, tableCode)
	alterColumnSQL := fmt.Sprintf("ALTER TABLE `%s` CHANGE `%s` `%s` %s;", tableName, originalFieldCode, fieldCode, sqlType)
	if err := tx.Exec(alterColumnSQL).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysTableRepositoryImpl) InsertTableField(tx *gorm.DB, field model.SysTableField) error {
	if tx == nil {
		tx = s.db
	}
	if err := tx.Create(&field).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysTableRepositoryImpl) DeleteTableField(tx *gorm.DB, id int) error {
	if tx == nil {
		tx = s.db
	}
	if err := tx.Where("id = ?", id).Delete(model.SysTableField{}).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysTableRepositoryImpl) DeleteTableFieldByTableId(tx *gorm.DB, tableId int) error {
	if tx == nil {
		tx = s.db
	}
	if err := tx.Where("table_id = ?", tableId).Delete(model.SysTableField{}).Error; err != nil {
		return err
	}
	return nil
}

// CreateTableColumn 添加实体字段
func (s *SysTableRepositoryImpl) CreateTableColumn(tx *gorm.DB, tableCode string, fieldCode string, sqlType string) error {
	if tx == nil {
		tx = s.db
	}
	tableName := utils.GetTableName(tx, tableCode)
	addColumnSQL := fmt.Sprintf("ALTER TABLE `%s` ADD COLUMN `%s` %s;", tableName, fieldCode, sqlType)
	if err := tx.Exec(addColumnSQL).Error; err != nil {
		return err
	}
	return nil
}

// DropTableColumn 删除实体字段
func (s *SysTableRepositoryImpl) DropTableColumn(tx *gorm.DB, tableCode string, fieldCode string) error {
	if tx == nil {
		tx = s.db
	}
	tableName := utils.GetTableName(tx, tableCode)
	// 构建删除字段的SQL语句
	dropColumnSQL := fmt.Sprintf("ALTER TABLE `%s` DROP COLUMN `%s`;", tableName, fieldCode)
	if err := tx.Exec(dropColumnSQL).Error; err != nil {
		return err
	}
	return nil
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
	if tx == nil {
		tx = s.db
	}
	if err := tx.Create(&relation).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysTableRepositoryImpl) DeleteTableRelation(tx *gorm.DB, id int) error {
	if tx == nil {
		tx = s.db
	}
	if err := tx.Where("id = ?", id).Delete(model.SysTableRelation{}).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysTableRepositoryImpl) CreateTable(tx *gorm.DB, tableCode string, model any) error {
	if tx == nil {
		tx = s.db
	}
	tableName := utils.GetTableName(tx, tableCode)
	return tx.Table(tableName).AutoMigrate(model)
}

func (s *SysTableRepositoryImpl) DropTable(tx *gorm.DB, tableCode string) error {
	if tx == nil {
		tx = s.db
	}
	tableName := utils.GetTableName(tx, tableCode)
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
	if tx == nil {
		tx = s.db
	}
	// 创建索引表数据
	if err := tx.Create(&index).Error; err != nil {
		return err
	}
	return nil
}

// UpdateTableIndex 修改表索引
func (s *SysTableRepositoryImpl) UpdateTableIndex(tx *gorm.DB, req request.TableIndexUpdateReq) error {
	if tx == nil {
		tx = s.db
	}
	// 修改表数据
	if err := tx.Model(model.SysTableIndex{}).Where("id = ?", req.Id).Updates(&req).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysTableRepositoryImpl) DeleteTableIndex(tx *gorm.DB, id int) error {
	if tx == nil {
		tx = s.db
	}
	if err := tx.Where("id = ?", id).Delete(model.SysTableIndex{}).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysTableRepositoryImpl) DeleteTableIndexByTableId(tx *gorm.DB, tableId int) error {
	if tx == nil {
		tx = s.db
	}
	if err := tx.Where("table_id = ?", tableId).Delete(model.SysTableIndex{}).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysTableRepositoryImpl) InsertTableIndexFields(tx *gorm.DB, indexFields []model.SysTableIndexField) error {
	if tx == nil {
		tx = s.db
	}
	if err := tx.Create(&indexFields).Error; err != nil {
		return err
	}
	return nil
}

// CreateTableIndex 删除实体表索引
func (s *SysTableRepositoryImpl) CreateTableIndex(tx *gorm.DB, isUnique bool, indexName string, tableCode string, fields string) error {
	if tx == nil {
		tx = s.db
	}
	var unique string
	if isUnique {
		unique = "UNIQUE"
	}
	tableName := utils.GetTableName(tx, tableCode)
	createIndexSql := fmt.Sprintf("CREATE %s INDEX %s ON %s (%s)", unique, indexName, tableName, fields)
	if err := tx.Exec(createIndexSql).Error; err != nil {
		return err
	}
	return nil
}

// DropTableIndex 删除实体表索引
func (s *SysTableRepositoryImpl) DropTableIndex(tx *gorm.DB, indexName string, tableCode string) error {
	if tx == nil {
		tx = s.db
	}
	tableName := utils.GetTableName(tx, tableCode)
	// 构建删除索引的SQL语句
	dropIndexSQL := fmt.Sprintf("DROP INDEX %s ON %s", indexName, tableName)
	if err := tx.Exec(dropIndexSQL).Error; err != nil {
		return err
	}
	return nil
}

// DeleteTableIndexFieldByIndexId 根据单个indexId删除中间表字段
func (s *SysTableRepositoryImpl) DeleteTableIndexFieldByIndexId(tx *gorm.DB, id int) error {
	if tx == nil {
		tx = s.db
	}
	if err := tx.Where("index_id = ?", id).Delete(model.SysTableIndexField{}).Error; err != nil {
		return err
	}
	return nil
}

// DeleteTableIndexFieldByIndexIds 根据所有indexId删除中间表字段
func (s *SysTableRepositoryImpl) DeleteTableIndexFieldByIndexIds(tx *gorm.DB, ids []int) error {
	if tx == nil {
		tx = s.db
	}
	if err := tx.Where("index_id in ?", ids).Delete(model.SysTableIndexField{}).Error; err != nil {
		return err
	}
	return nil
}

func (s *SysTableRepositoryImpl) FetchTableMetadata(tableSchema string, tableCode string) ([]model.TableColumnMate, error) {
	var columns []model.TableColumnMate
	tableName := utils.GetTableName(s.db, tableCode)
	query := `SELECT *  FROM INFORMATION_SCHEMA.COLUMNS  WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?;`
	err := s.db.Raw(query, tableSchema, tableName).Scan(&columns).Error
	if err != nil {
		return []model.TableColumnMate{}, err
	}
	return columns, nil
}

func (s *SysTableRepositoryImpl) FetchTableIndexMetadata(tableSchema string, tableCode string) ([]model.TableIndexMate, error) {
	var indexes []model.TableIndexMate
	tableName := utils.GetTableName(s.db, tableCode)
	query := `SELECT COLUMN_NAME, INDEX_NAME, NON_UNIQUE, INDEX_TYPE FROM
    INFORMATION_SCHEMA.STATISTICS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ? AND INDEX_NAME != 'PRIMARY';`
	err := s.db.Raw(query, tableSchema, tableName).Scan(&indexes).Error
	if err != nil {
		return []model.TableIndexMate{}, err
	}
	return indexes, nil
}

func (s *SysTableRepositoryImpl) InitTable(tx *gorm.DB, table model.SysTable) error {
	if tx == nil {
		tx = s.db
	}
	// 创建sys_table数据
	if err := tx.Create(&table).Error; err != nil {
		return err
	}
	return nil
}
