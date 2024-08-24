/**
 * @Author: Nan
 * @Date: 2024/7/26 下午5:14
 */

package request

type MenuCreateReq struct {
	Pid       int     `json:"pid"`
	Name      string  `json:"name"`
	Path      string  `json:"path"`
	Component string  `json:"component"`
	Title     string  `json:"title"`
	IsHidden  bool    `json:"isHidden"`
	Sequence  uint8   `json:"sequence"`
	Option    string  `json:"option"`
	Icon      *string `json:"icon"`
	Redirect  *string `json:"redirect"`
}

type MenuUpdateReq struct {
	Id        int     `json:"id" binding:"required"`
	Pid       *int    `json:"pid" binding:"required"`
	Name      string  `json:"name" binding:"required"`
	Path      string  `json:"path" binding:"required"`
	Component string  `json:"component" binding:"required"`
	Title     string  `json:"title" binding:"required"`
	IsHidden  *bool   `json:"isHidden" binding:"required"`
	Sequence  uint8   `json:"sequence" binding:"required"`
	Option    string  `json:"option"`
	Icon      *string `json:"icon"`
	Redirect  *string `json:"redirect"`
}
