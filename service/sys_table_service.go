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
	sysTableRepo           repository.SysTableRepository
	sysTableFieldRepo      repository.SysTableFieldRepository
	sysTableIndexRepo      repository.SysTableIndexRepository
	sysTableIndexFieldRepo repository.SysTableIndexFieldRepository
	sysTableRelationRepo   repository.SysTableRelationRepository
	sf                     *utils.Snowflake
	sysTableCache          *cache.SysTableCache
	sysTableFieldCache     *cache.SysTableFieldCache
	serverConfig           *config.Server
}

func NewSysTableService(
	sysTableRepo repository.SysTableRepository,
	sysTableFieldRepo repository.SysTableFieldRepository,
	sysTableIndexRepo repository.SysTableIndexRepository,
	sysTableIndexFieldRepo repository.SysTableIndexFieldRepository,
	sysTableRelationRepo repository.SysTableRelationRepository,
	sf *utils.Snowflake,
	sysTableCache *cache.SysTableCache,
	sysTableFieldCache *cache.SysTableFieldCache,
	serverConfig *config.Server,
) *SysTableService {
	return &SysTableService{
		sysTableRepo,
		sysTableFieldRepo,
		sysTableIndexRepo,
		sysTableIndexFieldRepo,
		sysTableRelationRepo,
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

func (s *SysTableService) CreateTable(ctx *gin.Context, req request.TableCreateReq) error {
	var data model.SysTable
	table, e := s.GetTableByTableCode(req.TableCode)
	if e != nil {
		return e
	}
	if table.Id != 0 {
		e := &response.AdminError{
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: "当前表已存在，请勿重复创建",
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
	// 自动在sys_table_field中为Basic结构体中的每个字段创建记录
	fields := []model.SysTableField{
		{TableId: data.Id, FieldName: "id", FieldCode: "id", FieldType: enum.IntFieldType, IsPrimaryKey: true, IsNull: false, InputType: enum.InputNumberInputType, IsSort: true, Sequence: 1, IsListShow: true},
		{TableId: data.Id, FieldName: "创建时间", FieldCode: "gmt_create", FieldType: enum.DatetimeFieldType, IsNull: false, InputType: enum.DatetimePickerInputType, IsSort: true, Sequence: 2, IsListShow: true},
		{TableId: data.Id, FieldName: "创建者", FieldCode: "gmt_create_user", FieldType: enum.IntFieldType, IsNull: false, InputType: enum.InputNumberInputType, Sequence: 3, IsListShow: true},
		{TableId: data.Id, FieldName: "修改时间", FieldCode: "gmt_modify", FieldType: enum.DatetimeFieldType, IsNull: false, InputType: enum.DatetimePickerInputType, IsSort: true, Sequence: 4, IsListShow: true},
		{TableId: data.Id, FieldName: "修改者", FieldCode: "gmt_modify_user", FieldType: enum.IntFieldType, IsNull: false, InputType: enum.InputNumberInputType, Sequence: 5, IsListShow: true},
		{TableId: data.Id, FieldName: "删除时间", FieldCode: "gmt_delete", FieldType: enum.DatetimeFieldType, IsNull: true, InputType: enum.DatetimePickerInputType, Sequence: 6},
		{TableId: data.Id, FieldName: "删除者", FieldCode: "gmt_delete_user", FieldType: enum.IntFieldType, IsNull: true, InputType: enum.InputNumberInputType, Sequence: 7},
		{TableId: data.Id, FieldName: "状态", FieldCode: "state", FieldType: enum.BooleanFieldType, IsNull: false, InputType: enum.SelectInputType, IsSort: true, DefaultValue: utils.StringPtr("true"), DictCode: utils.StringPtr("whether"), Sequence: 8, IsListShow: true},
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
		if e := s.sysTableRepo.Create(tx, &data); e != nil {
			return e
		}
		// 创建实例
		dynamicModel := s.sysTableRepo.Model(data.TableFields)
		// 先删除再创建
		if e := s.sysTableRepo.DropTable(tx, data.TableCode); e != nil {
			return e
		}
		if e := s.sysTableRepo.CreateTable(tx, data.TableCode, dynamicModel); e != nil {
			return e
		}
		return nil
	})
}

func (s *SysTableService) UpdateTable(ctx *gin.Context, req request.TableUpdateReq) error {
	tx := s.sysTableRepo.DBWithContext(ctx)
	err := s.sysTableRepo.Update(tx, &req)
	if err != nil {
		return err
	}
	// 删除缓存
	s.DeleteCache(req.Id)
	return nil
}

func (s *SysTableService) DeleteTableById(ctx *gin.Context, id int) error {
	err := s.sysTableRepo.ExecuteTx(ctx, func(tx *gorm.DB) error {
		if e := s.sysTableRepo.DeleteById(tx, id); e != nil {
			return e
		}
		// 删除字段元数据
		if e := s.sysTableFieldRepo.DeleteByField(tx, "table_id", id); e != nil {
			return e
		}
		// 查询表所有索引
		tableIndexes, e := s.sysTableIndexRepo.GetTableIndexesByTableId(id)
		if e != nil && !errors.Is(e, gorm.ErrRecordNotFound) {
			return e
		}
		// 删除索引信息
		if e := s.sysTableIndexRepo.DeleteByField(tx, "table_id", id); e != nil {
			return e
		}
		var indexIDs []int
		for _, index := range tableIndexes {
			indexIDs = append(indexIDs, index.Id)
		}
		// 删除索引中间表信息，需要使用 IN 查询
		if len(indexIDs) > 0 {
			slice := utils.ToInterfaceSlice(indexIDs)
			if e := s.sysTableIndexFieldRepo.DeleteByFieldIn(tx, "index_id", slice); e != nil {
				return e
			}
		}
		// 查询关联表数据
		relations, e := s.sysTableRelationRepo.GetTableRelationsByTableId(id)
		if e != nil && !errors.Is(e, gorm.ErrRecordNotFound) {
			return e
		}
		for _, relation := range relations {
			// 删除关联关系表
			if e := s.sysTableRelationRepo.DeleteById(tx, relation.Id); e != nil {
				return e
			}
			if relation.RelationType == enum.ManyToMany {
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
	result, err := s.sysTableFieldRepo.FindById(id)
	data = result.(model.SysTableField)
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
	fields, err := s.sysTableFieldRepo.GetTableFieldsByTableId(tableId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.SysTableField{}, nil
		}
		return nil, err
	}
	return fields, nil
}

func (s *SysTableService) CreateTableField(ctx *gin.Context, req request.TableFieldCreateReq) error {
	var data model.SysTableField
	fields, e := s.GetTableFieldsByTableId(req.TableId)
	if e != nil {
		return e
	}
	for _, field := range fields {
		if field.FieldCode == req.FieldCode {
			e = &response.AdminError{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: "该字段已存在，请勿重复创建",
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
		if e := s.sysTableFieldRepo.Create(tx, &data); err != nil {
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
						ErrorCode:    http.StatusBadRequest,
						ErrorMessage: "字段未发生变化，无需更新",
					}
				}
				zap.L().Info("变化值：", zap.String("diff", diff))
				data = field
				break
			}
		}
		err = s.sysTableRepo.ExecuteTx(ctx, func(tx *gorm.DB) error {
			if e := s.sysTableFieldRepo.Update(tx, &req); e != nil {
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
				if e := s.sysTableFieldRepo.DeleteById(tx, field.Id); e != nil {
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
	return s.sysTableRelationRepo.GetTableRelationsByTableId(tableId)
}

func (s *SysTableService) GetTableRelationById(id int) (model.SysTableRelation, error) {
	result, err := s.sysTableRelationRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.SysTableRelation{}, nil
		}
		return model.SysTableRelation{}, err
	}
	data := result.(model.SysTableRelation)
	return data, nil
}

func (s *SysTableService) CreateTableRelation(ctx *gin.Context, req request.TableRelationCreateReq) error {
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
		if e := s.sysTableRelationRepo.Create(tx, &data); e != nil {
			return e
		}
		// 如果是多对多 创建对应的表
		if data.RelationType == enum.ManyToMany {
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
					referenceKeyField.Tag = utils.StringPtr(`gorm:"primaryKey;autoIncrement:false"`)
				}
			}
			var foreignKeyField model.SysTableField
			for _, field := range relatedTable.TableFields {
				if field.FieldCode == data.ForeignKey {
					foreignKeyField = field
					foreignKeyField.Tag = utils.StringPtr(`gorm:"primaryKey;autoIncrement:false"`)
				}
			}
			if referenceKeyField.Id == 0 || foreignKeyField.Id == 0 {
				return errors.New("关联字段不存在")
			}

			//var relationList []reflect.StructField
			//referenceKey := reflect.StructField{
			//	Name: data.ReferenceKey,
			//	Type: util.GetFieldType(referenceKeyField.FieldType),
			//	Tag:  reflect.StructTag(`gorm:"primaryKey;autoIncrement:false"`),
			//}
			//foreignKey := reflect.StructField{
			//	Name: data.ForeignKey,
			//	Type: util.GetFieldType(foreignKeyField.FieldType),
			//	Tag:  reflect.StructTag(`gorm:"primaryKey;autoIncrement:false"`),
			//}
			//relationList = append(relationList, referenceKey, foreignKey)
			//reflect.StructOf(relationList)
			//relationModel := reflect.New(reflect.StructOf(relationList)).Interface()

			var relationFields []model.SysTableField
			relationFields = append(relationFields, referenceKeyField, foreignKeyField)
			relationModel := s.sysTableRepo.Model(relationFields)
			// 先删除再创建
			if e := s.sysTableRepo.DropTable(tx, data.ManyTableCode); e != nil {
				return e
			}
			return s.sysTableRepo.CreateTable(tx, data.ManyTableCode, relationModel)
		}
		return nil
	})
	// 删除缓存
	s.DeleteCache(req.TableId)
	return err
}

func (s *SysTableService) UpdateTableRelation(ctx *gin.Context, data request.TableRelationUpdateReq) error {
	// TODO 需要完善
	return nil
}

func (s *SysTableService) DeleteTableRelationById(ctx *gin.Context, id int) error {
	err := s.sysTableRepo.ExecuteTx(ctx, func(tx *gorm.DB) error {
		if e := s.sysTableRelationRepo.DeleteById(tx, id); e != nil {
			return e
		}
		result, e := s.sysTableRelationRepo.FindById(id)
		relation := result.(model.SysTableRelation)
		if e != nil && !errors.Is(e, gorm.ErrRecordNotFound) {
			return e
		}
		// 删除缓存
		s.DeleteCache(relation.TableId)
		return nil
	})
	return err
}

func (s *SysTableService) GetTableIndexById(id int) (model.SysTableIndex, error) {
	result, err := s.sysTableIndexRepo.FindById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.SysTableIndex{}, nil
		}
		return model.SysTableIndex{}, err
	}
	data := result.(model.SysTableIndex)
	return data, nil
}

func (s *SysTableService) GetTableIndexesByTableId(tableId int) ([]model.SysTableIndex, error) {
	return s.sysTableIndexRepo.GetTableIndexesByTableId(tableId)
}

func (s *SysTableService) CreateTableIndex(ctx *gin.Context, req request.TableIndexCreateReq) error {
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
		if e := s.sysTableIndexRepo.Create(tx, &data); e != nil {
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
		if e := s.sysTableIndexFieldRepo.Create(tx, &indexFields); e != nil {
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
		if e := s.sysTableIndexFieldRepo.DeleteByField(tx, "index_id", req.Id); e != nil {
			return e
		}
		if e := s.sysTableIndexRepo.Update(tx, &req); e != nil {
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
		if e := s.sysTableIndexFieldRepo.Create(tx, &indexFields); e != nil {
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

func (s *SysTableService) DeleteTableIndexById(ctx *gin.Context, id int) error {
	result, e := s.sysTableIndexRepo.FindById(id)
	if e != nil {
		return e
	}
	index := result.(model.SysTableIndex)
	table, e := s.GetTableById(index.TableId)
	if e != nil {
		return e
	}
	err := s.sysTableRepo.ExecuteTx(ctx, func(tx *gorm.DB) error {
		if e := s.sysTableIndexRepo.DeleteById(tx, id); e != nil {
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
	indexes, e := s.sysTableIndexRepo.GetTableIndexesByTableId(id)
	if e != nil {
		return e
	}
	err := s.sysTableRepo.ExecuteTx(ctx, func(tx *gorm.DB) error {
		if e := s.sysTableIndexRepo.DeleteByField(tx, "table_id", id); e != nil {
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
	fields := convertColumnsToSysTableFields(columns)
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
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: "当前表已初始化，请勿重复操作",
		}
		return e
	}
	table = model.SysTable{
		Basic: model.Basic{
			Id: int(id),
		},
		TableName: tableCode,
		TableCode: tableCode,
		TableType: enum.System,
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
		for j, _ := range tableIndexes {
			if tableIndexes[j].ColumnName == fields[i].FieldCode {
				indexId, err := s.sf.GenerateUniqueID()
				if err != nil {
					return err
				}
				if _, exists := indexesMap[tableIndexes[j].IndexName]; !exists {
					indexesMap[tableIndexes[j].IndexName] = model.SysTableIndex{
						Basic: model.Basic{
							Id: int(indexId),
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
		if e := s.sysTableRepo.Create(tx, &table); e != nil {
			return e
		}
		if indexFields != nil && len(indexFields) > 0 {
			if e := s.sysTableIndexFieldRepo.Create(tx, &indexFields); e != nil {
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

func convertColumnsToSysTableFields(columns []model.TableColumnMate) []model.SysTableField {
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
			FieldCategory:      enum.NormalField,
			Binding:            "required", // 根据实际逻辑调整
		}
		if column.ColumnComment != "" {
			field.FieldName = column.ColumnComment
		} else {
			field.FieldName = column.ColumnName
		}
		switch column.DataType {
		case "int", "bigint":
			field.FieldType = enum.IntFieldType
		case "tinyint":
			field.FieldType = enum.TinyintFieldType
			field.FieldLength = int(column.NumericPrecision.Int64)
		case "varchar":
			field.FieldType = enum.VarcharFieldType
			field.FieldLength = int(column.CharacterMaximumLength.Int64)
		case "text", "mediumtext", "longtext":
			field.FieldType = enum.TextFieldType
			field.FieldLength = int(column.CharacterMaximumLength.Int64)
		case "boolean", "bool":
			field.FieldType = enum.BooleanFieldType
		case "date":
			field.FieldType = enum.DateFieldType
		case "datetime", "timestamp":
			field.FieldType = enum.DatetimeFieldType
		case "time":
			field.FieldType = enum.TimeFieldType
		default:
			field.FieldType = enum.VarcharFieldType
			field.FieldLength = int(column.NumericPrecision.Int64)
		}
		// 检查DefaultValue是否有值
		if column.ColumnDefault.Valid {
			field.DefaultValue = &column.ColumnDefault.String
		}
		fields = append(fields, field)
	}
	return fields
}
