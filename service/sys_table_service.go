/**
 * @Author: Nan
 * @Date: 2024/5/17 上午11:30
 */

package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"sweet-cms/cache"
	"sweet-cms/config"
	"sweet-cms/enum"
	"sweet-cms/form/request"
	"sweet-cms/form/response"
	"sweet-cms/inter"
	"sweet-cms/model"
	"sweet-cms/repository"
	"sweet-cms/utils"
)

type SysTableService struct {
	sysTableRepo       repository.SysTableRepository
	sf                 *utils.Snowflake
	sysTableCache      *cache.SysTableCache
	sysTableFieldCache *cache.SysTableFieldCache
	serverConfig       *config.Server
}

func NewSysTableService(
	sysTableRepo repository.SysTableRepository,
	sf *utils.Snowflake,
	sysTableCache *cache.SysTableCache,
	sysTableFieldCache *cache.SysTableFieldCache,
	serverConfig *config.Server,
) *SysTableService {
	return &SysTableService{
		sysTableRepo,
		sf,
		sysTableCache,
		sysTableFieldCache,
		serverConfig,
	}
}

func (s *SysTableService) GetTableById(id int) (model.SysTable, error) {
	data, err := s.sysTableCache.Get(strconv.Itoa(id))
	if err == nil {
		return data, nil
	}
	if !errors.Is(err, inter.ErrCacheMiss) {
		return model.SysTable{}, err
	}
	data, err = s.sysTableRepo.GetTableById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.SysTable{}, nil
		}
		return model.SysTable{}, err
	}
	s.sysTableCache.Set(string(id), data)
	return data, nil
}

func (s *SysTableService) GetTableList(basic request.Basic) (response.ListResult[model.SysTable], error) {
	result, err := s.sysTableRepo.GetTableList(basic)
	return result, err
}

func (s *SysTableService) GetTableByTableCode(code string) (model.SysTable, error) {
	data, err := s.sysTableCache.Get(code)
	if err == nil {
		return data, nil
	}
	if !errors.Is(err, inter.ErrCacheMiss) {
		return model.SysTable{}, err
	}
	data, err = s.sysTableRepo.GetTableByTableCode(code)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.SysTable{}, nil
		}
		return model.SysTable{}, err
	}
	s.sysTableCache.Set(code, data)
	return data, nil
}

func (s *SysTableService) InsertTable(req request.TableCreateReq) error {
	var data model.SysTable
	table, e := s.GetTableByTableCode(req.TableCode)
	if e != nil {
		return e
	}
	if table.Id != 0 {
		e := &response.AdminError{
			Code:    http.StatusBadRequest,
			Message: "当前表已存在，请勿重复创建",
		}
		return e
	}
	err := mapstructure.Decode(req, &data)
	if err != nil {
		fmt.Println("Error during struct mapping:", err)
		return err
	}
	id, err := s.sf.GenerateUniqueID()
	if err != nil {
		return err
	}
	data.Id = int(id)
	return s.sysTableRepo.InsertTable(data)
}

func (s *SysTableService) UpdateTable(req request.TableUpdateReq) error {
	err := s.sysTableRepo.UpdateTable(req)
	if err != nil {
		return err
	}
	data, err := s.GetTableById(req.Id)
	if err != nil {
		return err
	}
	if data.Id != 0 {
		s.sysTableCache.Delete(strconv.Itoa(data.Id))
		s.sysTableCache.Delete(data.TableCode)
	}
	return nil
}

func (s *SysTableService) DeleteTableById(id int) error {
	data, err := s.GetTableById(id)
	if err != nil {
		return err
	}
	err = s.sysTableRepo.DeleteTableById(id)
	if err != nil {
		return err
	}
	if data.Id != 0 {
		s.sysTableCache.Delete(strconv.Itoa(data.Id))
		s.sysTableCache.Delete(data.TableCode)
	}
	return nil
}

func (s *SysTableService) GetTableFieldById(id int) (model.SysTableField, error) {
	data, err := s.sysTableFieldCache.Get(strconv.Itoa(id))
	if err == nil {
		return data, nil
	}
	if !errors.Is(err, inter.ErrCacheMiss) {
		return model.SysTableField{}, err
	}
	data, err = s.sysTableRepo.GetTableFieldById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.SysTableField{}, nil
		}
		return model.SysTableField{}, err
	}
	s.sysTableFieldCache.Set(strconv.Itoa(id), data)
	return data, nil
}

