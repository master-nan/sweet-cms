/**
 * @Author: Nan
 * @Date: 2023/9/7 16:25
 */

package enum

import "database/sql/driver"

// DataPermissionsEnum 数据权限字典
type DataPermissionsEnum uint8

const (
	WHOLE   DataPermissionsEnum = iota + 1 //全部
	CUSTOM                                 //自定义
	TACITLY                                // 默认
)

func (dp DataPermissionsEnum) Value() (driver.Value, error) {
	return int(dp), nil
}

type SysMenuBtnPosition uint8

const (
	COLUMN SysMenuBtnPosition = iota + 1
	LINE
)

func (sbp SysMenuBtnPosition) Value() (driver.Value, error) {
	return int(sbp), nil
}

type SysTableType uint8

const (
	SYSTEM SysTableType = iota + 1
	VIEW
)

func (stt SysTableType) Value() (driver.Value, error) {
	return int(stt), nil
}

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
)

func (stf SysTableFieldType) Value() (driver.Value, error) {
	return int(stf), nil
}

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

func (s SysTableFieldInputType) Value() (driver.Value, error) {
	return int(s), nil
}