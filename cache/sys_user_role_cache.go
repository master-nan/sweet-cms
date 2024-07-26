/**
 * @Author: Nan
 * @Date: 2024/7/25 下午11:19
 */

package cache

import (
	"sweet-cms/inter"
	"sweet-cms/model"
)

type SysUserRoleCache struct {
	*BasicCache[model.SysUserRole]
}

const UserRoleCacheKey = "USER_ROLE_CACHE_KEY_"

func NewSysUserRoleCache(cacheInterface inter.CacheInterface) *SysUserRoleCache {
	return &SysUserRoleCache{
		BasicCache: NewBasicCache[model.SysUserRole](cacheInterface, UserRoleCacheKey),
	}
}
