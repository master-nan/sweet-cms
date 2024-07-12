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
	"reflect"
	"strconv"
	"strings"
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

func (s *SysTableService) InsertTable(ctx *gin.Context, req request.TableCreateReq) error {
	var user model.SysUser
	obj, exists := ctx.Get("user")
	if exists {
		user, _ = obj.(model.SysUser)
	}
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
	data.CreateUser = &user.EmployeeId
	// 自动在sys_table_field中为Basic结构体中的每个字段创建记录
	fields := []model.SysTableField{
		{Basic: model.Basic{CreateUser: data.CreateUser}, TableId: data.Id, FieldName: "id", FieldCode: "id", FieldType: enum.INT, IsPrimaryKey: true, IsNull: false, InputType: enum.INPUT_NUMBER, IsSort: true, Sequence: 1, IsListShow: true},
		{Basic: model.Basic{CreateUser: data.CreateUser}, TableId: data.Id, FieldName: "创建时间", FieldCode: "gmt_create", FieldType: enum.DATETIME, IsNull: false, InputType: enum.DATETIME_PICKER, IsSort: true, Sequence: 2, IsListShow: true},
		{Basic: model.Basic{CreateUser: data.CreateUser}, TableId: data.Id, FieldName: "创建者", FieldCode: "gmt_create_user", FieldType: enum.INT, IsNull: false, InputType: enum.INPUT_NUMBER, Sequence: 3, IsListShow: true},
		{Basic: model.Basic{CreateUser: data.CreateUser}, TableId: data.Id, FieldName: "修改时间", FieldCode: "gmt_modify", FieldType: enum.DATETIME, IsNull: false, InputType: enum.DATETIME_PICKER, IsSort: true, Sequence: 4, IsListShow: true},
		{Basic: model.Basic{CreateUser: data.CreateUser}, TableId: data.Id, FieldName: "修改者", FieldCode: "gmt_modify_user", FieldType: enum.INT, IsNull: false, InputType: enum.INPUT_NUMBER, Sequence: 5, IsListShow: true},
		{Basic: model.Basic{CreateUser: data.CreateUser}, TableId: data.Id, FieldName: "删除时间", FieldCode: "gmt_delete", FieldType: enum.DATETIME, IsNull: true, InputType: enum.DATETIME_PICKER, Sequence: 6},
		{Basic: model.Basic{CreateUser: data.CreateUser}, TableId: data.Id, FieldName: "删除者", FieldCode: "gmt_delete_user", FieldType: enum.INT, IsNull: true, InputType: enum.INPUT_NUMBER, Sequence: 7},
		{Basic: model.Basic{CreateUser: data.CreateUser}, TableId: data.Id, FieldName: "状态", FieldCode: "state", FieldType: enum.BOOLEAN, IsNull: false, InputType: enum.SELECT, IsSort: true, DefaultValue: utils.StringPtr("true"), DictCode: utils.StringPtr("whether"), Sequence: 8, IsListShow: true},
	}
	for i := range fields {
		fieldId, err := s.sf.GenerateUniqueID()
		if err != nil {
			return err
		}
		fields[i].Id = int(fieldId)
	}
	data.TableFields = fields
	return s.sysTableRepo.ExecuteTx(ctx, func(tx *gorm.DB) error {
		if e := s.sysTableRepo.InsertTable(tx, data); e != nil {
			return e
		}
		// 动态创建结构体类型
		dynamicType := utils.CreateDynamicStruct(table.TableFields)
		// 创建实例
		dynamicModel := reflect.New(dynamicType).Interface()
		if e := s.sysTableRepo.CreateTable(tx, table.TableCode, dynamicModel); e != nil {
			return e
		}
		return nil
	})
}

func (s *SysTableService) UpdateTable(req request.TableUpdateReq) error {
	err := s.sysTableRepo.UpdateTable(req)
	if err != nil {
		return err
	}
	// 删除缓存
	s.DeleteCache(req.Id)
	return nil
}

