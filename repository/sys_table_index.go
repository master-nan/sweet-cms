/**
 * @Author: Nan
 * @Date: 2024/7/20 下午3:55
 */

package repository

import (
	"gorm.io/gorm"
	"sweet-cms/model"
)

type SysTableIndexRepository interface {
	BasicRepository
	GetTableIndexesByTableId(int) ([]model.SysTableIndex, error)
	GetTableIndexById(int) (model.SysTableIndex, error)
	DeleteTableIndexByTableId(*gorm.DB, int) error
}
