/**
 * @Author: Nan
 * @Date: 2024/7/20 下午3:48
 */

package repository

import (
	"sweet-cms/model"
)

type SysTableRelationRepository interface {
	BasicRepository
	GetTableRelationById(int) (model.SysTableRelation, error)
	GetTableRelationsByTableId(int) ([]model.SysTableRelation, error)
}
