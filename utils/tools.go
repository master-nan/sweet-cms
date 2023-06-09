package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-contrib/sessions"
	"math/rand"
	"sweet-cms/model"
	"time"

	//"github.com/gin-contrib/sessions"
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

func GetMenu() []model.MenuTree {
	var menuList []model.MenuTree
	menu1 := model.MenuTree{
		SysMenu: model.SysMenu{
			PID:    0,
			Name:   "index",
			Path:   "/sweet/admin/index",
			Title:  "首页",
			Option: "",
			Icon:   "speedometer2",
		},
		Children: nil,
	}
	menu2 := model.MenuTree{
		SysMenu: model.SysMenu{
			PID:    0,
			Name:   "#",
			Path:   "#",
			Title:  "系统设置",
			Option: "",
			Icon:   "gear",
		},
		Children: []model.MenuTree{
			{
				SysMenu: model.SysMenu{
					PID:    0,
					Name:   "#",
					Path:   "#",
					Title:  "角色管理",
					Option: "",
					Icon:   "gear",
				},
				Children: nil,
			},
			{
				SysMenu: model.SysMenu{
					PID:    0,
					Name:   "#",
					Path:   "#",
					Title:  "菜单管理",
					Option: "",
					Icon:   "gear",
				},
				Children: nil,
			},
			{
				SysMenu: model.SysMenu{
					PID:    0,
					Name:   "#",
					Path:   "#",
					Title:  "用户管理",
					Option: "",
					Icon:   "gear",
				},
				Children: nil,
			},
		},
	}
	menu3 := model.MenuTree{
		SysMenu: model.SysMenu{
			PID:    0,
			Name:   "article",
			Path:   "/sweet/admin/article",
			Title:  "内容设置",
			Option: "",
			Icon:   "journal-album",
		},
		Children: []model.MenuTree{
			{
				SysMenu: model.SysMenu{
					PID:    0,
					Name:   "#",
					Path:   "channel",
					Title:  "分类管理",
					Option: "",
					Icon:   "",
				},
				Children: nil,
			},
			{
				SysMenu: model.SysMenu{
					PID:    0,
					Name:   "index",
					Path:   "index",
					Title:  "内容管理",
					Option: "",
					Icon:   "",
				},
				Children: nil,
			},
		},
	}
	menuList = append(menuList, menu1, menu2, menu3)
	return menuList
}
