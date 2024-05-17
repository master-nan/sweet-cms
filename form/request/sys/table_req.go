/**
 * @Author: Nan
 * @Date: 2024/5/17 上午11:39
 */

package sys

import "sweet-cms/model"

type TableCreateReq struct {
	TableName string             `json:"table_name" binding:"required"`
	TableCode string             `json:"table_code" binding:"required"`
	TableType model.SysTableType `json:"table_type" binding:"required"`
}

type TableUpdateReq struct {
	ID        int    `json:"id" binding:"required"`
	TableName string `json:"table_name" binding:"required"`
}
