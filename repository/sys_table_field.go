/**
 * @Author: Nan
 * @Date: 2024/7/20 下午3:27
 */

package repository

import (
	"sweet-cms/model"
)

type SysTableFieldRepository interface {
	BasicRepository
	GetTableFieldsByTableId(int) ([]model.SysTableField, error)
}
