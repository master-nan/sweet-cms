/**
 * @Author: Nan
 * @Date: 2024/5/17 下午3:38
 */

package request

import (
	"sweet-cms/enum"
)

type Basic struct {
	Page        int               `json:"page"`
	Num         int               `json:"num"`
	Order       Order             `json:"order"`
	TableCode   string            `json:"table_code"`
	Expressions []ExpressionGroup `json:"expressions"`
	QuickQuery  *QuickQuery       `json:"quick_query"`
}

type ExpressionGroup struct {
	Logic  enum.ExpressionLogic `json:"logic"`  // "and" 或 "or"
	Rules  []QueryRule          `json:"rules"`  // 基础查询规则
	Nested []ExpressionGroup    `json:"nested"` // 嵌套的表达式组
}

type QueryRule struct {
	Field          string                 `json:"field"`
	ExpressionType enum.ExpressionType    `json:"expression_type"` // 比较器类型，如EQ, LT等
	Value          interface{}            `json:"value"`
	Type           enum.SysTableFieldType `json:"type"` // 字段类型
}

type Order struct {
	Field string `json:"field"`
	IsAsc bool   `json:"is_asc"`
}

type QuickQuery struct {
	KeyWord string `json:"keyword"`
}
