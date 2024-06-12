/**
 * @Author: Nan
 * @Date: 2024/6/3 下午9:59
 */

package request

// DictCreateReq 新增字典
type DictCreateReq struct {
	DictName string `json:"dict_name" binding:"required"`
	DictCode string `json:"dict_code" binding:"required"`
}

// DictUpdateReq 修改字典
type DictUpdateReq struct {
	ID       int    `json:"id" binding:"required"`
	DictName string `json:"dict_name" binding:"required"`
}

// DictItemCreateReq 新增字典明细
type DictItemCreateReq struct {
	DictID    int    `json:"dict_id" binding:"required"`
	ItemName  string `json:"item_name" binding:"required"`
	ItemCode  string `json:"item_code" binding:"required"`
	ItemValue string `json:"item_value" binding:"required"`
}

// DictItemUpdateReq 修改字典明细
type DictItemUpdateReq struct {
	ID        int    `json:"id" binding:"required"`
	ItemName  string `json:"item_name" binding:"required"`
	ItemCode  string `json:"item_code" binding:"required"`
	ItemValue string `json:"item_value" binding:"required"`
}
