/**
 * @Author: Nan
 * @Date: 2024/6/13 下午11:32
 */

package service

import (
	"sweet-cms/form/request"
	"sweet-cms/model"
	"sweet-cms/repository"
)

type GeneralizationService struct {
	generalizationRepo repository.GeneralizationRepository
}

func NewGeneralizationService(generalizationRepo repository.GeneralizationRepository) *GeneralizationService {
	return &GeneralizationService{
		generalizationRepo,
	}
}

func (gs *GeneralizationService) Query(basic request.Basic, table model.SysTable) (repository.GeneralizationListResult, error) {
	result, err := gs.generalizationRepo.Query(basic, table)
	if err != nil {
		return repository.GeneralizationListResult{}, err
	}
	return result, nil
}