func (s *SysTableService) DeleteTableById(ctx *gin.Context, id int) error {
	err := s.sysTableRepo.ExecuteTx(ctx, func(tx *gorm.DB) error {
		if e := s.sysTableRepo.DeleteTableById(tx, id); e != nil {
			return e
		}
		// 删除字段元数据
		if e := s.sysTableRepo.DeleteTableFieldByTableId(tx, id); e != nil {
			return e
		}
		// 查询表所有索引
		tableIndexes, e := s.sysTableRepo.GetTableIndexesByTableId(id)
		if e != nil && !errors.Is(e, gorm.ErrRecordNotFound) {
			return e
		}
		// 删除索引信息
		if e := s.sysTableRepo.DeleteTableIndexByTableId(tx, id); e != nil {
			return e
		}
		var indexIDs []int
		for _, index := range tableIndexes {
			indexIDs = append(indexIDs, index.Id)
		}
		// 删除索引中间表信息，需要使用 IN 查询
		if len(indexIDs) > 0 {
			if e := s.sysTableRepo.DeleteTableIndexFieldByIndexIds(tx, indexIDs); e != nil {
				return e
			}
		}
		// 查询关联表数据
		relations, e := s.sysTableRepo.GetTableRelationsByTableId(id)
		if e != nil && !errors.Is(e, gorm.ErrRecordNotFound) {
			return e
		}
		for _, relation := range relations {
			// 删除关联关系表
			if e := s.sysTableRepo.DeleteTableRelation(tx, relation.Id); e != nil {
				return e
			}
			if relation.RelationType == enum.MANY_TO_MANY {
				// 删除多对多中间表
				if e := s.sysTableRepo.DropTable(tx, relation.ManyTableCode); e != nil {
					return e
				}
			}
		}
		return nil
	})
	// 删除缓存
	s.DeleteCache(id)
	return err
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

func (s *SysTableService) InsertTableField(ctx *gin.Context, req request.TableFieldCreateReq) error {
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
	err = s.sysTableRepo.ExecuteTx(ctx, func(tx *gorm.DB) error {
		if e := s.sysTableRepo.InsertTableField(tx, data); err != nil {
			return e
		}
		// 构建SQL类型字符串，包括长度、默认值、是否可为空和备注
		sqlType := utils.SqlTypeFromFieldType(data.FieldType)
		if data.FieldLength > 0 {
			sqlType += fmt.Sprintf("(%d)", data.FieldLength)
		}
		if data.DefaultValue != nil {
			sqlType += fmt.Sprintf(" DEFAULT '%s'", data.DefaultValue)
		}
		if data.IsNull {
			sqlType += " NULL"
		} else {
			sqlType += " NOT NULL"
		}
		if data.FieldName != "" {
			sqlType += fmt.Sprintf(" COMMENT '%s'", data.FieldName)
		}
		if e := s.sysTableRepo.CreateTableColumn(tx, table.TableCode, data.FieldCode, sqlType); e != nil {
			return e
		}
		return nil
	})
	// 删除缓存
	s.DeleteCache(table.Id)
	return err
}

func (s *SysTableService) UpdateTableField(ctx *gin.Context, req request.TableFieldUpdateReq) error {
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
		err = s.sysTableRepo.ExecuteTx(ctx, func(tx *gorm.DB) error {
			if e := s.sysTableRepo.UpdateTableField(tx, req); e != nil {
				return e
			}
			var sqlType string
			if req.FieldType != data.FieldType || (req.FieldLength > 0 && req.FieldLength != data.FieldLength) {
				sqlType += fmt.Sprintf("%s(%d)", utils.SqlTypeFromFieldType(req.FieldType), req.FieldLength)
			}
			if req.DefaultValue != "" && req.DefaultValue != *data.DefaultValue {
				sqlType += fmt.Sprintf(" DEFAULT '%s'", req.DefaultValue)
			}
			if req.IsNull != data.IsNull {
				if req.IsNull {
					sqlType += " NULL"
				} else {
					sqlType += " NOT NULL"
				}
			}
			if req.FieldName != "" && req.FieldName != data.FieldName {
				sqlType += fmt.Sprintf(" COMMENT '%s'", req.FieldName)
			}
			if req.FieldCode == data.FieldCode {
				if sqlType != "" {
					if e := s.sysTableRepo.ModifyTableColumn(tx, table.TableCode, req.FieldCode, sqlType); e != nil {
						return e
					}
				}
			} else {
				if err := s.sysTableRepo.ChangeTableColumn(tx, table.TableCode, data.FieldCode, req.FieldCode, sqlType); err != nil {
					return e
				}
			}
			return nil
		})
		// 删除缓存
		s.DeleteCache(table.Id)
		return err
	}
	return errors.New("数据不存在")
}

