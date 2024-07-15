/**
 * @Author: Nan
 * @Date: 2023/9/7 16:25
 */

package enum

// DataPermissionsEnum 数据权限字典
type DataPermissionsEnum uint8

const (
	Whole   DataPermissionsEnum = iota + 1 //全部
	Custom                                 //自定义
	Tacitly                                // 默认
)

// SysMenuButtonPosition 按钮位置字典
type SysMenuButtonPosition uint8

const (
	Column SysMenuButtonPosition = iota + 1
	Line
)

// SysTableType 表类型字典
type SysTableType uint8

const (
	System SysTableType = iota + 1 // 系统表
	View                           // 视图
)

// SysTableFieldType 字段数据库存储类型
type SysTableFieldType uint8

const (
	IntFieldType SysTableFieldType = iota + 1 //
	FloatFieldType
	VarcharFieldType
	TextFieldType
	BooleanFieldType
	DateFieldType
	DatetimeFieldType
	TimeFieldType
	TinyintFieldType
)

// SysTableFieldInputType 字段页面输入类型
type SysTableFieldInputType uint8

const (
	InputInputType SysTableFieldInputType = iota + 1
	InputNumberInputType
	TextareaInputType
	SelectInputType
	DatePickerInputType
	DatetimePickerInputType
	TimePickerInputType
	YearPickerInputType
	YearMonthPickerInputType
	FilePickerInputType
)

// ExpressionType 表达式
type ExpressionType uint8

const (
	Gt  ExpressionType = iota + 1 // Gt
	Lt                            // Lt
	Gte                           // Gte
	Lte
	Eq
	Ne
	Like
	NotLike
	In
	NotIn
	IsNull
	IsNotNull
)

// ExpressionLogic 关联方式
type ExpressionLogic uint8

const (
	And ExpressionLogic = iota + 1
	Or
)

// SysTableRelationType 表关系
type SysTableRelationType uint8

const (
	OneToOne SysTableRelationType = iota + 1
	OneToMany
	ManyToOne
	ManyToMany
)

// SysTableFieldCategory 表字段类型
type SysTableFieldCategory string

const (
	NormalField     SysTableFieldCategory = "normal_field"     // 默认字段
	VirtualField    SysTableFieldCategory = "virtual_field"    // 虚拟列
	CalculatedField SysTableFieldCategory = "calculated_field" // 计算字段
)
