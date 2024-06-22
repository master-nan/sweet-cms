/**
 * @Author: Nan
 * @Date: 2024/6/22 下午2:31
 */

package impl

import "github.com/casbin/casbin/v2"

type CasbinRuleRepositoryImpl struct {
	enforcer *casbin.Enforcer
}

func NewCasbinRuleRepositoryImpl(enforcer *casbin.Enforcer) *CasbinRuleRepositoryImpl {
	return &CasbinRuleRepositoryImpl{enforcer: enforcer}
}

func (c CasbinRuleRepositoryImpl) AddPolicy(params ...interface{}) (bool, error) {
	return c.enforcer.AddPolicy(params...)
}

func (c CasbinRuleRepositoryImpl) RemovePolicy(params ...interface{}) (bool, error) {
	return c.enforcer.RemovePolicy(params...)
}