func (s *SysTableService) GetTableFieldsByTableId(tableId int) ([]model.SysTableField, error) {
	data, err := s.sysTableCache.Get(strconv.Itoa(tableId))
	if err == nil {
		return data.TableFields, nil
	}
	if !errors.Is(err, inter.ErrCacheMiss) {
		return nil, err
	}
	fields, err := s.sysTableRepo.GetTableFieldsByTableId(tableId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.SysTableField{}, nil
		}
		return nil, err
	}
	return fields, nil
}

func (s *SysTableService) InsertTableField(req request.TableFieldCreateReq) error {
	var data model.SysTableField
	fields, e := s.GetTableFieldsByTableId(req.TableId)
	if e != nil {
		return e
	}
	for _, field := range fields {
		if field.FieldCode == req.FieldCode {
			e = &response.AdminError{
				Code:    http.StatusBadRequest,
				Message: "该字段已存在，请勿重复创建",
			}
			return e
		}
	}
	err := mapstructure.Decode(req, &data)
	if err != nil {
		zap.L().Error("Error during struct mapping:", zap.Error(err))
		return err
	}
	id, err := s.sf.GenerateUniqueID()
	if err != nil {
		return err
	}
	table, err := s.GetTableById(data.TableId)
	if err != nil {
		return err
	}

	data.Id = int(id)
	err = s.sysTableRepo.InsertTableField(data, table.TableCode)
	if err != nil {
		return err
	}

	if table.Id != 0 {
		s.sysTableCache.Delete(strconv.Itoa(table.Id))
		s.sysTableCache.Delete(table.TableCode)
	}
	return nil
}

func (s *SysTableService) UpdateTableField(req request.TableFieldUpdateReq) error {
	table, err := s.GetTableById(req.Id)
	if err != nil {
		return err
	}
	if table.Id != 0 {
		fields, e := s.GetTableFieldsByTableId(req.TableId)
		if e != nil {
			return e
		}
		var data model.SysTableField
		for _, field := range fields {
			if field.Id == req.Id {
				diff := cmp.Diff(req, field)
				if diff == "" {
					return &response.AdminError{
						Code:    http.StatusBadRequest,
						Message: "字段未发生变化，无需更新",
					}
				}
				zap.L().Info("变化值：", zap.String("diff", diff))
				data = field
				break
			}
		}
		err = s.sysTableRepo.UpdateTableField(req, data, table.TableCode)
		if err != nil {
			return err
		}
		s.sysTableCache.Delete(strconv.Itoa(table.Id))
		s.sysTableCache.Delete(table.TableCode)
		s.sysTableFieldCache.Delete(strconv.Itoa(req.Id))
		return nil
	}
	return errors.New("数据不存在")
}

func (s *SysTableService) DeleteTableFieldById(id int) error {
	field, err := s.GetTableFieldById(id)
	if err != nil {
		return err
	}
	if field.Id != 0 {
		table, err := s.GetTableById(field.TableId)
		if err != nil {
			return err
		}
		if table.Id != 0 {
			err = s.sysTableRepo.DeleteTableField(field, table.TableCode)
			if err != nil {
				return err
			}
			s.sysTableFieldCache.Delete(strconv.Itoa(field.Id))
			s.sysTableCache.Delete(strconv.Itoa(table.Id))
			s.sysTableCache.Delete(table.TableCode)
			return nil
		}
	}
	return errors.New("数据不存在")
}

