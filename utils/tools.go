package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/go-playground/validator/v10"
	"math/rand"
	"sweet-cms/enum"
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
	option := sessions.Options{MaxAge: 3600}
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
	case enum.INT:
		return "INT"
	case enum.VARCHAR:
		return "VARCHAR" // 长度将在外部指定
	case enum.DATETIME:
		return "DATETIME"
	case enum.BOOLEAN:
		return "BOOLEAN"
	case enum.TEXT:
		return "TEXT"
	case enum.DATE:
		return "DATE"
	case enum.TIME:
		return "TIME"
	default:
		return "TEXT"
	}
}
