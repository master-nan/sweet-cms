/**
 * @Author: Nan
 * @Date: 2024/6/3 下午2:50
 */

package repository

import (
	"gorm.io/gorm"
	"sweet-cms/model"
)

type SysConfigureRepository interface {
	BasicRepository
	GetSysConfigure() (model.SysConfigure, error)
	UpdateSysConfigure(*gorm.DB, model.SysConfigure) error
}
