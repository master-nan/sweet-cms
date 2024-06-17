package model

import (
	"sweet-cms/enum"
)

type SysConfigure struct {
	Basic
	EnableCaptcha bool `gorm:"comment:登录验证码" json:"enable_captcha"`
}

type SysMenu struct {
	Basic
	PID       int          `gorm:"type:int" json:"pid"`
	Name      string       `gorm:"size:32;comment:路由" json:"name"`
	Path      string       `gorm:"size:128;comment:路径" json:"path"`
	Component string       `gorm:"size:64;comment:路由主体" json:"component"`
	Title     string       `gorm:"size:64;comment:显示标题" json:"title"`
	IsHidden  bool         `gorm:"default:false;comment:是否隐藏" json:"isHidden"`
	Sequence  uint8        `gorm:"comment:排序;type:tinyint" json:"sequence"`
	Option    string       `gorm:"size:64;comment:选项" json:"option"`
	Icon      *string      `gorm:"size:32;comment:图标" json:"icon"`
	Redirect  *string      `gorm:"size:128;comment:重定向地址" json:"redirect"`
	IsUnfold  bool         `gorm:"default:false;comment:默认展开" json:"isUnfold"`
	MenuBtns  []SysMenuBtn `gorm:"foreignKey:MenuId;references:ID" json:"menu_btns"`
}

type SysMenuBtn struct {
	Basic
	MenuID   int                     `gorm:"comment:menu_id" json:"menu_id" binding:"required"`
	Name     string                  `gorm:"size:128;comment:按钮名称" json:"name" binding:"required"`
	Code     string                  `gorm:"size:128;comment:按钮编码" json:"code" binding:"required"`
	Memo     string                  `gorm:"size:128;comment:备注" json:"memo"`
	Position enum.SysMenuBtnPosition `gorm:"type:tinyint;default:1;comment:位置" json:"position" binding:"required"`
}

type SysRole struct {
	Basic
	Name  string    `gorm:"size:128;comment:角色名称" json:"name"`
	Rs    string    `gorm:"size:128;comment:菜单ID集合" json:"rs"`
	Memo  string    `gorm:"size:128;comment:备注" json:"memo"`
	Users []SysUser `gorm:"foreignKey:RoleId;references:ID" json:"users"`
}

type SysUser struct {
	Basic
	UserName string  `gorm:"size:128;comment:用户名" json:"username"`
	RoleId   int     `gorm:"comment:角色ID" json:"roleId"`
	Password string  `gorm:"size:128;comment:密码" json:"-"`
	Role     SysRole `gorm:"foreignKey:RoleId" json:"role"`
}

type SysTable struct {
	Basic
	TableName   string            `gorm:"size:128;comment:表名" json:"table_name"`
	TableCode   string            `gorm:"size:128;comment:数据库中表名" json:"table_code"`
	TableType   enum.SysTableType `gorm:"type:tinyint;default:1;comment:表类型" json:"table_type"`
	ParentID    int               `gorm:"comment:父节点ID" json:"parent_id"`
	TableFields []SysTableField   `gorm:"foreignKey:TableID;references:ID" json:"table_fields"`
}

type SysTableField struct {
	Basic
	TableID            int                         `gorm:"comment:table_id" json:"table_id" binding:"required"`
	FieldName          string                      `gorm:"size:128;comment:列名" json:"field_name"`
	FieldCode          string                      `gorm:"size:128;comment:数据库中字段名" json:"field_code"`
	FieldType          enum.SysTableFieldType      `gorm:"type:tinyint;default:1;comment:字段类型" json:"type"`
	FieldLength        int                         `gorm:"default:0;comment:字段长度" json:"field_length"`
	FieldDecimalLength int                         `gorm:"default:0;comment:小数位数" json:"field_decimal_length"`
	InputType          enum.SysTableFieldInputType `gorm:"type:tinyint;default:1;comment:输入类型" json:"input_type"`
	DefaultValue       *string                     `gorm:"size:128;comment:默认值" json:"default_value"`
	DictCode           *string                     `gorm:"size:128;comment:所用字典" json:"dict_code"`
	IsPrimaryKey       bool                        `gorm:"default:false;comment:是否主键" json:"is_primary_key"`
	IsIndex            bool                        `gorm:"default:false;comment:是否索引" json:"is_index"`
	IsQuickSearch      bool                        `gorm:"default:false;comment:是否快捷搜索" json:"is_quick_search"`
	IsAdvancedSearch   bool                        `gorm:"default:false;comment:是否高级搜索" json:"is_advanced_search"`
	IsSort             bool                        `gorm:"default:false;comment:是否可排序" json:"is_sort"`
	IsNull             bool                        `gorm:"default:true;comment:是否可空" json:"is_null"`
	IsListShow         bool                        `gorm:"default:true;comment:是否列表显示" json:"is_list_show"`
	IsInsertShow       bool                        `gorm:"default:true;comment:是否新增显示" json:"is_insert_show"`
	IsUpdateShow       bool                        `gorm:"default:true;comment:是否更新显示" json:"is_update_show"`
	Sequence           uint8                       `gorm:"comment:排序;type:tinyint" json:"sequence"`
	OriginalFieldID    int                         `gorm:"comment:原字段ID" json:"original_field_id"`
}

type SysDict struct {
	Basic
	DictName  string        `gorm:"size:128;comment:字典名称" json:"dict_name"`
	DictCode  string        `gorm:"size:128;comment:字典编码" json:"dict_code"`
	DictItems []SysDictItem `gorm:"foreignKey:DictID;references:ID" json:"dict_items"`
}

type SysDictItem struct {
	Basic
	DictID    int    `gorm:"comment:dict_id" json:"dict_id"`
	ItemName  string `gorm:"size:128;comment:字典名称" json:"item_name"`
	ItemCode  string `gorm:"size:128;comment:字典编码" json:"item_code"`
	ItemValue string `gorm:"size:128;comment:字典值" json:"item_value"`
}
