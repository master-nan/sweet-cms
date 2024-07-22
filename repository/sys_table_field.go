/**
 * @Author: Nan
 * @Date: 2024/7/20 下午3:27
 */

package repository

import (
	"gorm.io/gorm"
	"sweet-cms/model"
)

type SysTableFieldRepository interface {
	BasicRepository
	GetTableFieldById(int) (model.SysTableField, error)
	GetTableFieldsByTableId(int) ([]model.SysTableField, error)
	DeleteTableFieldByTableId(*gorm.DB, int) error
}
