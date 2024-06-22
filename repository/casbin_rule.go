/**
 * @Author: Nan
 * @Date: 2024/6/22 下午2:30
 */

package repository

type CasbinRuleRepository interface {
	AddPolicy(params ...interface{}) (bool, error)
	RemovePolicy(params ...interface{}) (bool, error)
}
