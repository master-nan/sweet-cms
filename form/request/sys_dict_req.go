/**
 * @Author: Nan
 * @Date: 2024/6/3 下午9:59
 */

package request

// DictCreateReq 新增字典
type DictCreateReq struct {
	DictName string `json:"dictName" binding:"required"`
	DictCode string `json:"dictCode" binding:"required"`
}

// DictUpdateReq 修改字典
type DictUpdateReq struct {
	Id       int    `json:"id" binding:"required"`
	DictName string `json:"dictName" binding:"required"`
}

// DictItemCreateReq 新增字典明细
type DictItemCreateReq struct {
	DictId    int    `json:"dictId" binding:"required"`
	ItemName  string `json:"itemName" binding:"required"`
	ItemCode  string `json:"itemCode" binding:"required"`
	ItemValue string `json:"itemValue" binding:"required"`
}

// DictItemUpdateReq 修改字典明细
type DictItemUpdateReq struct {
	Id        int    `json:"id" binding:"required"`
	ItemName  string `json:"itemName" binding:"required"`
	ItemCode  string `json:"itemCode" binding:"required"`
	ItemValue string `json:"itemValue" binding:"required"`
}
