/**
 * @Author: Nan
 * @Date: 2024/5/17 下午3:38
 */

package request

import (
	"database/sql/driver"
	"sweet-cms/model"
)

type Basic struct {
	Page   int     `json:"page"`
	Num    int     `json:"num"`
	Query  []Query `json:"query"`
	Orders Order   `json:"order"`
}

type ExpressionType uint8

const (
	GT ExpressionType = iota + 1
	LT
	GTE
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

func (e ExpressionType) Value() (driver.Value, error) {
	return int(e), nil
}

type Query struct {
	Field      string                  `json:"field"`
	Expression ExpressionType          `json:"expression"`
	Value      interface{}             `json:"value"`
	Type       model.SysTableFieldType `json:"type"`
}

type Order struct {
	Field string `json:"field"`
	IsAsc bool   `json:"is_asc"`
}
