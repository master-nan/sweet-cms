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
	PID         int             `gorm:"type:int" json:"pid"`
	Name        string          `gorm:"size:32;comment:路由" json:"name"`
	Path        string          `gorm:"size:128;comment:路径" json:"path"`
	Component   string          `gorm:"size:64;comment:路由主体" json:"component"`
	Title       string          `gorm:"size:64;comment:显示标题" json:"title"`
	IsHidden    bool            `gorm:"default:false;comment:是否隐藏" json:"isHidden"`
	Sequence    uint8           `gorm:"comment:排序;type:tinyint" json:"sequence"`
	Option      string          `gorm:"size:64;comment:选项" json:"option"`
	Icon        *string         `gorm:"size:32;comment:图标" json:"icon"`
	Redirect    *string         `gorm:"size:128;comment:重定向地址" json:"redirect"`
	IsUnfold    bool            `gorm:"default:false;comment:默认展开" json:"isUnfold"`
	MenuButtons []SysMenuButton `gorm:"foreignKey:MenuID;references:ID" json:"menu_btns"`
}

type SysMenuButton struct {
	Basic
	MenuID   int                        `gorm:"comment:menu_id" json:"menu_id" binding:"required"`
	Name     string                     `gorm:"size:128;comment:按钮名称" json:"name" binding:"required"`
	Code     string                     `gorm:"size:128;comment:按钮编码" json:"code" binding:"required"`
	Memo     string                     `gorm:"size:128;comment:备注" json:"memo"`
	Position enum.SysMenuButtonPosition `gorm:"type:tinyint;default:1;comment:位置" json:"position" binding:"required"`
}

type SysMenuDataPermission struct {
	Basic
	UserID     int    `gorm:"comment:用户ID" json:"user_id"`
	MenuID     int    `gorm:"comment:菜单ID" json:"menu_id"`
	CompanyIDs string `gorm:"size:128;comment:公司ID集合" json:"company_ids"`
}

type SysRole struct {
	Basic
	Name  string    `gorm:"size:128;comment:角色名称" json:"name"`
	Memo  string    `gorm:"size:128;comment:备注" json:"memo"`
	Menus []SysMenu `gorm:"many2many:sys_role_menu;" json:"menus"`
}

type SysUser struct {
	Basic
	UserName     string     `gorm:"size:128;unique;comment:用户名" json:"user_name"`
	Password     string     `gorm:"size:128;comment:密码" json:"password"`
	Email        string     `gorm:"size:128;unique;comment:邮箱" json:"email"`
	PhoneNumber  string     `gorm:"size:128;unique;comment:电话" json:"phone_number"`
	IDCard       string     `gorm:"size:128;unique;comment:身份证号" json:"id_card"`
	EmployeeID   int        `gorm:"comment:员工ID" json:"employee_id"`
	GmtLastLogin CustomTime `gorm:"type:datetime;comment:最后登录时间" json:"gmt_last_login"`
	Language     string     `gorm:"size:32;comment:语言包" json:"language"`
	AccessTokens string     `gorm:"type:text;comment:用户最近5次Token" json:"access_tokens"`
	Roles        []SysRole  `gorm:"many2many:sys_user_role;" json:"roles"`
}

type SysUserRole struct {
	UserID int `gorm:"primaryKey;autoIncrement:false" json:"user_id"`
	RoleID int `gorm:"primaryKey;autoIncrement:false" json:"role_id"`
}

type SysRoleMenu struct {
	RoleID int `gorm:"primaryKey;autoIncrement:false" json:"role_id"`
	MenuID int `gorm:"primaryKey;autoIncrement:false" json:"menu_id"`
}

type SysGlobalDataPermission struct {
	Basic
	UserID     int    `gorm:"comment:用户ID" json:"user_id"`
	CompanyIDs string `gorm:"size:128;comment:公司ID集合" json:"company_ids"`
	IsAll      bool   `gorm:"default:false;comment:是否拥有全部公司权限" json:"is_all"`
}

