/**
 * @Author: Nan
 * @Date: 2024/6/13 下午11:32
 */

package service

import (
	"sweet-cms/form/request"
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

func (s *GeneralizationService) Query(basic request.Basic) (repository.GeneralizationListResult, error) {
	panic("something")
}
