package model

import (
	"database/sql"
	"sweet-cms/enum"
)

type SysConfigure struct {
	Basic
	EnableCaptcha bool `gorm:"comment:登录验证码" json:"enableCaptcha"`
}

type SysMenu struct {
	Basic
	Pid         int                         `gorm:"type:int" json:"pid"`
	Name        string                      `gorm:"size:32;comment:路由" json:"name"`
	Path        string                      `gorm:"size:128;comment:路径" json:"path"`
	Component   string                      `gorm:"size:64;comment:路由主体" json:"component"`
	Title       string                      `gorm:"size:64;comment:显示标题" json:"title"`
	IsHidden    bool                        `gorm:"default:false;comment:是否隐藏" json:"isHidden"`
	Sequence    uint8                       `gorm:"comment:排序;type:tinyint" json:"sequence"`
	Option      string                      `gorm:"size:64;comment:选项" json:"option"`
	Icon        *string                     `gorm:"size:32;comment:图标" json:"icon"`
	Redirect    *string                     `gorm:"size:128;comment:重定向地址" json:"redirect"`
	IsUnfold    bool                        `gorm:"default:false;comment:默认展开" json:"isUnfold"`
	MenuButtons []SysMenuButton             `gorm:"foreignKey:MenuId;references:Id" json:"menuButtons"`
	Roles       []SysRole                   `gorm:"many2many:sys_role_menu" json:"roles"`
	Permissions []SysUserMenuDataPermission `gorm:"foreignKey:MenuId;references:Id" json:"permissions"`
	Children    []SysMenu                   `gorm:"-" json:"children"` // 子菜单
}

type SysMenuButton struct {
	Basic
	MenuId      int                        `gorm:"comment:menu_id" json:"menuId" binding:"required"`
	Name        string                     `gorm:"size:128;comment:按钮名称" json:"name" binding:"required"`
	Code        string                     `gorm:"size:128;comment:按钮编码" json:"code" binding:"required"`
	Memo        string                     `gorm:"size:128;comment:备注" json:"memo"`
	Position    enum.SysMenuButtonPosition `gorm:"type:tinyint;default:1;comment:位置" json:"position" binding:"required"`
	EventType   string                     `gorm:"size:64;comment:事件类型" json:"eventType"`
	EventAction string                     `gorm:"size:256;comment:事件动作" json:"eventAction"`
	BeforeHooks []string                   `gorm:"-" json:"beforeHooks"` // 执行事件前的钩子函数
	AfterHooks  []string                   `gorm:"-" json:"afterHooks"`  // 执行事件后的钩子函数
	Roles       []SysRole                  `gorm:"many2many:sys_role_menu_button" json:"roles"`
	Menus       []SysMenu                  `gorm:"many2many:sys_role_menu_button" json:"menus"`
}

type SysUserMenuDataPermission struct {
	Basic
	UserId     int     `gorm:"comment:用户Id" json:"userId"`
	MenuId     int     `gorm:"comment:菜单Id" json:"menuId"`
	CompanyIds string  `gorm:"size:128;comment:公司Id集合" json:"companyIds"`
	Menu       SysMenu `gorm:"foreignKey:MenuId;references:Id" json:"menu"`
	User       SysUser `gorm:"foreignKey:UserId;references:Id" json:"user"`
}

type SysUser struct {
	Basic
	UserName     string                      `gorm:"size:128;uniqueIndex:uni_user_name;comment:用户名" json:"userName"`
	Password     string                      `gorm:"size:128;comment:密码" json:"password"`
	Email        string                      `gorm:"size:128;uniqueIndex:uni_email;comment:邮箱" json:"email"`
	PhoneNumber  string                      `gorm:"size:128;uniqueIndex:uni_phone_number;comment:电话" json:"phoneNumber"`
	IdCard       string                      `gorm:"size:128;uniqueIndex:uni_id_card;comment:身份证号" json:"idCard"`
	EmployeeId   int                         `gorm:"comment:员工Id" json:"employeeId"`
	GmtLastLogin CustomTime                  `gorm:"type:datetime;comment:最后登录时间" json:"gmtLastLogin"`
	Language     string                      `gorm:"size:32;comment:语言包" json:"language"`
	AccessTokens string                      `gorm:"type:text;comment:用户最近5次Token" json:"accessTokens"`
	Roles        []SysRole                   `gorm:"many2many:sys_user_role" json:"roles"`
	Permissions  []SysUserMenuDataPermission `gorm:"foreignKey:UserId;references:Id" json:"permissions"`
}

type SysUserRole struct {
	UserId int `gorm:"primaryKey;autoIncrement:false" json:"userId"`
	RoleId int `gorm:"primaryKey;autoIncrement:false" json:"roleId"`
}