type SysTable struct {
	Basic
	TableName      string             `gorm:"size:128;comment:表名" json:"table_name"`
	TableCode      string             `gorm:"size:128;comment:数据库中表名" json:"table_code"`
	TableType      enum.SysTableType  `gorm:"type:tinyint;default:1;comment:表类型" json:"table_type"`
	ParentID       int                `gorm:"comment:父节点ID" json:"parent_id"`
	SQL            string             `gorm:"type:text;comment:视图定义SQL" json:"sql"`
	TableFields    []SysTableField    `gorm:"foreignKey:TableID;references:ID" json:"table_fields"`
	TableRelations []SysTableRelation `gorm:"foreignKey:TableID"`
	TableIndexes   []SysTableIndex    `gorm:"foreignKey:TableID"`
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
	Dict               SysDict                     `gorm:"foreignKey:DictCode;references:DictCode" json:"dict"`
	IsPrimaryKey       bool                        `gorm:"default:false;comment:是否主键" json:"is_primary_key"`
	IsQuickSearch      bool                        `gorm:"default:false;comment:是否快捷搜索" json:"is_quick_search"`
	IsAdvancedSearch   bool                        `gorm:"default:false;comment:是否高级搜索" json:"is_advanced_search"`
	IsSort             bool                        `gorm:"default:false;comment:是否可排序" json:"is_sort"`
	IsNull             bool                        `gorm:"default:true;comment:是否可空" json:"is_null"`
	IsListShow         bool                        `gorm:"default:true;comment:是否列表显示" json:"is_list_show"`
	IsInsertShow       bool                        `gorm:"default:true;comment:是否新增显示" json:"is_insert_show"`
	IsUpdateShow       bool                        `gorm:"default:true;comment:是否更新显示" json:"is_update_show"`
	Sequence           uint8                       `gorm:"comment:排序;type:tinyint" json:"sequence"`
	OriginalFieldID    int                         `gorm:"comment:原字段ID" json:"original_field_id"`
	Binding            string                      `gorm:"size:256;comment:验证器" json:"binding"`        // 用于存储绑定规则
	FieldCategory      enum.SysTableFieldCategory  `gorm:"size:64;comment:字段类别" json:"field_category"` // 字段类别（普通字段、虚拟列、计算字段）
	Expression         *string                     `gorm:"size:256;comment:计算字段表达式" json:"expression"` // 计算字段表达式或虚拟列表达式
}

type SysTableIndex struct {
	Basic
	TableID     int             `gorm:"index;comment:表ID" json:"table_id"`
	IndexName   string          `gorm:"size:128;comment:索引名称" json:"index_name"`
	IsUnique    bool            `gorm:"comment:是否唯一索引" json:"is_unique"`
	IndexFields []SysTableField `gorm:"many2many:sys_table_index_field" json:"index_fields"`
}

type SysTableIndexField struct {
	IndexID int `gorm:"primaryKey;autoIncrement:false" json:"index_id"`
	FieldID int `gorm:"primaryKey;autoIncrement:false" json:"field_id"`
}

type SysTableRelation struct {
	Basic
	TableID        int                       `gorm:"index;comment:主表ID" json:"table_id"`
	RelatedTableID int                       `gorm:"index;comment:关联表ID" json:"related_table_id"`   // 关联的表的ID
	ReferenceKey   string                    `gorm:"size:128;comment:引用主表的字段" json:"reference_key"` // 主表对应字段
	ForeignKey     string                    `gorm:"size:128;comment:外键字段" json:"foreign_key"`      // 关联表 字段
	OnDelete       string                    `gorm:"size:128;comment:删除时策略" json:"on_delete"`
	OnUpdate       string                    `gorm:"size:128;comment:更新时策略" json:"on_update"`
	RelationType   enum.SysTableRelationType `gorm:"size:128;comment:关系类型" json:"relation_type"`
	ManyTableCode  *string                   `gorm:"size:128;comment:多对多关系中间表" json:"many_table_code"` // 多对多关系使用到的中间表
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
