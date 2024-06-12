package model

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/pkg/errors"
	"sweet-cms/enum"
)

type SysConfigure struct {
	Basic
	EnableCaptcha bool `gorm:"comment:登录验证码" json:"enable_captcha"`
}

type SysMenu struct {
	Basic
	PID       int    `gorm:"type:int" json:"pid"`
	Name      string `gorm:"size:32;comment:路由" json:"name"`
	Path      string `gorm:"size:128;comment:路径" json:"path"`
	Component string `gorm:"size:64;comment:路由主体" json:"component"`
	Title     string `gorm:"size:64;comment:显示标题" json:"title"`
	IsHidden  bool   `gorm:"comment:是否隐藏" json:"isHidden"`
	Sequence  uint8  `gorm:"comment:排序" gorm:"type:tinyint" json:"sequence"`
	Option    string `gorm:"size:64;comment:排序" json:"option"`
	Icon      string `gorm:"size:32;comment:图标" json:"icon"`
	Redirect  string `gorm:"size:128;comment:重定向地址" json:"redirect"`
	IsUnfold  bool   `gorm:"comment:默认展开" json:"isUnfold"`
}

type SysMenuBtn struct {
	Basic
	Name     string                  `gorm:"size:128;comment:按钮名称" json:"name"`
	Code     string                  `gorm:"size:128;comment:按钮编码" json:"code"`
	Memo     string                  `gorm:"size:128;comment:备注" json:"memo"`
	Position enum.SysMenuBtnPosition `gorm:"size:128;comment:位置" json:"position"`
}

type SysRole struct {
	Basic
	Name string `json:"name"`
	Rs   string `json:"rs"`
	Memo string `json:"memo"`
}

type SysUser struct {
	Basic
	UserName string `json:"username"`
	RoleId   string `json:"roleId"`
	Password string `json:"-"`
}

type SysTable struct {
	Basic
	TableName   string            `gorm:"size:128;comment:表名" json:"table_name"`
	TableCode   string            `gorm:"size:128;comment:数据库中表名" json:"table_code"`
	TableType   enum.SysTableType `gorm:"default:1" json:"table_type"`
	ParentID    int               `gorm:"type:int" json:"parent_id"`
	TableFields []SysTableField   `gorm:"-" json:"table_fields"`
}

func (st *SysTable) Value() (driver.Value, error) {
	return json.Marshal(st.TableFields)
}

func (st *SysTable) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Scan source is not []byte")
	}
	return json.Unmarshal(bytes, &st.TableFields)
}

type SysTableField struct {
	Basic
	TableID            int                         `gorm:"comment:table_id" json:"table_id" binding:"required"`
	FieldName          string                      `gorm:"size:128;comment:列名" json:"field_name"`
	FieldCode          string                      `gorm:"size:128;comment:数据库中字段名" json:"field_code"`
	FieldType          enum.SysTableFieldType      `gorm:"default:1;comment:字段类型" json:"type"`
	FieldLength        int                         `gorm:"comment:字段长度" json:"field_length"`
	FieldDecimalLength int                         `gorm:"comment:小数位数" json:"field_decimal_length"`
	InputType          enum.SysTableFieldInputType `gorm:"default:1;comment:输入类型" json:"input_type"`
	DefaultValue       string                      `gorm:"size:128;comment:默认值" json:"default_value"`
	DictCode           string                      `gorm:"size:128;comment:所用字典" json:"dict_code"`
	IsPrimaryKey       bool                        `gorm:"default:false;comment:是否主键" json:"is_primary_key"`
	IsIndex            bool                        `gorm:"default:false;comment:是否索引" json:"is_index"`
	IsQuickSearch      bool                        `gorm:"default:false;comment:是否快捷搜索" json:"is_quick_search"`
	IsAdvancedSearch   bool                        `gorm:"default:false;comment:是否高级搜索" json:"is_advanced_search"`
	IsSort             bool                        `gorm:"default:false;comment:是否可排序" json:"is_sort"`
	IsNull             bool                        `gorm:"default:true;comment:是否可空" json:"is_null"`
	OriginalFieldID    int                         `gorm:"comment:原字段ID" json:"original_field_id"`
}

type SysDict struct {
	Basic
	DictName  *string        `gorm:"size:128;comment:字典名称" json:"dict_name"`
	DictCode  *string        `gorm:"size:128;comment:字典编码" json:"dict_code"`
	DictItems *[]SysDictItem `gorm:"foreignKey:DictID;references:ID" json:"dict_items"`
}

type SysDictItem struct {
	Basic
	DictID    *int    `gorm:"comment:dict_id" json:"dict_id"`
	ItemName  *string `gorm:"size:128;comment:字典名称" json:"item_name"`
	ItemCode  *string `gorm:"size:128;comment:字典编码" json:"item_code"`
	ItemValue *string `gorm:"size:128;comment:字典值" json:"item_value"`
}

//func (st *SysDict) Value() (driver.Value, error) {
//	return json.Marshal(st.DictItems)
//}
//
//func (st *SysDict) Scan(value interface{}) error {
//	bytes, ok := value.([]byte)
//	if !ok {
//		return errors.New("Scan source is not []byte")
//	}
//	return json.Unmarshal(bytes, &st.DictItems)
//}