type SysRole struct {
	Basic
	Name    string          `gorm:"size:128;comment:角色名称" json:"name"`
	Memo    string          `gorm:"size:128;comment:备注" json:"memo"`
	Menus   []SysMenu       `gorm:"many2many:sys_role_menu" json:"menus"`
	Buttons []SysMenuButton `gorm:"many2many:sys_role_menu_button" json:"buttons"`
	Users   []SysUser       `gorm:"many2many:sys_user_role" json:"users"`
}

type SysRoleMenu struct {
	RoleId int `gorm:"primaryKey;autoIncrement:false" json:"roleId"`
	MenuId int `gorm:"primaryKey;autoIncrement:false" json:"menuId"`
}

type SysRoleMenuButton struct {
	RoleId   int `gorm:"primaryKey;autoIncrement:false" json:"roleId"`
	MenuId   int `gorm:"primaryKey;autoIncrement:false" json:"menuId"`
	ButtonId int `gorm:"primaryKey;autoIncrement:false" json:"buttonId"`
}

type SysGlobalDataPermission struct {
	Basic
	UserId     int    `gorm:"comment:用户Id" json:"user_id"`
	CompanyIds string `gorm:"size:128;comment:公司Id集合" json:"companyIds"`
	IsAll      bool   `gorm:"default:false;comment:是否拥有全部公司权限" json:"isAll"`
}

type SysTable struct {
	Basic
	TableName      string             `gorm:"size:128;comment:表名" json:"tableName"`
	TableCode      string             `gorm:"size:128;uniqueIndex:uni_table_code_index;comment:数据库中表名" json:"tableCode"`
	TableType      enum.SysTableType  `gorm:"type:tinyint;default:1;comment:表类型" json:"tableType"`
	ParentId       int                `gorm:"comment:父节点Id" json:"parentId"`
	SQL            string             `gorm:"type:text;comment:视图定义SQL" json:"sql"`
	TableFields    []SysTableField    `gorm:"foreignKey:TableId;references:Id" json:"tableFields"`
	TableRelations []SysTableRelation `gorm:"foreignKey:TableId"`
	TableIndexes   []SysTableIndex    `gorm:"foreignKey:TableId"`
}

type SysTableField struct {
	Basic
	TableId            int                         `gorm:"comment:table_id;uniqueIndex:union_uni_table_id_field_code_index" json:"tableId" binding:"required"`
	FieldName          string                      `gorm:"size:128;comment:列名" json:"fieldName"`
	FieldCode          string                      `gorm:"size:128;uniqueIndex:union_uni_table_id_field_code_index;comment:表字段名" json:"fieldCode"`
	FieldType          enum.SysTableFieldType      `gorm:"type:tinyint;default:1;comment:字段类型" json:"type"`
	FieldLength        int                         `gorm:"default:0;comment:字段长度" json:"fieldLength"`
	FieldDecimalLength int                         `gorm:"default:0;comment:小数位数" json:"fieldDecimalLength"`
	InputType          enum.SysTableFieldInputType `gorm:"type:tinyint;default:1;comment:输入类型" json:"inputType"`
	DefaultValue       *string                     `gorm:"size:128;comment:默认值" json:"defaultValue"`
	DictCode           *string                     `gorm:"size:128;comment:所用字典" json:"dictCode"`
	Dict               SysDict                     `gorm:"foreignKey:DictCode;references:DictCode" json:"dict"`
	IsPrimaryKey       bool                        `gorm:"default:false;comment:是否主键" json:"isPrimaryKey"`
	IsQuickSearch      bool                        `gorm:"default:false;comment:是否快捷搜索" json:"isQuickSearch"`
	IsAdvancedSearch   bool                        `gorm:"default:false;comment:是否高级搜索" json:"isAdvancedSearch"`
	IsSort             bool                        `gorm:"default:false;comment:是否可排序" json:"isSort"`
	IsNull             bool                        `gorm:"default:true;comment:是否可空" json:"isNull"`
	IsListShow         bool                        `gorm:"default:true;comment:是否列表显示" json:"isListShow"`
	IsInsertShow       bool                        `gorm:"default:true;comment:是否新增显示" json:"isInsertShow"`
	IsUpdateShow       bool                        `gorm:"default:true;comment:是否更新显示" json:"isUpdateShow"`
	Sequence           uint8                       `gorm:"comment:排序;type:tinyint" json:"sequence"`
	OriginalFieldId    int                         `gorm:"comment:原字段Id" json:"originalFieldId"`
	Binding            string                      `gorm:"size:256;comment:验证器" json:"binding"`        // 用于存储绑定规则
	FieldCategory      enum.SysTableFieldCategory  `gorm:"size:64;comment:字段类别" json:"fieldCategory"`  // 字段类别（普通字段、虚拟列、计算字段）
	Expression         *string                     `gorm:"size:256;comment:计算字段表达式" json:"expression"` // 计算字段表达式或虚拟列表达式
	Tag                *string                     `gorm:"size:256;comment:标签" json:"tag"`
}

