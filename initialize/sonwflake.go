/**
 * @Author: Nan
 * @Date: 2023/9/12 15:55
 */

package initialize

import (
	"sweet-cms/utils"
)

func InitSnowflake() (*utils.Snowflake, error) {
	return utils.NewSnowflake(1)
}
