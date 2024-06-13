/**
 * @Author: Nan
 * @Date: 2024/6/13 下午11:33
 */

package repository

import (
	"sweet-cms/form/request"
)

type GeneralizationListResult struct {
	Data  []map[string]interface{} `json:"data"`
	Total int                      `json:"total"`
}

type GeneralizationRepository interface {
	Query(basic request.Basic) (GeneralizationListResult, error)
}
