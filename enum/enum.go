/**
 * @Author: Nan
 * @Date: 2023/9/7 16:25
 */

package enum

// DataPermissionsEnum 数据权限字典
type DataPermissionsEnum uint8

const (
	WHOLE   DataPermissionsEnum = iota + 1 //全部
	CUSTOM                                 //自定义
	TACITLY                                // 默认
)

type SysMenuButtonPosition uint8

const (
	COLUMN SysMenuButtonPosition = iota + 1
	LINE
)

type SysTableType uint8

const (
	SYSTEM SysTableType = iota + 1
	VIEW
)

type SysTableFieldType uint8

const (
	INT SysTableFieldType = iota + 1
	FLOAT
	VARCHAR
	TEXT
	BOOLEAN
	DATE
	DATETIME
	TIME
	TINYINT
)

type SysTableFieldInputType uint8

const (
	INPUT SysTableFieldInputType = iota + 1
	INPUT_NUMBER
	TEXTAREA
	SELECT
	DATE_PICKER
	DATETIME_PICKER
	TIME_PICKER
	YEAR_PICKER
	YREA_MONTH_PICKER
	FILE_PICKER
)

// ExpressionType 表达式
type ExpressionType uint8

const (
	GT  ExpressionType = iota + 1 // GT
	LT                            // LT
	GTE                           // GTE
	LTE
	EQ
	NE
	LIKE
	NOT_LIKE
	IN
	NOT_IN
	IS_NULL
	IS_NOT_NULL
)

type ExpressionLogic uint8

const (
	AND ExpressionLogic = iota + 1
	OR
)

type SysTableRelationType uint8

const (
	ONE_TO_ONE SysTableRelationType = iota + 1
	ONE_TO_MANY
	MANY_TO_ONE
	MANY_TO_MANY
)

type ArticleType string

const (
	DRAFT   ArticleType = "draft"
	REVIEW  ArticleType = "review"
	RELEASE ArticleType = "release"
	REJECT  ArticleType = "reject"
)

type SysTableFieldCategory string

const (
	NORMAL_FIELD     SysTableFieldCategory = "normal_field"     // 默认字段
	VIRTUAL_FIELD    SysTableFieldCategory = "virtual_field"    // 虚拟列
	CALCULATED_FIELD SysTableFieldCategory = "calculated_field" // 计算字段
)
