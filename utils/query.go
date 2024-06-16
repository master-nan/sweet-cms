/**
 * @Author: Nan
 * @Date: 2024/5/29 上午11:44
 */

package utils

import (
	"fmt"
	"gorm.io/gorm"
	"reflect"
	"sweet-cms/enum"
	"sweet-cms/form/request"
	"sweet-cms/model"
	"sweet-cms/repository"
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

// 动态生成结构体并进行查询
func DynamicQuery(db *gorm.DB, basic request.Basic, table model.SysTable) (repository.GeneralizationListResult, error) {
	var result repository.GeneralizationListResult
	// 创建动态结构体
	modelType := createDynamicStruct(table.TableFields)

	// 构建查询
	query := BuildQuery(db.Table(table.TableCode), basic)

	// 查询结果
	results := reflect.New(reflect.SliceOf(modelType)).Elem()
	err := query.Find(results.Addr().Interface()).Error
	if err != nil {
		return result, err
	}
	// 转换结果为更通用的格式
	records := make([]map[string]interface{}, results.Len())
	for i := 0; i < results.Len(); i++ {
		record := make(map[string]interface{})
		val := results.Index(i)
		for _, field := range table.TableFields {
			fieldValue := val.FieldByName(field.FieldCode)
			if fieldValue.IsValid() {
				record[field.FieldCode] = fieldValue.Interface()
			}
		}
		records[i] = record
	}
	// 总数查询
	var total int64
	db.Table(table.TableCode).Count(&total)
	result.Data = records
	result.Total = int(total)
	return result, nil
}

// 根据表元数据创建动态结构体
func createDynamicStruct(fields []model.SysTableField) reflect.Type {
	var fieldsList []reflect.StructField
	for _, field := range fields {
		var fieldType reflect.Type
		switch field.FieldType {
		case enum.INT:
			fieldType = reflect.TypeOf(0)
		case enum.FLOAT:
			fieldType = reflect.TypeOf(float64(0.0))
		case enum.VARCHAR:
			fieldType = reflect.TypeOf("")
		case enum.TEXT:
			fieldType = reflect.TypeOf("")
		case enum.BOOLEAN:
			fieldType = reflect.TypeOf(false)
		case enum.DATE:
			fieldType = reflect.TypeOf(time.Time{})
		case enum.DATETIME:
			fieldType = reflect.TypeOf(time.Time{})
		case enum.TIME:
			fieldType = reflect.TypeOf(time.Time{})
		}
		fieldsList = append(fieldsList, reflect.StructField{
			Name: field.FieldName,
			Type: fieldType,
			Tag:  reflect.StructTag(fmt.Sprintf(`gorm:"column:%s" json:"%s"`, field.FieldCode, field.FieldName)),
		})
	}
	return reflect.StructOf(fieldsList)
}