type SysTableIndex struct {
	Basic
	TableId     int             `gorm:"index;comment:表Id" json:"tableId"`
	IndexName   string          `gorm:"size:128;comment:索引名称" json:"indexName"`
	IsUnique    bool            `gorm:"comment:是否唯一索引" json:"isUnique"`
	IndexFields []SysTableField `gorm:"many2many:sys_table_index_field" json:"indexFields"`
}

type SysTableIndexField struct {
	IndexId int `gorm:"primaryKey;autoIncrement:false" json:"indexId"`
	FieldId int `gorm:"primaryKey;autoIncrement:false" json:"fieldId"`
}

type SysTableRelation struct {
	Basic
	TableId        int                       `gorm:"index;comment:主表Id" json:"tableId"`
	RelatedTableId int                       `gorm:"index;comment:关联表Id" json:"relatedTableId"` // 关联的表的Id
	ReferenceKey   string                    `gorm:"size:128;comment:主表字段" json:"referenceKey"` // 主表对应字段
	ForeignKey     string                    `gorm:"size:128;comment:关联表字段" json:"foreignKey"`  // 关联表 字段
	OnDelete       string                    `gorm:"size:128;comment:删除时策略" json:"onDelete"`
	OnUpdate       string                    `gorm:"size:128;comment:更新时策略" json:"onUpdate"`
	RelationType   enum.SysTableRelationType `gorm:"size:128;comment:关系类型" json:"relationType"`
	ManyTableCode  string                    `gorm:"size:128;comment:多对多关系中间表" json:"manyTableCode"` // 多对多关系使用到的中间表
}

type SysDict struct {
	Basic
	DictName  string        `gorm:"size:128;comment:字典名称;uniqueIndex:uni_dict_name_index" json:"dictName"`
	DictCode  string        `gorm:"size:128;comment:字典编码;uniqueIndex:uni_dict_code_index" json:"dictCode"`
	DictItems []SysDictItem `gorm:"foreignKey:DictId;references:Id" json:"dictItems"`
}

type SysDictItem struct {
	Basic
	DictId    int    `gorm:"comment:dict_id" json:"dictId"`
	ItemName  string `gorm:"size:128;comment:字典名称" json:"itemName"`
	ItemCode  string `gorm:"size:128;comment:字典编码;uniqueIndex:uni_item_code_index" json:"itemCode"`
	ItemValue string `gorm:"size:128;comment:字典值" json:"itemValue"`
}

type TableColumnMate struct {
	ColumnName             string         `gorm:"column:COLUMN_NAME" json:"columnName"`                          // 列名
	OrdinalPosition        int            `gorm:"column:ORDINAL_POSITION" json:"ordinalPosition"`                // 列名所在位置，即排序-
	ColumnDefault          sql.NullString `gorm:"column:COLUMN_DEFAULT" json:"columnDefault"`                    // 默认值
	IsNullable             string         `gorm:"column:IS_NULLABLE" json:"isNullable"`                          // 是否可以为null
	DataType               string         `gorm:"column:DATA_TYPE" json:"dataType"`                              // 数据类型
	CharacterMaximumLength sql.NullInt64  `gorm:"column:CHARACTER_MAXIMUM_LENGTH" json:"characterMaximumLength"` // 字符类型最大长度
	CharacterOctetLength   sql.NullInt64  `gorm:"column:CHARACTER_OCTET_LENGTH" json:"characterOctetLength"`     // 字符类型最大字节长度
	NumericPrecision       sql.NullInt64  `gorm:"column:NUMERIC_PRECISION" json:"numericPrecision"`              // 数值型类的精度
	NumericScale           sql.NullInt64  `gorm:"column:NUMERIC_SCALE" json:"numericScale"`                      // 数值型列的小数位数
	DatetimePrecision      sql.NullInt64  `gorm:"column:DATETIME_PRECISION" json:"datetimePrecision"`            // 日期时间型列的精度
	ColumnType             string         `gorm:"column:COLUMN_TYPE" json:"columnType"`                          // 列的完整类型描述
	ColumnKey              string         `gorm:"column:COLUMN_KEY" json:"columnKey"`                            // 列是否被索引。
	Extra                  string         `gorm:"column:EXTRA" json:"extra"`                                     // 其他信息，是否自增
	ColumnComment          string         `gorm:"column:COLUMN_COMMENT" json:"columnComment"`                    // 列备注
}

type TableIndexMate struct {
	ColumnName string `gorm:"column:COLUMN_NAME" json:"columnName"` // 列名
	IndexName  string `gorm:"column:INDEX_NAME" json:"indexName"`   // 索引名称
	NonUnique  bool   `gorm:"column:NON_UNIQUE" json:"nonUnique"`   // 是否唯一索引
}