func (s *SysTableService) DeleteTableFieldById(ctx *gin.Context, id int) error {
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
			err = s.sysTableRepo.ExecuteTx(ctx, func(tx *gorm.DB) error {
				if e := s.sysTableRepo.DeleteTableField(tx, field.Id); e != nil {
					return e
				}
				if e := s.sysTableRepo.DropTableColumn(tx, table.TableCode, field.FieldCode); e != nil {
					return e
				}
				return nil
			})
			if err != nil {
				return err
			}
			// 删除缓存
			s.DeleteCache(table.Id)
			return nil
		}
	}
	return errors.New("数据不存在")
}

func (s *SysTableService) GetTableRelationsByTableId(tableId int) ([]model.SysTableRelation, error) {
	return s.sysTableRepo.GetTableRelationsByTableId(tableId)
}

func (s *SysTableService) GetTableRelationById(id int) (model.SysTableRelation, error) {
	return s.sysTableRepo.GetTableRelationById(id)
}

func (s *SysTableService) InsertTableRelation(ctx *gin.Context, req request.TableRelationCreateReq) error {
	err := s.sysTableRepo.ExecuteTx(ctx, func(tx *gorm.DB) error {
		var data model.SysTableRelation
		e := mapstructure.Decode(req, &data)
		if e != nil {
			zap.L().Error("Error during struct mapping:", zap.Error(e))
			return e
		}
		id, err := s.sf.GenerateUniqueID()
		if err != nil {
			return e
		}
		data.Id = int(id)
		if e := s.sysTableRepo.InsertTableRelation(tx, data); e != nil {
			return e
		}
		// 如果是多对多 创建对应的表
		if data.RelationType == enum.MANY_TO_MANY {
			mainTable, e := s.GetTableById(data.TableId)
			if e != nil {
				return e
			}
			relatedTable, e := s.GetTableById(data.RelatedTableId)
			if e != nil {
				return e
			}
			var referenceKeyField model.SysTableField
			for _, field := range mainTable.TableFields {
				if field.FieldCode == data.ReferenceKey {
					referenceKeyField = field
				}
			}
			var foreignKeyField model.SysTableField
			for _, field := range relatedTable.TableFields {
				if field.FieldCode == data.ForeignKey {
					foreignKeyField = field
				}
			}
			if referenceKeyField.Id == 0 || foreignKeyField.Id == 0 {
				return errors.New("关联字段不存在")
			}
			var relationList []reflect.StructField
			referenceKey := reflect.StructField{
				Name: data.ReferenceKey,
				Type: utils.GetFieldType(referenceKeyField.FieldType),
				Tag:  reflect.StructTag(`gorm:"primaryKey;autoIncrement:false"`),
			}
			foreignKey := reflect.StructField{
				Name: data.ForeignKey,
				Type: utils.GetFieldType(foreignKeyField.FieldType),
				Tag:  reflect.StructTag(`gorm:"primaryKey;autoIncrement:false"`),
			}
			relationList = append(relationList, referenceKey, foreignKey)
			reflect.StructOf(relationList)
			relationModel := reflect.New(reflect.StructOf(relationList)).Interface()
			return s.sysTableRepo.CreateTable(tx, data.ManyTableCode, relationModel)
		}
		return nil
	})
	// 删除缓存
	s.DeleteCache(req.TableId)
	return err
}
func (s *SysTableService) DeleteTableRelation(ctx *gin.Context, id int) error {
	err := s.sysTableRepo.ExecuteTx(ctx, func(tx *gorm.DB) error {
		if e := s.sysTableRepo.DeleteTableRelation(tx, id); e != nil {
			return e
		}
		relation, e := s.sysTableRepo.GetTableRelationById(id)
		if e != nil && !errors.Is(e, gorm.ErrRecordNotFound) {
			return e
		}
		// 删除缓存
		s.DeleteCache(relation.TableId)
		// TODO 删除表关系考虑是否需要删除多对多中间表
		//if relation.RelationType == enum.MANY_TO_MANY {
		//	if e := s.sysTableRepo.DropTable(tx, relation.ManyTableCode); e {
		//		return e
		//	}
		//}
		return nil
	})
	return err
}

