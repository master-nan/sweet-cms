/**
 * @Author: Nan
 * @Date: 2023/9/7 16:25
 */

package enum

// DataPermissionsEnum 数据权限字典
type DataPermissionsEnum uint8

const (
	WHOLE   DataPermissionsEnum = iota + 1 //全部
	CUSTOM                                 //自定义
	TACITLY                                // 默认
)
