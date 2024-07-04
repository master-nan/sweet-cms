/**
 * @Author: Nan
 * @Date: 2024/5/29 上午11:44
 */

package utils

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
	"reflect"
	"strings"
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

func ExecuteQuery(db *gorm.DB, basic request.Basic) *gorm.DB {
	// 构建基本查询
	query := buildQuery(db, basic)
	// 应用排序和分页
	query = finalizeQuery(query, basic)

	return query
}

func buildQuery(db *gorm.DB, basic request.Basic) *gorm.DB {
	query := db
	if !basic.IncludeDeleted {
		query = query.Where("gmt_delete IS NULL")
	}
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
			nestedQuery := buildQuery(db, request.Basic{Expressions: []request.ExpressionGroup{nestedExpr}}) // 递归处理嵌套表达式
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
	return query
}

func finalizeQuery(query *gorm.DB, basic request.Basic) *gorm.DB {
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

// DynamicQuery 动态生成结构体并进行查询
func DynamicQuery(db *gorm.DB, basic request.Basic, table model.SysTable) (repository.GeneralizationListResult, error) {
	var result repository.GeneralizationListResult
	// 创建动态结构体
	modelType := CreateDynamicStruct(table.TableFields)

	// 构建查询
	query := ExecuteQuery(db.Table(table.TableCode), basic)

	// 查询结果
	results := reflect.New(reflect.SliceOf(modelType)).Elem()
	err := query.Find(results.Addr().Interface()).Error
	if err != nil {
		return result, err
	}
	// 总数查询
	var total int64
	err = query.Limit(-1).Offset(-1).Count(&total).Error
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
	result.Data = records
	result.Total = int(total)
	return result, nil
}

// CreateDynamicStruct 根据表元数据创建动态结构体
func CreateDynamicStruct(fields []model.SysTableField) reflect.Type {
	var fieldsList []reflect.StructField
	for _, field := range fields {
		fieldType := getFieldType(field.FieldType)
		fieldTag := buildTag(field)
		fieldsList = append(fieldsList, reflect.StructField{
			Name: toCamelCaseGo(field.FieldCode),
			Type: fieldType,
			Tag:  reflect.StructTag(fieldTag),
		})
	}
	return reflect.StructOf(fieldsList)
}

// getFieldType 获取对应类型
func getFieldType(fieldType enum.SysTableFieldType) reflect.Type {
	switch fieldType {
	case enum.INT:
		return reflect.TypeOf(int(0)) // 或 reflect.TypeOf(int64(0)) 根据需要选择
	case enum.FLOAT:
		return reflect.TypeOf(float64(0.0)) // 使用 float64 是 Go 中最常用的浮点类型
	case enum.VARCHAR, enum.TEXT:
		return reflect.TypeOf("") // 字符串类型
	case enum.BOOLEAN:
		return reflect.TypeOf(false)
	case enum.DATE, enum.DATETIME, enum.TIME:
		return reflect.TypeOf(time.Time{}) // 对于所有的时间类型使用 time.Time
	default:
		return reflect.TypeOf(nil) // 未知类型返回 nil 类型，可能需要处理错误
	}
}

// buildTag 构建结构体tag
func buildTag(field model.SysTableField) string {
	gormParts := []string{
		fmt.Sprintf(`column:%s`, field.FieldCode),
		fmt.Sprintf(`type:%s`, getSQLType(field.FieldType, field.FieldLength, field.FieldDecimalLength)),
	}
	if field.FieldLength > 0 {
		gormParts = append(gormParts, fmt.Sprintf(`size:%d`, field.FieldLength))
	}
	if field.DefaultValue != nil && *field.DefaultValue != "" {
		gormParts = append(gormParts, fmt.Sprintf(`default:'%s'`, *field.DefaultValue))
	}
	if field.IsPrimaryKey {
		gormParts = append(gormParts, `primaryKey:true`)
	}
	if !field.IsNull {
		gormParts = append(gormParts, `notNull:true`)
	}
	//if field.IsIndex {
	//	gormParts = append(gormParts, `index:true`)
	//}
	gormParts = append(gormParts, fmt.Sprintf(`comment:'%s'`, field.FieldName))

	// JSON 标签
	jsonPart := fmt.Sprintf(`json:"%s"`, toCamelCaseJson(field.FieldCode))

	// Binding 标签，如果字段定义了 Binding 规则，使用该规则
	bindingPart := ""
	if field.Binding != "" {
		bindingPart = fmt.Sprintf(`binding:"%s"`, field.Binding)
	}
	// 组合 GORM, JSON 和 Binding 标签
	fullTag := fmt.Sprintf(`gorm:"%s" %s %s`, strings.Join(gormParts, ";"), jsonPart, bindingPart)
	return fullTag
}

// getSQLType 返回类型和长度
func getSQLType(fieldType enum.SysTableFieldType, length int, decimal int) string {
	switch fieldType {
	case enum.INT:
		return "int"
	case enum.FLOAT:
		if decimal > 0 {
			return fmt.Sprintf("decimal(%d,%d)", length, decimal)
		}
		return "float"
	case enum.VARCHAR:
		return fmt.Sprintf("varchar(%d)", length)
	case enum.TEXT:
		return "text"
	case enum.BOOLEAN:
		return "boolean"
	case enum.DATE:
		return "date"
	case enum.DATETIME:
		return "datetime"
	case enum.TIME:
		return "time"
	default:
		return "text"
	}
}

func GetTableName(db *gorm.DB, tableCode string) string {
	tableName := db.NamingStrategy.TableName(tableCode)
	return tableName
}

func toCamelCaseGo(input string) string {
	parts := strings.Split(input, "_")
	c := cases.Title(language.English) // 使用英语规则进行标题转换
	for i, part := range parts {
		parts[i] = c.String(part)
	}
	return strings.Join(parts, "")
}

func toCamelCaseJson(input string) string {
	parts := strings.Split(input, "_")
	c := cases.Title(language.English) // 使用英语规则进行标题转换
	for i, part := range parts {
		if i == 0 {
			// 第一个单词首字母小写
			parts[i] = strings.ToLower(part)
		} else {
			// 其余单词首字母大写
			parts[i] = c.String(part)
		}
	}
	return strings.Join(parts, "")
}