func (s *SysTableService) GetTableIndexesByTableId(tableId int) ([]model.SysTableIndex, error) {
	return s.sysTableRepo.GetTableIndexesByTableId(tableId)
}

func (s *SysTableService) InsertTableIndex(ctx *gin.Context, req request.TableIndexCreateReq) error {
	err := s.sysTableRepo.ExecuteTx(ctx, func(tx *gorm.DB) error {
		var data model.SysTableIndex
		e := mapstructure.Decode(req, &data)
		if e != nil {
			zap.L().Error("Error during struct mapping:", zap.Error(e))
			return e
		}
		id, err := s.sf.GenerateUniqueID()
		if err != nil {
			return e
		}
		data.Id = int(id)
		if e := s.sysTableRepo.InsertTableIndex(tx, data); e != nil {
			return e
		}
		var indexFields []model.SysTableIndexField
		fieldCodeList := make([]string, len(req.IndexFields))
		for _, field := range req.IndexFields {
			fieldCodeList = append(fieldCodeList, field.FieldCode)
			indexField := model.SysTableIndexField{
				IndexId: data.Id,
				FieldId: field.FieldId,
			}
			indexFields = append(indexFields, indexField)
		}
		if e := s.sysTableRepo.InsertTableIndexFields(tx, indexFields); e != nil {
			return e
		}
		table, err := s.GetTableById(data.TableId)
		if err != nil {
			return err
		}
		if table.Id == 0 {
			return errors.New("操作的表不存在")
		}
		fields := strings.Join(fieldCodeList, ",")
		if e := s.sysTableRepo.CreateTableIndex(tx, req.IsUnique, req.IndexName, table.TableCode, fields); e != nil {
			return e
		}
		return nil
	})
	// 删除缓存
	s.DeleteCache(req.TableId)
	return err
}

func (s *SysTableService) UpdateTableIndex(ctx *gin.Context, req request.TableIndexUpdateReq) error {
	err := s.sysTableRepo.ExecuteTx(ctx, func(tx *gorm.DB) error {
		// 删除中间表数据
		if e := s.sysTableRepo.DeleteTableIndexFieldByIndexId(tx, req.Id); e != nil {
			return e
		}
		if e := s.sysTableRepo.UpdateTableIndex(tx, req); e != nil {
			return e
		}
		table, e := s.GetTableById(req.TableId)
		if e != nil {
			return e
		}
		// 使用原索引名称删除表索引
		if e := s.sysTableRepo.DropTableIndex(tx, req.IndexName, table.TableCode); e != nil {
			return e
		}
		var indexFields []model.SysTableIndexField
		fieldCodeList := make([]string, len(req.IndexFields))
		for _, field := range req.IndexFields {
			fieldCodeList = append(fieldCodeList, field.FieldCode)
			indexField := model.SysTableIndexField{
				IndexId: req.Id,
				FieldId: field.FieldId,
			}
			indexFields = append(indexFields, indexField)
		}
		// 创建中间表数据
		if e := s.sysTableRepo.InsertTableIndexFields(tx, indexFields); e != nil {
			return e
		}
		fields := strings.Join(fieldCodeList, ",")
		// 创建表索引
		if err := s.sysTableRepo.CreateTableIndex(tx, req.IsUnique, req.IndexName, table.TableCode, fields); err != nil {
			return err
		}
		return nil
	})
	// 删除缓存
	s.DeleteCache(req.TableId)
	return err
}

