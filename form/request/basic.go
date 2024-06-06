/**
 * @Author: Nan
 * @Date: 2024/5/17 下午3:38
 */

package request

import (
	"sweet-cms/enum"
)

// Basic 请求参数参数
type Basic struct {
	Page        int               `json:"page" example:"1"`
	Num         int               `json:"num" example:"10"`
	Order       Order             `json:"order" example:"\{\"field\":\"name\",\"is_asc\":true\}"`
	TableCode   string            `json:"table_code" example:"sys_dict"`
	Expressions []ExpressionGroup `json:"expressions"`
	QuickQuery  *QuickQuery       `json:"quick_query" example:"\{'keyword':'search'\}"`
}

// ExpressionGroup 参数请求组
type ExpressionGroup struct {
	Logic  enum.ExpressionLogic `json:"logic"`  // "and" 或 "or"
	Rules  []QueryRule          `json:"rules"`  // 基础查询规则
	Nested []ExpressionGroup    `json:"nested"` // 嵌套的表达式组
}

// QueryRule 查询规则
type QueryRule struct {
	Field          string                 `json:"field"`           // 字段
	ExpressionType enum.ExpressionType    `json:"expression_type"` // 比较器类型，如EQ, LT等
	Value          interface{}            `json:"value"`           // 值
	Type           enum.SysTableFieldType `json:"type"`            // 字段类型
}

// Order 排序
type Order struct {
	Field string `json:"field"`
	IsAsc bool   `json:"is_asc"`
}

// QuickQuery 快速查询参数
type QuickQuery struct {
	KeyWord string `json:"keyword"`
}
