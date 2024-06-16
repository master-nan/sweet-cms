/**
 * @Author: Nan
 * @Date: 2024/6/13 下午11:33
 */

package repository

import (
	"sweet-cms/form/request"
	"sweet-cms/model"
)

type GeneralizationListResult struct {
	Data  []map[string]interface{} `json:"data"`
	Total int                      `json:"total"`
}

type GeneralizationRepository interface {
	Query(request.Basic, model.SysTable) (GeneralizationListResult, error)
}
