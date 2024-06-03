/**
 * @Author: Nan
 * @Date: 2024/6/3 下午4:30
 */

package repository

import "sweet-cms/model"

type LogRepository interface {
	CreateLoginLog(model.LoginLog) (int, error)
	CreateAccessLog(model.AccessLog) (int, error)
}
