/**
 * @Author: Nan
 * @Date: 2024/7/20 下午2:51
 */

package repository

import (
	"sweet-cms/model"
)

type SysDictItemRepository interface {
	BasicRepository
	GetSysDictItemById(int) (model.SysDictItem, error)
	GetSysDictItemsByDictId(int) ([]model.SysDictItem, error)
}