func (s *SysTableService) InitTable(ctx *gin.Context, tableCode string) error {
	columns, err := s.sysTableRepo.FetchTableMetadata(s.serverConfig.DB.Name, s.serverConfig.DB.Prefix+tableCode)
	tableIndexes, err := s.sysTableRepo.FetchTableIndexes(s.serverConfig.DB.Name, s.serverConfig.DB.Prefix+tableCode)
	fields, err := ConvertColumnsToSysTableFields(columns)
	if err != nil {
		return err
	}
	var user model.SysUser
	obj, exists := ctx.Get("user")
	if exists {
		user, _ = obj.(model.SysUser)
	}
	id, err := s.sf.GenerateUniqueID()
	if err != nil {
		return err
	}
	table := model.SysTable{
		Basic: model.Basic{
			Id:            int(id),
			GmtCreateUser: &user.EmployeeId,
		},
		TableName: tableCode,
		TableCode: tableCode,
		TableType: enum.SYSTEM,
	}
	indexesMap := make(map[string]model.SysTableIndex)
	var indexes []model.SysTableIndex
	var indexFields []model.SysTableIndexField
	for i, _ := range fields {
		fields[i].TableId = table.Id
		fieldId, err := s.sf.GenerateUniqueID()
		if err != nil {
			return err
		}
		fields[i].Id = int(fieldId)
		fields[i].GmtCreateUser = &user.EmployeeId

		for j, _ := range tableIndexes {
			if tableIndexes[j].ColumnName == fields[i].FieldCode {
				indexId, err := s.sf.GenerateUniqueID()
				if err != nil {
					return err
				}
				if _, exists := indexesMap[tableIndexes[j].IndexName]; !exists {
					indexesMap[tableIndexes[j].IndexName] = model.SysTableIndex{
						Basic: model.Basic{
							Id:            int(indexId),
							GmtCreateUser: &user.EmployeeId,
						},
						TableId:   table.Id,
						IndexName: tableIndexes[j].IndexName,
						IsUnique:  !tableIndexes[j].NonUnique,
					}
					indexes = append(indexes, indexesMap[tableIndexes[j].IndexName])
				} else {
					indexId = int64(indexesMap[tableIndexes[j].IndexName].Id)
				}
				indexFields = append(indexFields, model.SysTableIndexField{
					IndexId: int(indexId),
					FieldId: fields[i].Id,
				})
			}
		}
	}
	table.TableFields = fields
	table.TableIndexes = indexes
	err = s.sysTableRepo.InitTable(table, indexFields)
	return err
}

func ConvertColumnsToSysTableFields(columns []model.TableColumn) ([]model.SysTableField, error) {
	var fields []model.SysTableField
	for _, column := range columns {
		field := model.SysTableField{
			FieldCode:          column.ColumnName,              // 通常 FieldCode 会是数据库的真实列名
			FieldDecimalLength: int(column.NumericScale.Int64), // 根据需要设置
			IsNull:             column.IsNullable == "YES",
			IsPrimaryKey:       column.ColumnKey == "PRI",
			IsQuickSearch:      false,
			IsAdvancedSearch:   false,
			IsSort:             true,
			IsListShow:         true,
			IsInsertShow:       false,
			IsUpdateShow:       false,
			Sequence:           uint8(column.OrdinalPosition),
			OriginalFieldId:    0,
			FieldLength:        0,
			FieldCategory:      enum.NORMAL_FIELD,
			Binding:            "required", // 根据实际逻辑调整
		}
		if column.ColumnComment != "" {
			field.FieldName = column.ColumnComment
		} else {
			field.FieldName = column.ColumnName
		}
		switch column.DataType {
		case "int", "bigint":
			field.FieldType = enum.INT
		case "tinyint":
			field.FieldType = enum.TINYINT
			field.FieldLength = int(column.NumericPrecision.Int64)
		case "varchar":
			field.FieldType = enum.VARCHAR
			field.FieldLength = int(column.CharacterMaximumLength.Int64)
		case "text", "mediumtext", "longtext":
			field.FieldType = enum.TEXT
			field.FieldLength = int(column.CharacterMaximumLength.Int64)
		case "boolean", "bool":
			field.FieldType = enum.BOOLEAN
		case "date":
			field.FieldType = enum.DATE
		case "datetime", "timestamp":
			field.FieldType = enum.DATETIME
		case "time":
			field.FieldType = enum.TIME
		default:
			field.FieldType = enum.VARCHAR
			field.FieldLength = int(column.NumericPrecision.Int64)
		}
		// 检查DefaultValue是否有值
		if column.ColumnDefault.Valid {
			field.DefaultValue = &column.ColumnDefault.String
		}
		fields = append(fields, field)
	}
	return fields, nil
}
