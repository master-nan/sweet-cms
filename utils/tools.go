package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/go-playground/validator/v10"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"reflect"
)

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
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if !val.IsValid() {
		return true
	}
	for i := 0; i < val.NumField(); i++ {
		if !isZero(val.Field(i)) {
			return false
		}
	}
	return true
}

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		return v.Len() == 0
	case reflect.Map, reflect.Ptr, reflect.Interface, reflect.Chan, reflect.Func:
		return v.IsNil()
	}
	zero := reflect.Zero(v.Type()).Interface()
	current := v.Interface()
	return reflect.DeepEqual(current, zero)
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
