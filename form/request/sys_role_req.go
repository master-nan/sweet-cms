/**
 * @Author: Nan
 * @Date: 2024/8/1 下午10:32
 */

package request

type RoleCreateReq struct {
	Name string `json:"name" binding:"required"`
	Memo string `json:"memo" binding:"required"`
}

type RoleUpdateReq struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Memo string `json:"memo" binding:"required"`
}

type RoleMenuCreateReq struct {
	RoleId int `json:"roleId" binding:"required"`
	MenuId int `json:"menuId" binding:"required"`
}
