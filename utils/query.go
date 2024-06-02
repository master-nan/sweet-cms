/**
 * @Author: Nan
 * @Date: 2024/5/29 上午11:44
 */

package utils

import (
	"gorm.io/gorm"
	"sweet-cms/enum"
	"sweet-cms/form/request"
	"time"
)

func parseValue(value interface{}, valueType enum.SysTableFieldType) interface{} {
	switch valueType {
	case enum.INT:
		return value.(int)
	case enum.FLOAT:
		return value.(float64)
	case enum.VARCHAR:
		return value.(string)
	case enum.BOOLEAN:
		return value.(bool)
	case enum.TEXT:
		return value.(string)
	case enum.DATE:
		t, _ := time.Parse(time.DateOnly, value.(string))
		return t
	case enum.DATETIME:
		t, _ := time.Parse(time.DateTime, value.(string))
		return t
	case enum.TIME:
		t, _ := time.Parse(time.TimeOnly, value.(string))
		return t
	default:
		return value
	}
}

func BuildQuery(db *gorm.DB, basic request.Basic) *gorm.DB {
	query := db
	// 构建查询条件
	for _, exprGroup := range basic.Expressions {
		var subQuery *gorm.DB
		for _, rule := range exprGroup.Rules {
			value := parseValue(rule.Value, rule.Type)
			switch rule.ExpressionType {
			case enum.GT:
				subQuery = query.Where(rule.Field+" > ?", value)
			case enum.LT:
				subQuery = query.Where(rule.Field+" < ?", value)
			case enum.GTE:
				subQuery = query.Where(rule.Field+" >= ?", value)
			case enum.LTE:
				subQuery = query.Where(rule.Field+" <= ?", value)
			case enum.EQ:
				subQuery = query.Where(rule.Field+" = ?", value)
			case enum.NE:
				subQuery = query.Where(rule.Field+" != ?", value)
			case enum.LIKE:
				subQuery = query.Where(rule.Field+" LIKE %?%", value)
			case enum.NOT_LIKE:
				subQuery = query.Where(rule.Field+" NOT LIKE %?%", value)
			case enum.IN:
				subQuery = query.Where(rule.Field+" IN (?)", value)
			case enum.NOT_IN:
				subQuery = query.Where(rule.Field+" NOT IN (?)", value)
			case enum.IS_NULL:
				subQuery = query.Where(rule.Field + " IS NULL")
			case enum.IS_NOT_NULL:
				subQuery = query.Where(rule.Field + " IS NOT NULL")
			default:
				continue
			}
		}

		// 处理嵌套表达式
		for _, nestedExpr := range exprGroup.Nested {
			nestedQuery := BuildQuery(db, request.Basic{Expressions: []request.ExpressionGroup{nestedExpr}}) // 递归处理嵌套表达式
			switch exprGroup.Logic {
			case enum.OR:
				if subQuery == nil {
					subQuery = nestedQuery
				} else {
					subQuery = subQuery.Or(nestedQuery)
				}
			case enum.AND:
				if subQuery == nil {
					subQuery = nestedQuery
				} else {
					subQuery = subQuery.Where(nestedQuery)
				}
			}
		}

		// 应用当前表达式组的逻辑
		if subQuery != nil {
			switch exprGroup.Logic {
			case enum.AND:
				query = query.Where(subQuery)
			case enum.OR:
				query = query.Or(subQuery)
			default:
				continue
			}
		}
	}

	// 添加排序
	if basic.Order.Field != "" {
		order := basic.Order.Field
		if !basic.Order.IsAsc {
			order += " desc"
		}
		query = query.Order(order)
	}

	// 设置 Page 和 Num 的默认值
	if basic.Page <= 0 {
		basic.Page = 1 // 默认页码为 1
	}
	if basic.Num <= 0 {
		basic.Num = 10 // 默认每页数量为 10
	}
	// 设置 Num 的上限
	const maxNum = 5000
	if basic.Num > maxNum {
		basic.Num = maxNum
	}
	// 添加分页
	if basic.Page > 0 && basic.Num > 0 {
		query = query.Limit(basic.Num).Offset((basic.Page - 1) * basic.Num)
	}

	return query
}
