/**
 * @Author: Nan
 * @Date: 2024/7/20 下午3:55
 */

package repository

import (
	"sweet-cms/model"
)

type SysTableIndexRepository interface {
	BasicRepository
	GetTableIndexesByTableId(int) ([]model.SysTableIndex, error)
}
