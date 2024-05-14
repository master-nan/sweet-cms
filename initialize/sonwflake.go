/**
 * @Author: Nan
 * @Date: 2023/9/12 15:55
 */

package initialize

import (
	"sweet-cms/global"
	"sweet-cms/utils"
)

func SF() {
	sf, err := utils.NewSnowflake(1)
	if err != nil {
		panic(err)
	}
	global.SF = sf
}
