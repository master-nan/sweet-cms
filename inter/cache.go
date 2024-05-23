/**
 * @Author: Nan
 * @Date: 2024/5/23 下午2:33
 */

package inter

import (
	"github.com/pkg/errors"
	"time"
)

var (
	ErrCacheMiss = errors.New("cache miss")
)

type CacheInterface interface {
	Get(key string, value interface{}) error
	Set(key string, value interface{}, expiration time.Duration) error
}
