/**
 * @Author: Nan
 * @Date: 2024/5/17 上午11:39
 */

package request

import (
	"sweet-cms/enum"
)

type TableCreateReq struct {
	TableName string            `json:"table_name" binding:"required"`
	TableCode string            `json:"table_code" binding:"required"`
	TableType enum.SysTableType `json:"table_type" binding:"required"`
}

type TableUpdateReq struct {
	ID        int    `json:"id" binding:"required"`
	TableName string `json:"table_name" binding:"required"`
}

type TableFieldCreateReq struct {
	TableID            int                         `json:"table_id" binding:"required"`
	FieldName          string                      `json:"field_name" binding:"required"`         // 列名
	FieldCode          string                      `json:"field_code" binding:"required"`         // 数据库中字段名
	FieldType          enum.SysTableFieldType      `json:"type" binding:"required"`               // 字段类型
	FieldLength        int                         `json:"field_length" binding:"required"`       // 字段长度
	FieldDecimalLength int                         `json:"field_decimal_length"`                  // 小数位数
	InputType          enum.SysTableFieldInputType `json:"input_type" binding:"required"`         // 输入类型
	DefaultValue       string                      `json:"default_value"`                         // 默认值
	DictCode           string                      `json:"dict_code"`                             // 所用字典
	IsPrimaryKey       bool                        `json:"is_primary_key" binding:"required"`     // 是否主键
	IsQuickSearch      bool                        `json:"is_quick_search" binding:"required"`    // 是否快捷搜索
	IsAdvancedSearch   bool                        `json:"is_advanced_search" binding:"required"` // 是否高级搜索
	IsSort             bool                        `json:"is_sort" binding:"required"`            // 是否可排序
	IsNull             bool                        `json:"is_null" binding:"required"`            // 是否可空
	OriginalFieldID    int                         `json:"original_field_id"`                     // 原字段ID
}

type TableFieldUpdateReq struct {
	ID                 int                         `json:"id" binding:"required"`
	TableID            int                         `json:"table_id" binding:"required"`
	FieldName          string                      `json:"field_name" binding:"required"`         // 列名
	FieldCode          string                      `json:"field_code" binding:"required"`         // 数据库中字段名
	FieldType          enum.SysTableFieldType      `json:"type" binding:"required"`               // 字段类型
	FieldLength        int                         `json:"field_length" binding:"required"`       // 字段长度
	FieldDecimalLength int                         `json:"field_decimal_length"`                  // 小数位数
	InputType          enum.SysTableFieldInputType `json:"input_type" binding:"required"`         // 输入类型
	DefaultValue       string                      `json:"default_value"`                         // 默认值
	DictCode           string                      `json:"dict_code"`                             // 所用字典
	IsPrimaryKey       bool                        `json:"is_primary_key" binding:"required"`     // 是否主键
	IsQuickSearch      bool                        `json:"is_quick_search" binding:"required"`    // 是否快捷搜索
	IsAdvancedSearch   bool                        `json:"is_advanced_search" binding:"required"` // 是否高级搜索
	IsSort             bool                        `json:"is_sort" binding:"required"`            // 是否可排序
	IsNull             bool                        `json:"is_null" binding:"required"`            // 是否可空
	OriginalFieldID    int                         `json:"original_field_id"`                     // 原字段ID
}

type TableRelationCreateReq struct {
	TableID        int                       `json:"table_id" binding:"required"`
	RelatedTableID int                       `json:"related_table_id" binding:"required"` // 关联的表的ID
	ReferenceKey   string                    `json:"reference_key" binding:"required"`    // 主表对应字段
	ForeignKey     string                    `json:"foreign_key" binding:"required"`      // 关联表 字段
	RelationType   enum.SysTableRelationType `json:"relation_type" binding:"required"`
}

type TableRelationUpdateReq struct {
	ID             int                       `json:"id" binding:"required"`
	TableID        int                       `json:"table_id" binding:"required"`
	RelatedTableID int                       `json:"related_table_id" binding:"required"` // 关联的表的ID
	ReferenceKey   string                    `json:"reference_key" binding:"required"`    // 主表对应字段
	ForeignKey     string                    `json:"foreign_key" binding:"required"`      // 关联表 字段
	RelationType   enum.SysTableRelationType `json:"relation_type" binding:"required"`
}
