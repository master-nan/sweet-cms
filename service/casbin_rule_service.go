/**
 * @Author: Nan
 * @Date: 2024/6/22 下午2:34
 */

package service

import "sweet-cms/repository"

type CasbinRuleService struct {
	casbinRuleRepo repository.CasbinRuleRepository
}

func NewCasbinRuleService(casbinRuleRepo repository.CasbinRuleRepository) *CasbinRuleService {
	return &CasbinRuleService{casbinRuleRepo: casbinRuleRepo}
}

func (s *CasbinRuleService) AddPolicy(role, path, method string) (bool, error) {
	return s.casbinRuleRepo.AddPolicy(role, path, method)
}

func (s *CasbinRuleService) RemovePolicy(role, path, method string) (bool, error) {
	return s.casbinRuleRepo.RemovePolicy(role, path, method)
}
