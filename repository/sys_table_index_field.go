/**
 * @Author: Nan
 * @Date: 2024/7/22 上午10:18
 */

package repository

import (
	"gorm.io/gorm"
)

type SysTableIndexFieldRepository interface {
	BasicRepository
	DeleteTableIndexFieldByIndexId(*gorm.DB, int) error
	DeleteTableIndexFieldByIndexIds(*gorm.DB, []int) error
}
