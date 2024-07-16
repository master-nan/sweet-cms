/**
 * @Author: Nan
 * @Date: 2024/6/13 下午11:34
 */

package impl

import (
	"gorm.io/gorm"
	"sweet-cms/form/request"
	"sweet-cms/model"
	"sweet-cms/repository"
	"sweet-cms/repository/util"
)

type GeneralizationRepositoryImpl struct {
	db *gorm.DB
	*BasicImpl
}

func NewGeneralizationRepositoryImpl(db *gorm.DB, basicImpl *BasicImpl) *GeneralizationRepositoryImpl {
	return &GeneralizationRepositoryImpl{db, basicImpl}
}

func (g GeneralizationRepositoryImpl) Query(basic request.Basic, table model.SysTable) (repository.GeneralizationListResult, error) {
	result, err := util.DynamicQuery(g.db, basic, table)
	if err != nil {
		return repository.GeneralizationListResult{}, err
	}
	return result, nil
}
