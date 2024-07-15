package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"sweet-cms/enum"
	"sweet-cms/form/response"
	"sweet-cms/model"
	"time"

	"github.com/gin-gonic/gin"
	"reflect"
)

// ConvertStruct 通用结构体转换函数
func ConvertStruct(source, target interface{}) error {
	srcVal := reflect.ValueOf(source)
	destVal := reflect.ValueOf(target).Elem() // 获取指针指向的元素

	if srcVal.Kind() == reflect.Slice && destVal.Kind() == reflect.Slice {
		// 处理切片类型
		itemType := destVal.Type().Elem()
		resultSlice := reflect.MakeSlice(destVal.Type(), srcVal.Len(), srcVal.Len())

		for i := 0; i < srcVal.Len(); i++ {
			newItem := reflect.New(itemType.Elem()) // 创建新的元素实例
			err := ConvertStruct(srcVal.Index(i).Interface(), newItem.Interface())
			if err != nil {
				return err
			}
			resultSlice.Index(i).Set(newItem.Elem())
		}
		destVal.Set(resultSlice)
	} else {
		// 处理单个对象
		srcType := srcVal.Type()

		for i := 0; i < srcType.NumField(); i++ {
			srcField := srcType.Field(i)
			srcFieldValue := srcVal.Field(i)

			if destField := destVal.FieldByName(srcField.Name); destField.IsValid() && destField.CanSet() {
				// 直接设置值
				destField.Set(srcFieldValue)
			}
		}
	}
	return nil
}

func Assignment(source interface{}, target interface{}) {
	s1 := reflect.ValueOf(source).Elem()
	t1 := reflect.ValueOf(target).Elem()
	for i := 0; i < s1.NumField(); i++ {
		name := s1.Type().Field(i).Name
		if ok := t1.FieldByName(name).IsValid(); ok {
			value := s1.FieldByName(name).Interface()
			if reflect.ValueOf(value).Kind() == reflect.Ptr {
				if !reflect.ValueOf(value).IsNil() {
					t1.FieldByName(name).Set(reflect.ValueOf(value).Elem())
				}
			} else {
				t1.FieldByName(name).Set(reflect.ValueOf(value))
			}
		}
	}
}

func GetStructSqlSelect(data interface{}) []string {
	var arr []string
	refData := reflect.ValueOf(data)
	for i := 0; i < refData.NumField(); i++ {
		tag := refData.Type().Field(i).Tag
		value := refData.Field(i).Interface() //获取属性的值
		refValue := reflect.ValueOf(value)    //获取值的类型
		if refValue.Kind() == reflect.Ptr {   //是否为指针  判断是否为nil
			if !refValue.IsNil() {
				arr = append(arr, tag.Get("json"))
			}
		} else {
			if value != "" {
				arr = append(arr, tag.Get("json"))
			}
		}
	}
	return arr
}

func IsStructEmpty(source interface{}, target interface{}) bool {
	return reflect.DeepEqual(source, target)
}

func SaveSession(ctx *gin.Context, key string, value interface{}) {
	session := sessions.Default(ctx)
	option := sessions.Options{Path: "/", MaxAge: 3600}
	session.Options(option)
	session.Set(key, value)
	session.Save()
}

func DeleteSession(ctx *gin.Context, key string) {
	session := sessions.Default(ctx)
	session.Delete(key)
	session.Save()
}

func GetSessionString(ctx *gin.Context, key string) string {
	session := sessions.Default(ctx)
	value := session.Get(key)
	if value == nil {
		return ""
	}
	return value.(string)
}

func ClearSession(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
}

func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func Encryption(password string, salt string) string {
	str := fmt.Sprintf("%s%s", password, salt)
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func IsEmpty(s interface{}) bool {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return true
	}
	return isZero(v)
}

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Func, reflect.Map, reflect.Slice, reflect.Interface, reflect.Ptr, reflect.Chan:
		return v.IsNil()
	case reflect.Array, reflect.String:
		return v.Len() == 0
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if !isZero(v.Field(i)) {
				return false
			}
		}
		return true
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Complex64, reflect.Complex128:
		return v.Complex() == 0
	default:
		// 未处理的类型
		return false
	}
}

func TranslateError(err validator.FieldError) string {
	// 你可以根据err.Tag()来判断是哪种验证错误，然后返回不同的错误信息
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s 是必填项", err.Field())
	default:
		return fmt.Sprintf("%s 验证错误", err.Field())
	}
}

// 辅助函数用于创建各种类型的指针
func BoolPtr(b bool) *bool {
	return &b
}

func IntPtr(i int) *int {
	return &i
}

func StringPtr(s string) *string {
	return &s
}

func SqlTypeFromFieldType(fieldType enum.SysTableFieldType) string {
	switch fieldType {
	case enum.IntFieldType:
		return "IntFieldType"
	case enum.VarcharFieldType:
		return "VarcharFieldType" // 长度将在外部指定
	case enum.DatetimeFieldType:
		return "DatetimeFieldType"
	case enum.BooleanFieldType:
		return "BooleanFieldType"
	case enum.TextFieldType:
		return "TextFieldType"
	case enum.DateFieldType:
		return "DateFieldType"
	case enum.TimeFieldType:
		return "TimeFieldType"
	default:
		return "TextFieldType"
	}
}

// UpdateAccessTokens 替换当前token字符串
// existingTokens string: 现有token字符串
// newToken string: 新token
// 返回更新后的token字符串
func UpdateAccessTokens(existingTokens string, newToken string) string {
	// 分隔符
	delimiter := ","
	var tokens []string
	// 检查是否有现有的Tokens，防止创建一个包含空字符串的slice
	if existingTokens != "" {
		tokens = strings.Split(existingTokens, delimiter)
	}

	// 确保只保留最近的4个token（因为我们将添加一个新的）
	if len(tokens) >= 5 {
		tokens = tokens[1:] // 删除最老的Token
	}
	// 添加新的Token
	tokens = append(tokens, newToken)
	// 将更新后的Token列表连接成一个新的字符串
	updatedTokens := strings.Join(tokens, delimiter)
	return updatedTokens
}

// ContainsToken 查找token是否在当前token集合内
func ContainsToken(existingTokens string, newToken string) bool {
	// 如果现有的tokens字符串为空，直接返回false
	if existingTokens == "" {
		return false
	}
	// 分隔符
	delimiter := ","
	// 分割现有tokens
	tokens := strings.Split(existingTokens, delimiter)

	// 检查newToken是否在tokens切片中
	for _, token := range tokens {
		if token == newToken {
			return true // 找到了，返回true
		}
	}
	return false // 没有找到，返回false
}

func ConvertColumnsToSysTableFields(columns []model.TableColumnMate) ([]model.SysTableField, error) {
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
	return fields, nil
}

func ValidatorBody[T any](ctx *gin.Context, data *T, translator ut.Translator) error {
	err := ctx.ShouldBindBodyWith(data, binding.JSON)
	if err != nil {
		if err == io.EOF {
			// 客户端请求体为空
			e := &response.AdminError{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: "请求参数数据",
			}
			return e
		}
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			// 如果是验证错误，则翻译错误信息
			var errorMessages []string
			for _, e := range ve {
				errMsg := e.Translate(translator)
				errorMessages = append(errorMessages, errMsg)
			}
			e := &response.AdminError{
				ErrorCode:    http.StatusBadRequest,
				ErrorMessage: strings.Join(errorMessages, ","),
			}
			return e
		}
		return err
	}
	return nil
}
