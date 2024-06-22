/**
 * @Author: Nan
 * @Date: 2024/6/22 下午2:20
 */

package initialize

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

func InitCasbin(db *gorm.DB) (*casbin.Enforcer, error) {
	adapter, err := gormadapter.NewAdapterByDB(db) // 使用GORM适配器
	if err != nil {
		return nil, err
	}
	enforcer, err := casbin.NewEnforcer("casbin_model.conf", adapter)
	if err != nil {
		return nil, err
	}
	err = enforcer.LoadPolicy()
	if err != nil {
		return nil, err
	}
	return enforcer, nil
}
