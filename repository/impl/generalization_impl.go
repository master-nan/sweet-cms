/**
 * @Author: Nan
 * @Date: 2024/6/13 下午11:34
 */

package impl

import (
	"gorm.io/gorm"
	"sweet-cms/form/request"
	"sweet-cms/repository"
)

type GeneralizationRepositoryImpl struct {
	db *gorm.DB
}

func NewGeneralizationRepositoryImpl(db *gorm.DB) *GeneralizationRepositoryImpl {
	return &GeneralizationRepositoryImpl{db}
}

func (g GeneralizationRepositoryImpl) Query(basic request.Basic) (repository.GeneralizationListResult, error) {
	//TODO implement me
	panic("implement me")
}
