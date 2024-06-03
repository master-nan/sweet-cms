/**
 * @Author: Nan
 * @Date: 2024/6/3 下午2:50
 */

package repository

import "sweet-cms/model"

type SysConfigureRepository interface {
	GetSysConfigure() (model.SysConfigure, error)
	UpdateSysConfigure(d model.SysConfigure) error
}
