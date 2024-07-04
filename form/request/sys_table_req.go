/**
 * @Author: Nan
 * @Date: 2024/5/17 上午11:39
 */

package request

import (
	"sweet-cms/enum"
)

type TableCreateReq struct {
	TableName string            `json:"tableName" binding:"required"`
	TableCode string            `json:"tableCode" binding:"required"`
	TableType enum.SysTableType `json:"tableType" binding:"required"`
}

type TableUpdateReq struct {
	Id        int    `json:"id" binding:"required"`
	TableName string `json:"tableName" binding:"required"`
}

type TableFieldCreateReq struct {
	TableId            int                         `json:"tableId" binding:"required"`
	FieldName          string                      `json:"fieldName" binding:"required"`        // 列名
	FieldCode          string                      `json:"fieldCode" binding:"required"`        // 数据库中字段名
	FieldType          enum.SysTableFieldType      `json:"type" binding:"required"`             // 字段类型
	FieldLength        int                         `json:"fieldLength" binding:"required"`      // 字段长度
	FieldDecimalLength int                         `json:"fieldDecimalLength"`                  // 小数位数
	InputType          enum.SysTableFieldInputType `json:"inputType" binding:"required"`        // 输入类型
	DefaultValue       string                      `json:"defaultValue"`                        // 默认值
	DictCode           string                      `json:"dictCode"`                            // 所用字典
	IsPrimaryKey       bool                        `json:"isPrimaryKey" binding:"required"`     // 是否主键
	IsQuickSearch      bool                        `json:"IsQuickSearch" binding:"required"`    // 是否快捷搜索
	IsAdvancedSearch   bool                        `json:"isAdvancedSearch" binding:"required"` // 是否高级搜索
	IsSort             bool                        `json:"isSort" binding:"required"`           // 是否可排序
	IsNull             bool                        `json:"isNull" binding:"required"`           // 是否可空
	OriginalFieldId    int                         `json:"originalFieldId"`                     // 原字段Id
}

type TableFieldUpdateReq struct {
	Id                 int                         `json:"id" binding:"required"`
	TableId            int                         `json:"table_id" binding:"required"`
	FieldName          string                      `json:"field_name" binding:"required"`         // 列名
	FieldCode          string                      `json:"field_code" binding:"required"`         // 数据库中字段名
	FieldType          enum.SysTableFieldType      `json:"type" binding:"required"`               // 字段类型
	FieldLength        int                         `json:"field_length" binding:"required"`       // 字段长度
	FieldDecimalLength int                         `json:"field_decimal_length"`                  // 小数位数
	InputType          enum.SysTableFieldInputType `json:"input_type" binding:"required"`         // 输入类型
	DefaultValue       string                      `json:"default_value"`                         // 默认值
	DictCode           string                      `json:"dict_code"`                             // 所用字典
	IsQuickSearch      bool                        `json:"is_quick_search" binding:"required"`    // 是否快捷搜索
	IsAdvancedSearch   bool                        `json:"is_advanced_search" binding:"required"` // 是否高级搜索
	IsSort             bool                        `json:"is_sort" binding:"required"`            // 是否可排序
	IsNull             bool                        `json:"is_null" binding:"required"`            // 是否可空
	OriginalFieldId    int                         `json:"original_field_id"`                     // 原字段Id
}

type TableRelationCreateReq struct {
	TableId        int                       `json:"table_id" binding:"required"`
	RelatedTableId int                       `json:"related_table_id" binding:"required"` // 关联的表的Id
	ReferenceKey   string                    `json:"reference_key" binding:"required"`    // 主表对应字段
	ForeignKey     string                    `json:"foreign_key" binding:"required"`      // 关联表 字段
	RelationType   enum.SysTableRelationType `json:"relation_type" binding:"required"`
}

type TableRelationUpdateReq struct {
	Id             int                       `json:"id" binding:"required"`
	TableId        int                       `json:"table_id" binding:"required"`
	RelatedTableId int                       `json:"related_table_id" binding:"required"` // 关联的表的Id
	ReferenceKey   string                    `json:"reference_key" binding:"required"`    // 主表对应字段
	ForeignKey     string                    `json:"foreign_key" binding:"required"`      // 关联表 字段
	RelationType   enum.SysTableRelationType `json:"relation_type" binding:"required"`
}

type TableIndexFieldReq struct {
	TableId   int    `json:"table_id" binding:"required"`
	FieldId   int    `json:"field_id" binding:"required"`
	FieldCode string `json:"field_code"  binding:"required"`
}

type TableIndexCreateReq struct {
	TableId     int                  `json:"table_id" binding:"required"`
	IndexName   string               `json:"index_name" binding:"required"`
	IsUnique    bool                 `json:"is_unique" binding:"required"`
	IndexFields []TableIndexFieldReq `json:"index_fields" binding:"required,min=1"`
}

type TableIndexUpdateReq struct {
	Id          int                  `json:"id" binding:"required"`
	TableId     int                  `json:"table_id" binding:"required"`
	IndexName   string               `json:"index_name" binding:"required"`
	IsUnique    bool                 `json:"is_unique" binding:"required"`
	IndexFields []TableIndexFieldReq `json:"index_fields" binding:"required"`
}