func (s *SysTableService) DeleteTableIndex(ctx *gin.Context, id int) error {
	index, e := s.sysTableRepo.GetTableIndexById(id)
	if e != nil {
		return e
	}
	table, e := s.GetTableById(index.TableId)
	if e != nil {
		return e
	}
	err := s.sysTableRepo.ExecuteTx(ctx, func(tx *gorm.DB) error {
		if e := s.sysTableRepo.DeleteTableIndex(tx, id); e != nil {
			return e
		}
		// 使用索引名称删除表索引
		if e := s.sysTableRepo.DropTableIndex(tx, index.IndexName, table.TableCode); e != nil {
			return e
		}
		return nil
	})
	// 删除缓存
	s.DeleteCache(table.Id)
	return err
}

func (s *SysTableService) DeleteTableIndexByTableId(ctx *gin.Context, id int) error {
	table, e := s.GetTableById(id)
	if e != nil {
		return e
	}
	indexes, e := s.sysTableRepo.GetTableIndexesByTableId(id)
	if e != nil {
		return e
	}
	err := s.sysTableRepo.ExecuteTx(ctx, func(tx *gorm.DB) error {
		if e := s.sysTableRepo.DeleteTableIndexByTableId(tx, id); e != nil {
			return e
		}
		for _, index := range indexes {
			// 使用索引名称删除表索引
			if e := s.sysTableRepo.DropTableIndex(tx, index.IndexName, table.TableCode); e != nil {
				return e
			}
		}
		return nil
	})
	// 删除缓存
	s.DeleteCache(table.Id)
	return err
}

func (s *SysTableService) InitTable(ctx *gin.Context, tableCode string) error {
	columns, err := s.sysTableRepo.FetchTableMetadata(s.serverConfig.DB.Name, tableCode)
	tableIndexes, err := s.sysTableRepo.FetchTableIndexMetadata(s.serverConfig.DB.Name, tableCode)
	fields, err := utils.ConvertColumnsToSysTableFields(columns)
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
	table, e := s.GetTableByTableCode(tableCode)
	if err != nil {
		return err
	}
	if e != nil {
		return e
	}
	if table.Id != 0 {
		e := &response.AdminError{
			Code:    http.StatusBadRequest,
			Message: "当前表已初始化，请勿重复操作",
		}
		return e
	}
	table = model.SysTable{
		Basic: model.Basic{
			Id:         int(id),
			CreateUser: &user.EmployeeId,
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
		fields[i].CreateUser = &user.EmployeeId

		for j, _ := range tableIndexes {
			if tableIndexes[j].ColumnName == fields[i].FieldCode {
				indexId, err := s.sf.GenerateUniqueID()
				if err != nil {
					return err
				}
				if _, exists := indexesMap[tableIndexes[j].IndexName]; !exists {
					indexesMap[tableIndexes[j].IndexName] = model.SysTableIndex{
						Basic: model.Basic{
							Id:         int(indexId),
							CreateUser: &user.EmployeeId,
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
	return s.sysTableRepo.ExecuteTx(ctx, func(tx *gorm.DB) error {
		if e := s.sysTableRepo.InsertTable(tx, table); e != nil {
			return e
		}
		if indexFields != nil && len(indexFields) > 0 {
			if e := s.sysTableRepo.InsertTableIndexFields(tx, indexFields); e != nil {
				return e
			}
		}
		return nil
	})
}

func (s *SysTableService) DeleteCache(tableId int) {
	go func() {
		table, _ := s.GetTableById(tableId)
		if table.Id != 0 {
			s.sysTableCache.Delete(strconv.Itoa(table.Id))
			s.sysTableCache.Delete(table.TableCode)
			for _, field := range table.TableFields {
				s.sysTableFieldCache.Delete(strconv.Itoa(field.Id))
			}
		}
	}()
}
